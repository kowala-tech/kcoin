package features

import (
	"context"
	"fmt"
	"errors"
	"math/big"

	"github.com/kowala-tech/kcoin"
	"github.com/kowala-tech/kcoin/accounts"
	"github.com/kowala-tech/kcoin/core/types"
	"github.com/kowala-tech/kcoin/cluster"
)

func (ctx *Context) CurrentBlock() (uint64, error) {
	block, err := ctx.client.BlockNumber(context.Background())
	if err != nil {
		return 0, err
	}

	return block.Uint64(), nil
}

func (ctx *Context) ITransferKUSD(kcoin int64, from, to string) error {
	return ctx.waiter.Do(
		func() error {
			var err error
			ctx.lastTx, err = ctx.sendFunds(ctx.accounts[from], ctx.accounts[to], kcoin)
			return err
		},
		func() error {
			isInBlockchain, err := ctx.isTransactionInBlockchain(ctx.lastTx)
			if err != nil {
				return err
			}
			if !isInBlockchain {
				return fmt.Errorf("tx %q is not in the blockchain", ctx.lastTx.String())
			}
			return nil
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
	if err != nil {
		return false, err
	}
	return receipt.Status == types.ReceiptStatusSuccessful, nil
}

func (ctx *Context) TransactionHashTheSame() error {
	txBlock, err := ctx.transactionBlock(ctx.lastTx)
	if err != nil {
		return err
	}

	command := fmt.Sprintf("web3.eth.getTransactionFromBlock('%x', 0);", txBlock.NumberU64())
	resp, err := ctx.nodeRunner.Exec(ctx.genesisValidatorNodeID, cluster.KcoinExecCommand(command))
	if err != nil {
		return err
	}

	//ctx.lastTx.Hash() != resp.StdOut

	fmt.Println("!!!!!!", resp.StdOut)
	return nil
}

func (ctx *Context) transactionBlock(tx *types.Transaction) (*types.Block, error) {
	currentBlock, err := ctx.client.BlockNumber(context.Background())
	if err != nil {
		return nil, err
	}

	fmt.Println("Searching for", tx.String(), tx.Value().String(), tx.ChainID().String())

	for i:=1; i <= int(currentBlock.Uint64()); i++ {
		block, err := ctx.client.BlockByNumber(context.Background(), big.NewInt(int64(i)))
		if err != nil {
			return nil, err
		}

		txs := block.Transactions()
		for _, blockTx := range txs {
			fmt.Println("Got", block.NumberU64(), blockTx.String())
			if blockTx.Hash() == tx.Hash() {
				return block, nil
			}
		}
	}

	return nil, errors.New("the transaction is not in the chain")
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
	var tx *types.Transaction
	return tx, ctx.waiter.Do(
		func() error {
			var err error
			tx, err = ctx.sendFunds(from, to, kcoins)
			return err
		},
		func() error {
			balance, err := ctx.client.BalanceAt(context.Background(), to.Address, nil)
			if err != nil {
				return err
			}
			if balance.Cmp(toWei(kcoins)) != 0 {
				return fmt.Errorf("want %d, got %d coins", toWei(kcoins).Uint64(), balance.Uint64())
			}
			return nil
		})
}
