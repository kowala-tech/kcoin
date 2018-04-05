package main

import (
	"encoding/json"
	"github.com/kowala-tech/kcoin/cmd/genesis/kcoin"
	"io"
)

type GenerateGenesisCommandHandler struct {
	w io.Writer
}

func (h *GenerateGenesisCommandHandler) Handle(command kcoin.GenesisOptions) error {
	genesis, err := kcoin.GenerateGenesis(command)
	if err != nil {
		return err
	}

	out, _ := json.MarshalIndent(genesis, "", "  ")

	_, err = h.w.Write(out)
	if err != nil {
		return err
	}

	return nil
}