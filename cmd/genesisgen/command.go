package main

import (
	"fmt"
	"github.com/kowala-tech/kcoin/common"
	"github.com/kowala-tech/kcoin/core"
	"github.com/kowala-tech/kcoin/params"
	"github.com/pkg/errors"
	"math/big"
	"strings"
	"time"
	"bytes"
)

var (
	ErrEmptyMaxNumValidators                        = errors.New("max number of validators is mandatory")
	ErrEmptyUnbondingPeriod                         = errors.New("unbonding period in days is mandatory")
	ErrEmptyWalletAddressValidator                  = errors.New("Wallet address of genesis validator is mandatory")
	ErrInvalidWalletAddressValidator                = errors.New("Wallet address of genesis validator is invalid")
	ErrEmptyPrefundedAccounts                       = errors.New("empty prefunded accounts, at least the validator wallet address should be included")
	ErrWalletAddressValidatorNotInPrefundedAccounts = errors.New("prefunded accounts should include genesis validator account")
	ErrInvalidAddressInPrefundedAccounts = errors.New("address in prefunded accounts is invalid")
)

type GenerateGenesisCommand struct {
	network                       string
	maxNumValidators              string
	unbondingPeriod               string
	walletAddressGenesisValidator string
	prefundedAccounts             []PrefundedAccount
}

type PrefundedAccount struct {
	walletAddress string
	balance       int64
}

type validPrefundedAccount struct {
	walletAddress *common.Address
	balance       *big.Int
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

	walletAddressValidator, err := h.createWalletAddress(command.walletAddressGenesisValidator)
	if err != nil {
		return err
	}

	validPrefundedAccounts, err := h.validatePrefundedAccounts(command.prefundedAccounts)
	if err != nil {
		return err
	}

	if !h.prefundedIncludesValidatorWallet(validPrefundedAccounts, walletAddressValidator) {
		return ErrWalletAddressValidatorNotInPrefundedAccounts
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
	fmt.Printf("%v\n", walletAddressValidator)
	fmt.Printf("%v\n", validPrefundedAccounts)

	return nil
}

func (h *GenerateGenesisCommandHandler) getMaxNumValidators(s string) (*big.Int, error) {
	var numValidators *big.Int

	if s = strings.TrimSpace(s); s == "" {
		return nil, ErrEmptyMaxNumValidators
	}

	return numValidators, nil
}
func (h *GenerateGenesisCommandHandler) getUnbondingPeriod(uP string) (*big.Int, error) {
	if text := strings.TrimSpace(uP); text == "" {
		return nil, ErrEmptyUnbondingPeriod
	}

	return nil, nil
}

func (h *GenerateGenesisCommandHandler) createWalletAddress(wA string) (*common.Address, error) {
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

	bigaddr, _ := new(big.Int).SetString(stringAddr, 16)
	address := common.BigToAddress(bigaddr)

	return &address, nil
}

func (h *GenerateGenesisCommandHandler) validatePrefundedAccounts(accounts []PrefundedAccount) ([]*validPrefundedAccount, error) {
	var validAccounts []*validPrefundedAccount

	if len(accounts) == 0 {
		return nil, ErrEmptyPrefundedAccounts
	}

	for _, a := range accounts {
		address, err := h.createWalletAddress(a.walletAddress)
		if err != nil {
			return nil, ErrInvalidAddressInPrefundedAccounts
		}

		balance := big.NewInt(a.balance)

		validAccount := &validPrefundedAccount{
			walletAddress: address,
			balance:       balance,
		}

		validAccounts = append(validAccounts, validAccount)
	}

	return validAccounts, nil
}

func (h *GenerateGenesisCommandHandler) prefundedIncludesValidatorWallet(
	accounts []*validPrefundedAccount,
	addresses *common.Address,
) bool {
	for _, account := range accounts {
		if bytes.Equal(account.walletAddress.Bytes(), addresses.Bytes()) {
			return true
		}
	}

	return false
}