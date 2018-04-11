package features

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var txRegexp = regexp.MustCompile(`0x[0-9a-f]{64}`)

func (ctx *Context) ITransferKUSD(kcoin int, from, to string) error {
	command := fmt.Sprintf(
		`
			personal.unlockAccount(eth.coinbase, "test");
			eth.sendTransaction({from:eth.coinbase, to: "%s", value: web3.toWei(%v, 'ether')})
		`,
		ctx.accountsCoinbase[to],
		kcoin)
	res, err := ctx.cluster.Exec(ctx.accountsNodeNames[from], command)
	if err != nil {
		return err
	}
	if !txRegexp.MatchString(res.StdOut) {
		return fmt.Errorf("Expected transaction, received: %v", res.StdOut)
	}
	err = waitFor("transaction in the blockhain", 1*time.Second, 5*time.Second, func() bool {
		isInBlockchain, err := ctx.isTransactionInBlockchain(res.StdOut)
		return err == nil && isInBlockchain
	})
	if err != nil {
		return err
	}
	return nil
}

func (ctx *Context) ITryTransferKUSD(kcoin int, from, to string) error {
	command := fmt.Sprintf(
		`
			personal.unlockAccount(eth.coinbase, "test");
			eth.sendTransaction({from:eth.coinbase, to: "%s", value: web3.toWei(%v, 'ether')})
		`,
		ctx.accountsCoinbase[to],
		kcoin)
	res, err := ctx.cluster.Exec(ctx.accountsNodeNames[from], command)
	if err != nil {
		return err
	}
	ctx.lastTxStdout = res.StdOut
	return nil
}

func (ctx *Context) LastTransactionFailed() error {
	if !txRegexp.MatchString(ctx.lastTxStdout) {
		return nil // Failed at submitting the transaction, all good
	}

	isInBlockchain, err := ctx.isTransactionInBlockchain(ctx.lastTxStdout)
	if err != nil {
		return err
	}
	if isInBlockchain {
		return fmt.Errorf("the last transaction is part of the blockchain, but shouldn't")
	}
	return nil
}

func (ctx *Context) isTransactionInBlockchain(tx string) (bool, error) {
	command := fmt.Sprintf(`eth.getTransaction(%v).blockNumber`, tx)
	res, err := ctx.cluster.Exec(ctx.genesisValidatorName, command)
	if err != nil {
		return false, err
	}
	block, err := strconv.Atoi(strings.TrimSpace(res.StdOut))
	if err != nil {
		return false, err
	}
	return block > 0, nil
}
