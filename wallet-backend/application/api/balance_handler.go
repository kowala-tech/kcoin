package api

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/wallet-backend/application/command"
)

//NewBalanceHandler returns an http.Handler to use as an endpoint api to get the balance. This handler wraps
//for http the use case to get the balance from an account.
func NewBalanceHandler(log log.Logger, cmd command.GetBalanceHandler) http.Handler {
	return setHandlerCors(
		&balanceHandler{
			logger:     log,
			balanceCmd: cmd,
		},
	)
}

type balanceHandler struct {
	logger     log.Logger
	balanceCmd command.GetBalanceHandler
}

func (h *balanceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	v := mux.Vars(r)

	account, ok := v["account"]
	if !ok {
		h.logger.Log(
			"type",
			"alert",
			"msg",
			"account required",
			"action",
			"balance",
		)
		json.NewEncoder(w).Encode(getErrorResponse(errors.New("account required")))
		return
	}

	a := common.HexToAddress(account)

	resp, err := h.balanceCmd.Handle(ctx, a)
	if err != nil {
		h.logger.Log(
			"type",
			"alert",
			"msg",
			"error getting balance "+err.Error(),
			"action",
			"balance",
		)
		json.NewEncoder(w).Encode(getErrorResponse(err))
		return
	}

	json.NewEncoder(w).Encode(resp)
}
