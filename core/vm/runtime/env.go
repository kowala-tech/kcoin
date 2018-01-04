package runtime

import (
	"math/big"

	"github.com/kowala-tech/kUSD/common"
	"github.com/kowala-tech/kUSD/core"
	"github.com/kowala-tech/kUSD/core/vm"
)

func NewEnv(cfg *Config) *vm.EVM {
	context := vm.Context{
		CanTransfer: core.CanTransfer,
		Transfer:    core.Transfer,
		GetHash:     func(uint64) common.Hash { return common.Hash{} },

		Origin:      cfg.Origin,
		Coinbase:    cfg.Coinbase,
		BlockNumber: cfg.BlockNumber,
		Time:        cfg.Time,
		GasLimit:    new(big.Int).SetUint64(cfg.GasLimit),
		GasPrice:    cfg.GasPrice,
	}

	return vm.NewEVM(context, cfg.State, cfg.ChainConfig, cfg.EVMConfig)
}
