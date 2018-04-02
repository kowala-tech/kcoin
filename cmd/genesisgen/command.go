package main

import (
	"time"
	"github.com/kowala-tech/kcoin/core"
	"github.com/kowala-tech/kcoin/params"
	"fmt"
	"math/big"
	"strings"
	"github.com/pkg/errors"
)

var (
	ErrEmptyMaxNumValidators = errors.New("max number of validators is mandatory")
)

type GenerateGenesisCommand struct {
	network string
	maxNumValidators string
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

	genesis := &core.Genesis{
		Timestamp: uint64(time.Now().Unix()),
		GasLimit:  4700000,
		Alloc:     make(core.GenesisAlloc),
		Config:    &params.ChainConfig{},
	}

	fmt.Printf("%v\n", network)
	fmt.Printf("%v\n", genesis)
	fmt.Printf("%v\n", maxNumValidators)

	return nil
}

func (h *GenerateGenesisCommandHandler) getMaxNumValidators(s string) (*big.Int, error) {
	if s = strings.TrimSpace(s); s == "" {
		return nil, ErrEmptyMaxNumValidators
	}



	return nil, nil
}
