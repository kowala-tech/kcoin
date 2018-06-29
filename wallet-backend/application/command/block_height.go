package command

import (
	"context"
	"math/big"

	"github.com/kowala-tech/kcoin/wallet-backend/blockchain"
)

//GetBlockHeightHandler represents the use case of getting the block height from the
type GetBlockHeightHandler struct {
	Client blockchain.Client
}

//Handle executes the use case of getting the block height of the blockchain and returns a BlockHeightResponse
func (h *GetBlockHeightHandler) Handle(c context.Context) (*BlockHeightResponse, error) {
	blockHeight, err := h.Client.BlockNumber(c)
	if err != nil {
		return nil, err
	}

	resp := &BlockHeightResponse{
		BlockHeight: blockHeight,
	}

	return resp, nil
}

//BlockHeightResponse represents the response of the use case of getting the block height of the Blockchain
type BlockHeightResponse struct {
	BlockHeight *big.Int `json:"block_height"`
}
