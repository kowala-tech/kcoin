package impl

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/kowala-tech/kcoin/client"
	"github.com/kowala-tech/kcoin/client/accounts"
	"github.com/kowala-tech/kcoin/client/core/types"
	"github.com/kowala-tech/kcoin/e2e/cluster"
)

func (ctx *Context) ITransferKUSD(kcoin int64, from, to string) error {
	return ctx.waiter.Do(
		func() error {
			var err error
			ctx.lastTxStartingBlock, err = ctx.client.BlockNumber(context.Background())
			if err != nil {
				return err
			}
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

	if receipt.Status != types.ReceiptStatusSuccessful {
		return false, nil
	}

	if tx.Hash().String() != receipt.TxHash.String() {
		return false, fmt.Errorf("sent transaction hash %q, receipt transaction hash %q", tx.Hash().String(), receipt.TxHash.String())
	}

	return true, nil
}

func (ctx *Context) OnlyOneTransactionIsDone() error {
	// wait some for new blocks
	time.Sleep(3 * time.Second)

	currentBlock, err := ctx.client.BlockNumber(context.Background())
	if err != nil {
		return err
	}

	var txs []*types.Transaction
	var txsLog string
	for i := ctx.lastTxStartingBlock.Uint64() + 1; i <= currentBlock.Uint64(); i++ {
		block, err := ctx.client.BlockByNumber(context.Background(), big.NewInt(int64(i)))
		if err != nil {
			return err
		}

		blockTxs := block.Transactions()
		txs = append(txs, blockTxs...)

		for _, tx := range blockTxs {
			txsLog += fmt.Sprintf("Block %d(%q), transaction with value %d and full info %s\n\n",
				block.NumberU64(), block.Hash().String(), tx.Value().Uint64(), tx.String())
		}
	}

	if len(txs) > 1 {
		return fmt.Errorf("expected single transaction with value %d and full info %s\n\n.\n\ngot:\n%s",
			ctx.lastTx.Value().Uint64(), ctx.lastTx.String(), txsLog)
	}

	return nil
}
func (ctx *Context) TransactionHashTheSame() error {
	txBlock, err := ctx.transactionBlock(ctx.lastTx)
	if err != nil {
		return err
	}

	command := fmt.Sprintf("web3.eth.getTransactionFromBlock('%d', 0);", txBlock.NumberU64())
	resp, err := ctx.nodeRunner.Exec(ctx.genesisValidatorNodeID, cluster.KcoinExecCommand(command))
	if err != nil {
		return err
	}

	type Hash struct {
		Hash string
	}

	txFromConsole := new(Hash)
	err = json.Unmarshal(fixUnquotedJSON(resp.StdOut), txFromConsole)
	if err != nil {
		return err
	}

	txFromRPC, err := ctx.client.TransactionInBlock(context.Background(), txBlock.Hash(), 0)
	if err != nil {
		return err
	}

	if txFromRPC.Hash().String() != txFromConsole.Hash {
		return fmt.Errorf("transaction hash via console %q, via RPC %q", txFromConsole.Hash, txFromRPC.Hash().String())
	}

	return nil
}

func (ctx *Context) transactionBlock(tx *types.Transaction) (*types.Block, error) {
	currentBlock, err := ctx.client.BlockNumber(context.Background())
	if err != nil {
		return nil, err
	}

	for i := 1; i <= int(currentBlock.Uint64()); i++ {
		block, err := ctx.client.BlockByNumber(context.Background(), big.NewInt(int64(i)))
		if err != nil {
			return nil, err
		}

		txs := block.Transactions()
		for _, blockTx := range txs {
			if blockTx.Hash() == tx.Hash() {
				return block, nil
			}

			if tx.To() == blockTx.To() && tx.Value().Uint64() == blockTx.Value().Uint64() {
				return nil, fmt.Errorf("got wrong transaction hash. expected %s. got %s",
					tx.Hash().String(), blockTx.Hash().String())
			}
		}
	}

	return nil, errors.New("the transaction is not in the chain")
}

func (ctx *Context) buildTx(from, to accounts.Account, kcoin int64) (*types.Transaction, error) {
	nonce, gp, gas, err := ctx.getTxParams(from, to, kcoin)
	if err != nil {
		return nil, err
	}

	tx := types.NewTransaction(nonce, to.Address, toWei(kcoin), gas, gp, nil)

	return ctx.AccountsStorage.SignTx(from, tx, ctx.chainID)
}

func (ctx *Context) getTxParams(from, to accounts.Account, kcoin int64) (uint64, *big.Int, uint64, error) {
	nonce, err := ctx.client.NonceAt(context.Background(), from.Address, nil)
	if err != nil {
		return 0, nil, 0, err
	}

	gasPrice, err := ctx.client.SuggestGasPrice(context.Background())
	if err != nil {
		return 0, nil, 0, err
	}

	gas, err := ctx.client.EstimateGas(context.Background(), kowala.CallMsg{
		From:     from.Address,
		To:       &to.Address,
		Value:    toWei(kcoin),
		GasPrice: gasPrice,
	})
	if err != nil {
		return 0, nil, 0, err
	}
	return nonce, gasPrice, gas, nil
}

func (ctx *Context) sendFunds(from, to accounts.Account, kcoin int64) (*types.Transaction, error) {
	tx, err := ctx.buildTx(from, to, kcoin)

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
