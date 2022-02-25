package types

import (
	context "context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	auth "github.com/cosmos/cosmos-sdk/x/auth/types"

	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/cosmos/cosmos-sdk/x/staking/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	conntypes "github.com/cosmos/ibc-go/v3/modules/core/03-connection/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	ibcexported "github.com/cosmos/ibc-go/v3/modules/core/exported"

	abci "github.com/tendermint/tendermint/abci/types"
)

// StakingKeeper defines the contract expected by parent-chain ccv module from a Staking Module that will keep track
// of the parent validator set. This version of the interchain-security protocol will mirror the parent chain's changes
// so we do not need a registry module between the staking module and CCV.
type StakingKeeper interface {
	GetValidatorUpdates(ctx sdk.Context) []abci.ValidatorUpdate
	CompleteStoppedUnbonding(ctx sdk.Context, id uint64) (found bool, err error)
	UnbondingTime(ctx sdk.Context) time.Duration
	GetValidatorByConsAddr(ctx sdk.Context, consAddr sdk.ConsAddress) (validator types.Validator, found bool)
	// slash the validator and delegators of the validator, specifying offence height, offence power, and slash fraction
	Jail(sdk.Context, sdk.ConsAddress) // jail a validator
	Slash(sdk.Context, sdk.ConsAddress, int64, int64, sdk.Dec)
}

type SlashingKeeper interface {
	JailUntil(sdk.Context, sdk.ConsAddress, time.Time)
	GetValidatorSigningInfo(ctx sdk.Context, address sdk.ConsAddress) (info slashingtypes.ValidatorSigningInfo, found bool)
	SetValidatorSigningInfo(ctx sdk.Context, address sdk.ConsAddress, info slashingtypes.ValidatorSigningInfo)
	DowntimeJailDuration(sdk.Context) time.Duration
	SlashFractionDowntime(sdk.Context) sdk.Dec
	ClearValidatorMissedBlockBitArray(ctx sdk.Context, address sdk.ConsAddress)
}

// ChannelKeeper defines the expected IBC channel keeper
type ChannelKeeper interface {
	GetChannel(ctx sdk.Context, srcPort, srcChan string) (channel channeltypes.Channel, found bool)
	GetNextSequenceSend(ctx sdk.Context, portID, channelID string) (uint64, bool)
	SendPacket(ctx sdk.Context, channelCap *capabilitytypes.Capability, packet ibcexported.PacketI) error
	WriteAcknowledgement(ctx sdk.Context, chanCap *capabilitytypes.Capability, packet ibcexported.PacketI, acknowledgement ibcexported.Acknowledgement) error
	ChanCloseInit(ctx sdk.Context, portID, channelID string, chanCap *capabilitytypes.Capability) error
}

// PortKeeper defines the expected IBC port keeper
type PortKeeper interface {
	BindPort(ctx sdk.Context, portID string) *capabilitytypes.Capability
}

// ConnectionKeeper defines the expected IBC connection keeper
type ConnectionKeeper interface {
	GetConnection(ctx sdk.Context, connectionID string) (conntypes.ConnectionEnd, bool)
}

// ClientKeeper defines the expected IBC client keeper
type ClientKeeper interface {
	CreateClient(ctx sdk.Context, clientState ibcexported.ClientState, consensusState ibcexported.ConsensusState) (string, error)
	GetClientState(ctx sdk.Context, clientID string) (ibcexported.ClientState, bool)
	GetLatestClientConsensusState(ctx sdk.Context, clientID string) (ibcexported.ConsensusState, bool)
}

// TODO: Expected interfaces for distribution on parent and baby chains
// SlashingHooks event hooks for jailing and slashing validator
type SlashingHooks interface {
	// Is triggered when the validator missed too many blocks
	AfterValidatorDowntime(ctx sdk.Context, consAddr sdk.ConsAddress, power int64)
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
}

// AccountKeeper defines the expected account keeper used for simulations
type AccountKeeper interface {
	GetModuleAccount(ctx sdk.Context, name string) auth.ModuleAccountI
}

// IBCTransferKeeper defines the expected interface needed for distribution transfer
// of tokens from the consumer to the provider chain
type IBCTransferKeeper interface {
	SendTransfer(
		ctx sdk.Context,
		sourcePort,
		sourceChannel string,
		token sdk.Coin,
		sender sdk.AccAddress,
		receiver string,
		timeoutHeight clienttypes.Height,
		timeoutTimestamp uint64,
	) error
}

// IBCKeeper defines the expected interface needed for openning a
// channel
type IBCCoreKeeper interface {
	ChannelOpenInit(
		goCtx context.Context,
		msg *channeltypes.MsgChannelOpenInit,
	) (*channeltypes.MsgChannelOpenInitResponse, error)
}
