package oracle

import (
	"context"
	"math/big"
	"time"

	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/core/types"
	"github.com/kowala-tech/kcoin/client/log"
	"github.com/kowala-tech/kcoin/client/params"
)

var (
	oracleEpochStart = new(big.Int).Sub(params.OracleEpochDuration, params.OracleUpdatePeriod)
)

type Backend interface {
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
}

// waitMined waits for tx to be mined on the blockchain.
// It stops waiting when the context is canceled.
func waitMined(ctx context.Context, b Backend, txHash common.Hash) (*types.Receipt, error) {
	queryTicker := time.NewTicker(time.Second)
	defer queryTicker.Stop()

	logger := log.New("hash", txHash)
	for {
		receipt, err := b.TransactionReceipt(ctx, txHash)
		if receipt != nil {
			return receipt, nil
		}
		if err != nil {
			logger.Trace("Receipt retrieval failed", "err", err)
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

func IsUpdatePeriod(blockNumber *big.Int) bool {
	return new(big.Int).Mod(blockNumber, params.OracleEpochDuration).Cmp(oracleEpochStart) >= 0
}
