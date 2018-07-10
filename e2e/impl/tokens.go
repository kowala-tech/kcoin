package impl

import (
	"fmt"
	"math/big"

	"github.com/kowala-tech/kcoin/client/accounts"
	"github.com/kowala-tech/kcoin/client/common/hexutil"
	"github.com/kowala-tech/kcoin/client/knode"
	"github.com/kowala-tech/kcoin/client/params"
	"github.com/kowala-tech/kcoin/e2e/cluster"
)

func (ctx *ValidationContext) sendTokensAndWait(from, to accounts.Account, tokens int64) error {
	res := &cluster.ExecResponse{}
	if err := ctx.execCommand(getTokenBalance(to.Address), res); err != nil {
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

func (ctx *ValidationContext) sendTokens(from, to accounts.Account, tokens int64) error {
	weis := toWei(tokens)
	hexWeis := hexutil.Big(*weis)
	args := knode.TransferArgs{
			From:  from.Address,
			To:    &to.Address,
			Value: &hexWeis,
	}

	res := &cluster.ExecResponse{}
	return ctx.execCommand(transferTokens(args), res)
}

func (ctx *ValidationContext) mintTokensAndWait(from, to accounts.Account, tokens int64) error {
	res := &cluster.ExecResponse{}
	if err := ctx.execCommand(getTokenBalance(to.Address), res); err != nil {
		return err
	}
	currentBalanceBig, ok := new(big.Int).SetString(res.StdOut, 10)
	if !ok {
		return fmt.Errorf("incorrect mToken deposit %q of %s", res.StdOut, to.Address.String())
	}
	currentBalance := new(big.Int).Div(currentBalanceBig, big.NewInt(params.Kcoin)).Int64()

	return ctx.waiter.Do(
		func() error {
			return ctx.mintTokens(from, to, tokens, AccountPass)
		},
		func() error {
			return ctx.checkTokenBalance(to, currentBalance+tokens)
		})
}

func (ctx *ValidationContext) mintTokens(from, to accounts.Account, tokens int64, pass string) error {
	weis := toWei(tokens)
	hexWeis := hexutil.Big(*weis)
	args := knode.TransferArgs{
		From:  from.Address,
		To:    &to.Address,
		Value: &hexWeis,
	}

	res := &cluster.ExecResponse{}
	return ctx.execCommand(mintTokens(args, pass), res)
}
