package main

import (
	"encoding/json"
	"github.com/kowala-tech/kcoin/kcoin/genesis"
	"io"
)

type generateGenesisFileCommandHandler struct {
	w io.Writer
}

func (h *generateGenesisFileCommandHandler) handle(options genesis.Options) error {
	gns, err := genesis.GenerateGenesis(options)
	if err != nil {
		return err
	}

	out, _ := json.MarshalIndent(gns, "", "  ")

	if _, err = h.w.Write(out); err != nil {
		return err
	}

	return nil
}
