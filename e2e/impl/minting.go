package impl

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/kowala-tech/kcoin/client/accounts"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/consensus"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/ownership"
	"github.com/kowala-tech/kcoin/client/core/types"
)

func (ctx *Context) MintMTokens(m, n int, mTokens int64, to string) error {
	governors := ctx.mtokensGovernanceAccounts
	if len(governors) != n {
		return fmt.Errorf("expected %v governors in the genesis, there are %v", n, len(governors))
	}
	toAccount, ok := ctx.accounts[to]
	if !ok {
		return fmt.Errorf("can't get account for %q", to)
	}

	return ctx.mintTokensAndWait(governors[:m], toAccount, mTokens)
}

func (ctx *Context) mintTokensAndWait(governance []accounts.Account, to accounts.Account, tokens int64) error {
	c, err := consensus.Binding(ctx.client, ctx.chainID)
	if err != nil {
		return err
	}
	if err := c.MintInit(); err != nil {
		return err
	}

	transactionID, err := ctx.submitTransactionToMint(c, governance[0], to, tokens)
	if err != nil {
		return err
	}

	for _, acct := range governance[1:] {
		if err := ctx.confirmMintTransaction(c, acct, transactionID); err != nil {
			return err
		}
	}
	return nil
}

func (ctx *Context) submitTransactionToMint(c consensus.Consensus, acct accounts.Account, to accounts.Account, tokens int64) (*big.Int, error) {
	weis := toWei(tokens)
	var transaction common.Hash
	var transactionID *big.Int

	err := ctx.waiter.Do(
		func() error {
			tOpts := &accounts.TransactOpts{
				From: acct.Address,
				Signer: func(signer types.Signer, address common.Address, tx *types.Transaction) (*types.Transaction, error) {
					return ctx.AccountsStorage.SignTx(acct, tx, ctx.chainID)
				},
			}
			tx, err := c.Mint(tOpts, to.Address, weis)
			if err != nil {
				return err
			}
			transaction = tx

			return nil
		},
		func() error {
			receipt, err := ctx.client.TransactionReceipt(context.Background(), transaction)
			if err != nil {
				return err
			}
			if receipt.Status != types.ReceiptStatusSuccessful {
				return fmt.Errorf("transaction status is %v, expected %v", receipt.Status, types.ReceiptStatusSuccessful)
			}
			contract := c.MultiSigWalletContract()
			event := new(ownership.MultiSigWalletSubmission)
			for _, log := range receipt.Logs {
				err := contract.UnpackLog(event, "Submission", *log)
				if err == nil {
					break
				}
			}
			if event == nil {
				return errors.New("submission event not found.")
			}
			transactionID = event.TransactionId
			return nil
		})
	return transactionID, err
}

func (ctx *Context) confirmMintTransaction(c consensus.Consensus, acct accounts.Account, transactionID *big.Int) error {
	var transaction common.Hash
	return ctx.waiter.Do(
		func() error {
			tOpts := &accounts.TransactOpts{
				From: acct.Address,
				Signer: func(signer types.Signer, address common.Address, tx *types.Transaction) (*types.Transaction, error) {
					return ctx.AccountsStorage.SignTx(acct, tx, ctx.chainID)
				},
			}
			tx, err := c.Confirm(tOpts, transactionID)
			if err != nil {
				return err
			}
			transaction = tx

			return nil
		},
		func() error {
			receipt, err := ctx.client.TransactionReceipt(context.Background(), transaction)
			if err != nil {
				return err
			}
			if receipt.Status != types.ReceiptStatusSuccessful {
				return fmt.Errorf("transaction status is %v, expected %v", receipt.Status, types.ReceiptStatusSuccessful)
			}
			return nil
		})
}
