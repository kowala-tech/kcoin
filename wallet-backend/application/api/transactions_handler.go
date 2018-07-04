package api

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"math/big"

	"strconv"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/wallet-backend/application/command"
)

const (
	fromBlockRequestVar = "fromblock"
	toBlockRequestVar   = "toblock"
)

//NewGetTransactionsHandler returns an http.Handler for the use case of getting transactions of a given
//address.
func NewGetTransactionsHandler(log log.Logger, cmd command.GetTransactionsHandler) http.Handler {
	return setHandlerCors(
		&getTransactionsHandler{
			logger:             log,
			getTransactionsCmd: cmd,
		},
	)
}

type getTransactionsHandler struct {
	logger             log.Logger
	getTransactionsCmd command.GetTransactionsHandler
}

func (h *getTransactionsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	v := mux.Vars(r)

	account, okFrom := v["account"]
	if !okFrom {
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

	from, to, err := h.parseRange(v)
	if err != nil {
		json.NewEncoder(w).Encode(getErrorResponse(err))
		return
	}

	cmd := command.GetTransactions{
		Address: common.HexToAddress(account),
		From:    from,
		To:      to,
	}

	resp, err := h.getTransactionsCmd.Handle(ctx, cmd)
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

//parseRange parses from request vars the from and to values to use as range
//to get transactions from.
func (h *getTransactionsHandler) parseRange(vars map[string]string) (from *big.Int, to *big.Int, err error) {
	if vars[fromBlockRequestVar] == "" || vars[toBlockRequestVar] == "" {
		return from, to, nil
	}

	f, err := strconv.Atoi(vars[fromBlockRequestVar])
	if err != nil {
		return nil, nil, errors.New("invalid from field")
	}

	t, err := strconv.Atoi(vars[toBlockRequestVar])
	if err != nil {
		return nil, nil, errors.New("invalid to field")
	}

	from = big.NewInt(int64(f))
	to = big.NewInt(int64(t))

	return from, to, nil
}
