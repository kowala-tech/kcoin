package api

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"github.com/kowala-tech/wallet-backend/application/command"
)

//NewBroadcastTransactionHandler creates an http.Handler for the use case of broadcast a signed transaction.
func NewBroadcastTransactionHandler(log log.Logger, cmd command.BroadcastTransactionHandler) http.Handler {
	return setHandlerCors(
		&broadcastTransactionHandler{
			logger:                  log,
			broadcastTransactionCmd: cmd,
		},
	)
}

type broadcastTransactionHandler struct {
	logger                  log.Logger
	broadcastTransactionCmd command.BroadcastTransactionHandler
}

func (h *broadcastTransactionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	v := mux.Vars(r)

	rawTx, ok := v["rawtx"]
	if !ok {
		h.logger.Log(
			"type",
			"alert",
			"msg",
			"raw transaction required",
			"action",
			"balance",
		)
		json.NewEncoder(w).Encode(getErrorResponse(errors.New("raw transaction required")))
		return
	}

	rawTxAsBytes, err := hex.DecodeString(rawTx)
	if err != nil {
		h.logger.Log(
			"type",
			"alert",
			"msg",
			fmt.Sprintf("Error with signed transaction: %s", err),
			"action",
			"balance",
		)
		json.NewEncoder(w).Encode(getErrorResponse(errors.New("error with signed transaction")))
		return
	}

	resp, err := h.broadcastTransactionCmd.Handle(ctx, rawTxAsBytes)
	if err != nil {
		h.logger.Log(
			"type",
			"alert",
			"msg",
			fmt.Sprintf("Error broadcasting transaction: %s", err),
			"action",
			"balance",
		)
		json.NewEncoder(w).Encode(getErrorResponse(fmt.Errorf("error broadcasting transaction: %s", err)))
		return
	}

	json.NewEncoder(w).Encode(resp)
}
