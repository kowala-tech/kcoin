package tx

import (
	"context"
	"time"

	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/core/types"
	"github.com/kowala-tech/kcoin/client/log"
)

type Backend interface {
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
}

// WaitMinedWithTimeout waits for tx to be mined on the blockchain within a given period of time.
func WaitMinedWithTimeout(backend Backend, txHash common.Hash, duration time.Duration) (*types.Receipt, error) {
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	return WaitMined(ctx, backend, txHash)
}

// WaitMined waits for tx to be mined on the blockchain.
// It stops waiting when the context is canceled.
func WaitMined(ctx context.Context, backend Backend, txHash common.Hash) (*types.Receipt, error) {
	queryTicker := time.NewTicker(time.Second)
	defer queryTicker.Stop()

	logger := log.New("hash", txHash)
	for {
		receipt, err := backend.TransactionReceipt(ctx, txHash)
		if receipt != nil {
			return receipt, nil
		}
		if err != nil {
			return nil, err
		} else {
			logger.Trace("Transaction not yet mined")
		}
		// Wait for the next round.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-queryTicker.C:
		}
	}
}
