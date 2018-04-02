package main

import (
	"math/big"
	"errors"
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
)

type Options struct {
	network string
	maxValidators *big.Int
	unbondingPeriod int
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

	if options.maxValidators == nil {
		return ErrEmptyMaxNumValidators
	}

	return nil
}

