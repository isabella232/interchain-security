package keeper

import (
	"encoding/binary"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v3/modules/core/24-host"
	"github.com/cosmos/ibc-go/v3/modules/core/exported"
	"github.com/cosmos/interchain-security/x/ccv/child/types"
	ccv "github.com/cosmos/interchain-security/x/ccv/types"
	utils "github.com/cosmos/interchain-security/x/ccv/utils"
	abci "github.com/tendermint/tendermint/abci/types"
)

// OnRecvPacket sets the pending validator set changes that will be flushed to ABCI on Endblock
// and set the unbonding time for the packet so that we can WriteAcknowledgement after unbonding time is over.
func (k Keeper) OnRecvPacket(ctx sdk.Context, packet channeltypes.Packet, newChanges ccv.ValidatorSetChangePacketData) exported.Acknowledgement {
	// packet is not sent on parent channel, return error acknowledgement and close channel
	if parentChannel, ok := k.GetParentChannel(ctx); ok && parentChannel != packet.DestinationChannel {
		ack := channeltypes.NewErrorAcknowledgement(
			fmt.Sprintf("packet sent on a channel %s other than the established parent channel %s", packet.DestinationChannel, parentChannel),
		)
		chanCap, _ := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(packet.DestinationPort, packet.DestinationChannel))
		k.channelKeeper.ChanCloseInit(ctx, packet.DestinationPort, packet.DestinationChannel, chanCap)
		return &ack
	}
	if status := k.GetChannelStatus(ctx, packet.DestinationChannel); status != ccv.VALIDATING {
		// Set CCV channel status to Validating and set parent channel
		k.SetChannelStatus(ctx, packet.DestinationChannel, ccv.VALIDATING)
		k.SetParentChannel(ctx, packet.DestinationChannel)
	}

	// Set pending changes by accumulating changes from this packet with all prior changes
	var pendingChanges []abci.ValidatorUpdate
	currentChanges, exists := k.GetPendingChanges(ctx)
	if !exists {
		pendingChanges = newChanges.ValidatorUpdates
	} else {
		pendingChanges = utils.AccumulateChanges(currentChanges.ValidatorUpdates, newChanges.ValidatorUpdates)
	}
	k.SetPendingChanges(ctx, ccv.ValidatorSetChangePacketData{ValidatorUpdates: pendingChanges})

	// Save unbonding time and packet
	unbondingTime := ctx.BlockTime().Add(types.UnbondingTime)
	k.SetUnbondingTime(ctx, packet.Sequence, uint64(unbondingTime.UnixNano()))
	k.SetUnbondingPacket(ctx, packet.Sequence, packet)
	// ack will be sent asynchronously
	return nil
}

// UnbondMaturePackets will iterate over the unbonding packets in order and write acknowledgements for all
// packets that have finished unbonding.
func (k Keeper) UnbondMaturePackets(ctx sdk.Context) error {
	store := ctx.KVStore(k.storeKey)
	unbondingIterator := sdk.KVStorePrefixIterator(store, []byte(types.UnbondingTimePrefix))
	defer unbondingIterator.Close()

	currentTime := uint64(ctx.BlockTime().UnixNano())

	channelID, ok := k.GetParentChannel(ctx)
	if !ok {
		return sdkerrors.Wrap(channeltypes.ErrChannelNotFound, "parent channel not set")
	}
	channelCap, ok := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(types.PortID, channelID))
	if !ok {
		return sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	for unbondingIterator.Valid() {
		sequence := types.GetSequenceFromUnbondingTimeKey(unbondingIterator.Key())
		if currentTime > binary.BigEndian.Uint64(unbondingIterator.Value()) {
			// write successful ack and delete unbonding information
			packet, err := k.GetUnbondingPacket(ctx, sequence)
			if err != nil {
				return err
			}
			ack := channeltypes.NewResultAcknowledgement([]byte{byte(1)})
			k.channelKeeper.WriteAcknowledgement(ctx, channelCap, packet, ack)
			k.DeleteUnbondingTime(ctx, sequence)
			k.DeleteUnbondingPacket(ctx, sequence)
		} else {
			break
		}
		unbondingIterator.Next()
	}
	return nil
}

// SendPacket sends a packet that initiates the given validator
// slashing and jailing on the provider chain.
func (k Keeper) SendPacket(ctx sdk.Context, val abci.Validator, slashFraction, jailedUntil int64) error {

	// check the setup
	channelID, ok := k.GetParentChannel(ctx)
	if !ok {
		return sdkerrors.Wrap(channeltypes.ErrChannelNotFound, "parent channel not set")
	}
	channel, ok := k.channelKeeper.GetChannel(ctx, types.PortID, channelID)
	if !ok {
		return sdkerrors.Wrapf(channeltypes.ErrChannelNotFound, "channel not found for channel ID: %s", channelID)
	}
	channelCap, ok := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(types.PortID, channelID))
	if !ok {
		return sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	// get the next sequence
	sequence, found := k.channelKeeper.GetNextSequenceSend(ctx, types.PortID, channelID)
	if !found {
		return sdkerrors.Wrapf(
			channeltypes.ErrSequenceSendNotFound,
			"source port: %s, source channel: %s", types.PortID, channelID,
		)
	}

	// add the last ValsetUpdateId to the packet data so that the provider
	// can find the block height when the downtime happened
	valsetUpdateId := k.GetLastUnbondingPacket(ctx).ValsetUpdateId
	if valsetUpdateId == 0 {
		return sdkerrors.Wrapf(ccv.ErrInvalidChildState, "last valset update id not set")
	}
	packetData := ccv.NewValidatorDowtimePacketData(val, valsetUpdateId, slashFraction, jailedUntil)
	packetDataBytes := packetData.GetBytes()

	// send ValidatorDowntime infractions in IBC packet
	packet := channeltypes.NewPacket(
		packetDataBytes, sequence,
		types.PortID, channelID,
		channel.Counterparty.PortId, channel.Counterparty.ChannelId,
		clienttypes.Height{}, uint64(ccv.GetTimeoutTimestamp(ctx.BlockTime()).UnixNano()),
	)
	if err := k.channelKeeper.SendPacket(ctx, channelCap, packet); err != nil {
		return err
	}

	return nil
}

func (k Keeper) OnAcknowledgementPacket(ctx sdk.Context, packet channeltypes.Packet, ack channeltypes.Acknowledgement) error {
	return nil
}
