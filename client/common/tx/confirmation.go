package tx

import (
	"context"
	"errors"
	"time"

	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/core/types"
	"github.com/kowala-tech/kcoin/client/log"
	"github.com/kowala-tech/kcoin/client/params"
)

type Backend interface {
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
}

// WaitMinedWithTimeout waits for tx to be mined on the blockchain within a given period of time.
func WaitMinedWithTimeout(ctx context.Context, backend Backend, txHash common.Hash) (*types.Receipt, error) {
	return WaitMined(ctx, backend, txHash)
}

// WaitMined waits for tx to be mined on the blockchain.
// It stops waiting when the context is canceled.
func WaitMined(ctx context.Context, backend Backend, txHash common.Hash) (*types.Receipt, error) {
	queryTicker := time.NewTicker(time.Duration(params.PreCommitDeltaDuration) * time.Millisecond)
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

		select {
		case <-ctx.Done():
			logger.Trace("Transaction failed by timeout", "err", ctx.Err())
			if receipt == nil && ctx.Err() == nil {
				return nil, errors.New("context deadline exceeded")
			}
			return nil, ctx.Err()
		case <-queryTicker.C:
			// Wait for the next round.
		}
	}
}
