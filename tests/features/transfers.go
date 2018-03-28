package features

import (
	"fmt"
	"regexp"
)

var txRegexp = regexp.MustCompile(`0x[0-9a-f]{64}`)

func (context *Context) ITransferKUSD(kusd int, from, to string) error {
	command := fmt.Sprintf(
		`
			personal.unlockAccount(eth.coinbase, "test");
			eth.sendTransaction({from:eth.coinbase, to: "%s", value: web3.toWei(%v, 'ether')})
		`,
		context.accountsCoinbase[to],
		kusd)
	res, err := context.cluster.Exec(context.accountsNodeNames[from], command)
	if err != nil {
		return err
	}
	if !txRegexp.MatchString(res.StdOut) {
		return fmt.Errorf("Invalid transaction response: %v", res.StdOut)
	}
	return nil
}
