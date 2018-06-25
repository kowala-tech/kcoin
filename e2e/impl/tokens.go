package impl

import (
	"math/big"

	"github.com/kowala-tech/kcoin/e2e/cluster"
	"github.com/kowala-tech/kcoin/client/accounts"
	"github.com/kowala-tech/kcoin/client/knode"
	"github.com/kowala-tech/kcoin/client/common/hexutil"
)

func (ctx *ValidationContext) sendTokensAndWait(from, to accounts.Account, tokens int) error {
	return ctx.waiter.Do(
		func() error {
			var err error
			err = ctx.sendTokens(from, to, tokens)
			return err
		},
		func() error {
			return ctx.checkTokenBalance(to, tokens)
		})
}

func (ctx *ValidationContext) sendTokens(from, to accounts.Account, tokens int) error {
	bigPointer := big.NewInt(int64(tokens))
	hexBig := hexutil.Big(*bigPointer)
	args := knode.TransferArgs{
		From: from.Address,
		To: &to.Address,
		Value: &hexBig,
	}

	res := &cluster.ExecResponse{}
	return ctx.execCommand(transferTokens(args), res)
}
