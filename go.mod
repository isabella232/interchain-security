module github.com/cosmos/interchain-security

go 1.16

require (
	github.com/cosmos/cosmos-sdk v0.45.2-0.20220210215401-58c103ca4daf
	github.com/cosmos/ibc-go/v3 v3.0.0-alpha1.0.20220210141024-fb2f0416254b
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2
	github.com/google/go-cmp v0.5.7 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/gravity-devs/liquidity v1.4.5
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/prometheus/common v0.30.0 // indirect
	github.com/rakyll/statik v0.1.7
	github.com/rs/zerolog v1.25.0 // indirect
	github.com/spf13/cast v1.4.1
	github.com/spf13/cobra v1.3.0
	github.com/strangelove-ventures/packet-forward-middleware v1.0.1
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/spm v0.1.9
	github.com/tendermint/tendermint v0.34.14
	github.com/tendermint/tm-db v0.6.4
	google.golang.org/genproto v0.0.0-20220118154757-00ab72f36ad5
	google.golang.org/grpc v1.44.0
	google.golang.org/protobuf v1.27.1
)

replace (
	github.com/99designs/keyring => github.com/cosmos/keyring v1.1.7-0.20210622111912-ef00f8ac3d76
	github.com/cosmos/ibc-go/v3 => github.com/rigelrozanski/ibc-go/v3 v3.0.0-alpha1.0.20220214205135-4e834c9c6adc
	// github.com/cosmos/cosmos-sdk => github.com/cosmos/cosmos-sdk v0.45.2-0.20220210215401-58c103ca4daf
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	google.golang.org/grpc => google.golang.org/grpc v1.33.2
)
