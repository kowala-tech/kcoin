package runtime

import (
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/core"
	"github.com/kowala-tech/kcoin/client/core/vm"
)

func NewEnv(cfg *Config) *vm.VM {
	context := vm.Context{
		CanTransfer: core.CanTransfer,
		Transfer:    core.Transfer,
		GetHash:     func(uint64) common.Hash { return common.Hash{} },

		Origin:           cfg.Origin,
		Coinbase:         cfg.Coinbase,
		BlockNumber:      cfg.BlockNumber,
		Time:             cfg.Time,
		ComputeCapacity:  cfg.ComputeLimit,
		ComputeUnitPrice: cfg.ComputeUnitPrice,
	}

	return vm.New(context, cfg.State, cfg.ChainConfig, cfg.VMConfig)
}
