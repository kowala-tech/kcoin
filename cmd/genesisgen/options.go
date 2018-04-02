package main

import (
	"math/big"
	"errors"
	"github.com/kowala-tech/kcoin/common"
)

var (
	MainNetwork  = "main"
	TestNetwork  = "test"
	OtherNetwork = "other"

	AvailableNetworks = map[string]bool{
		MainNetwork : true,
		TestNetwork : true,
		OtherNetwork: true,
	}

	ErrInvalidNetwork = errors.New("invalid network, use main, test or other")
	ErrEmptyMaxNumValidators = errors.New("maximum number of validators is mandatory")
	ErrEmptyUnbondingPeriod = errors.New("unbound period in days is mandatory")
)

type Options struct {
	network string
	maxValidators *big.Int
	unbondingPeriod *big.Int
	genesisWalletAddr string
	accounts PrefundAccounts
	optional OptionalOpts
}

type PrefundAccounts []string

type OptionalOpts struct {
	consensusEngine string
	smartContractsAccount string
	message string
}

func validateOptions(options *Options) error {
	if !AvailableNetworks[options.network] {
		return ErrInvalidNetwork
	}

	if options.maxValidators == nil || options.maxValidators.Cmp(common.Big0) == 0 {
		return ErrEmptyMaxNumValidators
	}

	if options.unbondingPeriod == nil {
		return ErrEmptyUnbondingPeriod
	}

	return nil
}
