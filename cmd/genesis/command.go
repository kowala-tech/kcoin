package main

import (
	"encoding/json"
	"io"

	"github.com/kowala-tech/kcoin/kcoin/genesis"
)

type generateGenesisFileCommandHandler struct {
	w io.Writer
}

func (h *generateGenesisFileCommandHandler) handle(options *genesis.Options) error {
	gns, err := genesis.Generate(options)
	if err != nil {
		return err
	}

	out, _ := json.MarshalIndent(gns, "", "  ")

	if _, err = h.w.Write(out); err != nil {
		return err
	}

	return nil
}
