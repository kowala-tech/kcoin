package consensus

import (
	"log"
	"math/big"

	"github.com/kowala-tech/kcoin/client/common"
)

//go:generate gencodec -type Config -field-override configMarshaling -formats toml -out gen_config.go

type MiningConfig struct {
	Coinbase common.Address `toml:",omitempty"`

	Deposit *big.Int `toml:",omitempty"`

	ExtraData []byte `toml:",omitempty"`

	// Logger is a custom logger to use with the p2p.Server.
	Logger log.Logger `toml:",omitempty"`
}
