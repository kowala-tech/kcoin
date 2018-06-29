package blockchain

import (
	"context"
	"math/big"

	"github.com/kowala-tech/kcoin/common"
)

//Client is an interface of a generic client to connect to a blockchain instance.
type Client interface {
	BlockNumber(ctx context.Context) (*big.Int, error)
	BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error)
	SendRawTransaction(ctx context.Context, rawTx []byte) error
}
