package impl

import (
	"fmt"
	"math/big"

	"github.com/kowala-tech/kcoin/client/accounts"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/consensus"
	"github.com/kowala-tech/kcoin/client/params"
	"github.com/kowala-tech/kcoin/e2e/cluster"
)

func (ctx *Context) IsMTokensBalanceExact(account string, expectedMTokens int64) error {
	acc, ok := ctx.accounts[account]
	if !ok {
		return fmt.Errorf("can't get account for %q", account)
	}

	return ctx.checkTokenBalance(acc, expectedMTokens)
}

func (ctx *Context) ITransferMTokens(mTokens int64, from, to string) error {
	fromAccount, ok := ctx.accounts[from]
	if !ok {
		return fmt.Errorf("can't get account for %q", from)
	}

	toAccount, ok := ctx.accounts[to]
	if !ok {
		return fmt.Errorf("can't get account for %q", to)
	}

	return ctx.sendTokensAndWait(fromAccount, toAccount, mTokens)
}

func (ctx *Context) sendTokensAndWait(from, to accounts.Account, tokens int64) error {
	res := &cluster.ExecResponse{}
	if err := ctx.execCommand(ctx.rpcNodeID, getTokenBalance(to.Address), res); err != nil {
		return err
	}
	currentBalanceBig, ok := new(big.Int).SetString(res.StdOut, 10)
	if !ok {
		return fmt.Errorf("incorrect mToken deposit %q of %s", res.StdOut, to.Address.String())
	}
	currentBalance := new(big.Int).Div(currentBalanceBig, big.NewInt(params.Kcoin)).Int64()

	return ctx.waiter.Do(
		func() error {
			return ctx.sendTokens(from, to, tokens)
		},
		func() error {
			return ctx.checkTokenBalance(to, currentBalance+tokens)
		})
}

func (ctx *Context) checkTokenBalance(account accounts.Account, expectedMTokens int64) error {
	weis := toWei(expectedMTokens)
	res := &cluster.ExecResponse{}
	if err := ctx.execCommand(ctx.rpcNodeID, getTokenBalance(account.Address), res); err != nil {
		return err
	}

	tokenBalance, ok := new(big.Int).SetString(res.StdOut, 10)
	if !ok {
		return fmt.Errorf("incorrect mToken balance %q for %s", res.StdOut, account.Address.String())
	}

	if tokenBalance.Cmp(weis) != 0 {
		return fmt.Errorf("account %s have %v mTokens, expected %v", account.Address.String(), tokenBalance, expectedMTokens)
	}

	return nil
}

func (ctx *Context) sendTokens(from, to accounts.Account, tokens int64) error {
	weis := toWei(tokens)

	musd, err := consensus.NewMUSD(ctx.client, ctx.chainID)
	if err != nil {
		return err
	}
	wallet, err := ctx.findWalletFor(from)
	if err != nil {
		return err
	}
	walletAccount, err := accounts.NewWalletAccount(wallet, from)
	if err != nil {
		return err
	}

	_, err = musd.Transfer(walletAccount, to.Address, weis, nil, "")
	if err != nil {
		return err
	}

	return nil
}
