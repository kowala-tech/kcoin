package main

import (
	"encoding/json"
	"io"
	"github.com/kowala-tech/kcoin/kcoin/genesis"
)

type GenerateGenesisCommandHandler struct {
	w io.Writer
}

func (h *GenerateGenesisCommandHandler) Handle(command genesis.GenesisOptions) error {
	gns, err := genesis.GenerateGenesis(command)
	if err != nil {
		return err
	}

	out, _ := json.MarshalIndent(gns, "", "  ")

	if _, err = h.w.Write(out); err != nil {
		return err
	}

	return nil
}