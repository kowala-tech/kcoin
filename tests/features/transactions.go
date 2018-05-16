package features

import (
	"context"
	"fmt"
	"time"

	"github.com/kowala-tech/kcoin"
	"github.com/kowala-tech/kcoin/accounts"
	"github.com/kowala-tech/kcoin/core/types"
)

func (ctx *Context) ITransferKUSD(kcoin int64, from, to string) error {
	currentBlock, err := ctx.currentBlock()
	if err != nil {
		return err
	}

	tx, err := ctx.sendFunds(ctx.accounts[from], ctx.accounts[to], kcoin)
	if err != nil {
		return err
	}

	err = ctx.waitBlocksFrom(currentBlock, 1)
	if err != nil {
		return err
	}

	isInBlockchain, err := ctx.isTransactionInBlockchain(tx)
	if err != nil {
		return err
	}

	if !isInBlockchain {
		return fmt.Errorf("tx %q is not in the blockchain by the block %d", tx.String(), currentBlock+1)
	}

	return nil
}

func (ctx *Context) ITryTransferKUSD(kcoin int64, from, to string) error {
	tx, err := ctx.sendFunds(ctx.accounts[from], ctx.accounts[to], kcoin)
	ctx.lastTx = tx
	ctx.lastTxErr = err
	return nil
}

func (ctx *Context) LastTransactionFailed() error {
	if ctx.lastTxErr != nil {
		return nil // Failed at submitting the transaction, all good
	}

	isInBlockchain, err := ctx.isTransactionInBlockchain(ctx.lastTx)
	if err != nil {
		return err
	}
	if isInBlockchain {
		return fmt.Errorf("the last transaction is part of the blockchain, but shouldn't")
	}
	return nil
}

func (ctx *Context) isTransactionInBlockchain(tx *types.Transaction) (bool, error) {
	receipt, err := ctx.client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		return false, err
	}
	return receipt.Status == types.ReceiptStatusSuccessful, nil
}

func (ctx *Context) sendFunds(from, to accounts.Account, kcoin int64) (*types.Transaction, error) {
	nonce, err := ctx.client.NonceAt(context.Background(), from.Address, nil)
	if err != nil {
		return nil, err
	}

	gp, err := ctx.client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	gas, err := ctx.client.EstimateGas(context.Background(), kowala.CallMsg{
		From:     from.Address,
		To:       &to.Address,
		Value:    toWei(kcoin),
		GasPrice: gp,
	})
	if err != nil {
		return nil, err
	}

	tx := types.NewTransaction(nonce, to.Address, toWei(kcoin), gas, gp, nil)

	tx, err = ctx.AccountsStorage.SignTx(from, tx, ctx.chainID)
	if err != nil {
		return nil, err
	}

	return tx, ctx.client.SendTransaction(context.Background(), tx)
}

func (ctx *Context) sendFundsAndWait(from, to accounts.Account, kcoins int64) (*types.Transaction, error) {
	currentBlock, err := ctx.currentBlock()
	if err != nil {
		return nil, err
	}

	tx, err := ctx.sendFunds(from, to, kcoins)
	if err != nil {
		return nil, err
	}

	err = ctx.waitBlocksFrom(currentBlock, 1)
	if err != nil {
		return nil, err
	}

	balance, err := ctx.client.BalanceAt(context.Background(), to.Address, nil)
	if err != nil {
		return nil, err
	}

	if balance.Cmp(toWei(kcoins)) != 0 {
		return nil, fmt.Errorf("want %d, got %d coins", toWei(kcoins).Uint64(), balance.Uint64())
	}

	return tx, nil
}

// Do executes the command on the node and waits 1 block then
func (ctx *Context) Do(f func() error) error {
	currentBlock, err := ctx.currentBlock()
	if err != nil {
		return err
	}

	if err = f(); err != nil {
		return err
	}

	err = ctx.waitBlocksFrom(currentBlock, 1)
	if err != nil {
		return err
	}

	return nil
}

func (ctx *Context) waitBlocksFrom(block, n uint64) error {
	t := time.NewTicker(200 * time.Millisecond)
	timeout := time.NewTimer(20 * time.Second)
	defer t.Stop()

	var (
		err      error
		newBlock uint64
	)

waitLoop:
	for {
		select {
		case <-timeout.C:
			return fmt.Errorf("timeout. started with block %d, finished with %d", block, newBlock)
		case <-t.C:
			newBlock, err = ctx.currentBlock()
			if err != nil {
				return err
			}

			blocks := newBlock - block

			if blocks >= n {
				break waitLoop
			}
		}
	}

	return nil
}

func (ctx *Context) currentBlock() (uint64, error) {
	block, err := ctx.client.BlockNumber(context.Background())
	if err != nil {
		return 0, err
	}

	return block.Uint64(), nil
}
