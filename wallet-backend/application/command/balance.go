package command

import (
	"context"
	"math/big"

	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/wallet-backend/blockchain"
)

//GetBalanceHandler is the domain use case that represents getting a balance from a specific address.
type GetBalanceHandler struct {
	Client blockchain.Client
}

//Handle executes the use case and returns a BalanceResponse with the results, or an error if there was
//a problem getting the Balance from the blockchain.
func (h *GetBalanceHandler) Handle(ctx context.Context, address common.Address) (*BalanceResponse, error) {
	balance, err := h.Client.BalanceAt(ctx, address, nil)
	if err != nil {
		return nil, err
	}

	resp := &BalanceResponse{
		Balance: balance,
	}

	return resp, nil
}

//BalanceResponse represents the response of the use case of getting the balance from an account.
type BalanceResponse struct {
	Balance *big.Int `json:"balance"`
}
