package main

import (
	"time"
	"github.com/kowala-tech/kcoin/core"
	"github.com/kowala-tech/kcoin/params"
)

type GenerateGenesisCommand struct {
}

func (c *GenerateGenesisCommand) Run(options *Options) error {
	err := validateOptions(options)
	if err != nil {
		return err
	}

	genesis := &core.Genesis{
		Timestamp: uint64(time.Now().Unix()),
		GasLimit:  4700000,
		Alloc:     make(core.GenesisAlloc),
		Config:    &params.ChainConfig{},
	}

	switch options.network {
	case "main":
		genesis.Config.ChainID = params.MainnetChainConfig.ChainID
	case "test":
		genesis.Config.ChainID = params.TestChainConfig.ChainID
	case "other":

	default:

	}

	return nil
}
