package command

import (
	"context"

	"github.com/kowala-tech/kcoin/wallet-backend/blockchain"
)

//BroadcastTransactionHandler represents the use case of sending a signed transaction to the Blockchain.
type BroadcastTransactionHandler struct {
	Client blockchain.Client
}

//Handle executes the use case to send a signed raw transaction to the blockchain. It returns error or
//BroadcastTransactionResponse with the status.
func (h *BroadcastTransactionHandler) Handle(ctx context.Context, rawTx []byte) (*BroadcastTransactionResponse, error) {
	err := h.Client.SendRawTransaction(ctx, rawTx)
	if err != nil {
		return nil, err
	}

	return &BroadcastTransactionResponse{
		Status: "ok",
	}, nil
}

//BroadcastTransactionResponse represents the response with information for the use case of broadcast
//a signed transaction to the blockchain.
type BroadcastTransactionResponse struct {
	Status string `json:"status"`
}
