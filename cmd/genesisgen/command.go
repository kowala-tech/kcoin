package main

import (
	"time"
	"github.com/kowala-tech/kcoin/core"
	"github.com/kowala-tech/kcoin/params"
	"fmt"
	"math/big"
	"strings"
	"github.com/pkg/errors"
	"github.com/kowala-tech/kcoin/common"
)

var (
	ErrEmptyMaxNumValidators = errors.New("max number of validators is mandatory")
	ErrEmptyUnbondingPeriod = errors.New("unbonding period in days is mandatory")
	ErrEmptyWalletAddressValidator = errors.New("Wallet address of genesis validator is mandatory")
	ErrInvalidWalletAddressValidator = errors.New("unbonding period in days is mandatory")
)

type GenerateGenesisCommand struct {
	network string
	maxNumValidators string
	unbondingPeriod string
	walletAddressGenesisValidator string
}

type GenerateGenesisCommandHandler struct {
}

func (h *GenerateGenesisCommandHandler) Handle(command GenerateGenesisCommand) error {
	network, err := NewNetwork(command.network)
	if err != nil {
		return err
	}

	maxNumValidators, err := h.getMaxNumValidators(command.maxNumValidators)
	if err != nil {
		return err
	}

	unbondingPeriod, err := h.getUnbondingPeriod(command.unbondingPeriod)
	if err != nil {
		return err
	}

	walletAddressGenValidator, err := h.getWalletAddressGenesisValidator(command.walletAddressGenesisValidator)
	if err != nil {
		return err
	}

	genesis := &core.Genesis{
		Timestamp: uint64(time.Now().Unix()),
		GasLimit:  4700000,
		Alloc:     make(core.GenesisAlloc),
		Config:    &params.ChainConfig{},
	}

	fmt.Printf("%v\n", network)
	fmt.Printf("%v\n", genesis)
	fmt.Printf("%v\n", maxNumValidators)
	fmt.Printf("%v\n", unbondingPeriod)
	fmt.Printf("%v\n", walletAddressGenValidator)

	return nil
}

func (h *GenerateGenesisCommandHandler) getMaxNumValidators(s string) (*big.Int, error) {
	var numValidators *big.Int

	if s = strings.TrimSpace(s); s == "" {
		return nil, ErrEmptyMaxNumValidators
	}

	return numValidators, nil
}
func (handler *GenerateGenesisCommandHandler) getUnbondingPeriod(uP string) (*big.Int, error) {
	if text := strings.TrimSpace(uP); text == "" {
		return nil, ErrEmptyUnbondingPeriod
	}

	return nil, nil
}
func (handler *GenerateGenesisCommandHandler) getWalletAddressGenesisValidator(wA string) (*common.Address, error) {
	stringAddr := wA

	if text := strings.TrimSpace(wA); text == "" {
		return nil, ErrEmptyWalletAddressValidator
	}

	if strings.HasPrefix(stringAddr, "0x") {
		stringAddr = strings.TrimPrefix(stringAddr, "0x")
	}

	if len(stringAddr) != 40 {
		return nil, ErrInvalidWalletAddressValidator
	}

	return nil, nil
}

