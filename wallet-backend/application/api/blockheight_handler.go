package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/kowala-tech/kcoin/wallet-backend/application/command"
)

//NewBlockHeightHandler returns an http.Handler for the use case of getting the Block Height from the Blockchain.
func NewBlockHeightHandler(logger log.Logger, cmd command.GetBlockHeightHandler) http.Handler {
	return setHandlerCors(
		&blockHeightHandler{
			logger:      logger,
			getBlockCmd: cmd,
		},
	)
}

type blockHeightHandler struct {
	logger      log.Logger
	getBlockCmd command.GetBlockHeightHandler
}

func (h *blockHeightHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	resp, err := h.getBlockCmd.Handle(ctx)
	if err != nil {
		h.logger.Log(
			"type",
			"alert",
			"msg",
			"error getting block "+err.Error(),
			"action",
			"blockheight",
		)
		json.NewEncoder(w).Encode(getErrorResponse(err))
		return
	}

	json.NewEncoder(w).Encode(resp)
}
