package features

import (
	"context"
	"fmt"
	"time"

	kowala "github.com/kowala-tech/kcoin"
	"github.com/kowala-tech/kcoin/accounts"
	"github.com/kowala-tech/kcoin/common"
	"github.com/kowala-tech/kcoin/core/types"
)

func (ctx *Context) ITransferKUSD(kcoin int64, from, to string) error {

	tx, err := ctx.sendFunds(ctx.accounts[from], ctx.accounts[to], kcoin)
	if err != nil {
		return err
	}

	return common.WaitFor("transaction in the blockhain", 1*time.Second, 5*time.Second, func() bool {
		isInBlockchain, err := ctx.isTransactionInBlockchain(tx)
		return err == nil && isInBlockchain
	})
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
	return receipt.Status == types.ReceiptStatusSuccessful, err
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
	tx, err := ctx.sendFunds(from, to, kcoins)
	if err != nil {
		return nil, err
	}
	return tx, common.WaitFor("account receives the balance", 1*time.Second, 10*time.Second, func() bool {
		balance, err := ctx.client.BalanceAt(context.Background(), to.Address, nil)
		if err != nil {
			return false
		}
		return balance.Cmp(toWei(kcoins)) == 0
	})
}
