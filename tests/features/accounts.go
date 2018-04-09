package features

import (
	"context"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/DATA-DOG/godog/gherkin"
	"github.com/kowala-tech/kcoin/cluster"
	"github.com/kowala-tech/kcoin/common"
)

type AccountEntry struct {
	AccountName string
	Funds       int64
}

func parseAccountsDataTable(accountsDataTable *gherkin.DataTable) ([]*AccountEntry, error) {
	var fields []string
	head := accountsDataTable.Rows[0].Cells
	for _, cell := range head {
		fields = append(fields, cell.Value)
	}

	var accounts []*AccountEntry

	for i := 1; i < len(accountsDataTable.Rows); i++ {
		account := &AccountEntry{}
		for n, cell := range accountsDataTable.Rows[i].Cells {
			switch head[n].Value {
			case "account":
				account.AccountName = cell.Value
			case "funds":
				parsed, err := strconv.ParseInt(cell.Value, 10, 64)
				if err != nil {
					return nil, err
				}
				account.Funds = parsed
			default:
				return nil, fmt.Errorf("unexpected column name: %s", head[n].Value)
			}
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}

func (ctx *Context) IHaveTheFollowingAccounts(accountsDataTable *gherkin.DataTable) error {
	accounts, err := parseAccountsDataTable(accountsDataTable)
	if err != nil {
		return err
	}

	// Create an archive node for each account and send them funds
	for _, account := range accounts {
		nodeName, err := ctx.cluster.RunArchiveNode()
		if err != nil {
			return err
		}

		res, err := ctx.cluster.Exec(nodeName, `eth.coinbase`)
		if err != nil {
			return err
		}
		coinbaseQuotes := res.StdOut

		ctx.accountsNodeNames[account.AccountName] = nodeName
		ctx.accountsCoinbase[account.AccountName] = strings.TrimSpace(strings.Replace(coinbaseQuotes, `"`, "", 2))

		res, err = ctx.cluster.Exec(ctx.genesisValidatorName,
			fmt.Sprintf(
				`eth.sendTransaction({
				from:eth.coinbase,
				to: %s,
				value: %v})`, coinbaseQuotes, toWei(account.Funds)))
		if err != nil {
			return err
		}
	}

	// Wait for funds to be available
	for _, account := range accounts {
		err = cluster.WaitFor(1*time.Second, 10*time.Second, func() bool {
			acct := common.HexToAddress(ctx.accountsCoinbase[account.AccountName])
			balance, err := ctx.client.BalanceAt(context.Background(), acct, nil)
			if err != nil {
				return false
			}
			expected := toWei(account.Funds)
			return balance.Cmp(expected) == 0
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (ctx *Context) TheBalanceIsExactly(account string, kcoin int64) error {
	err := cluster.WaitFor(1*time.Second, 10*time.Second, func() bool {
		acct := common.HexToAddress(ctx.accountsCoinbase[account])
		balance, err := ctx.client.BalanceAt(context.Background(), acct, nil)
		if err != nil {
			return false
		}
		return balance.Cmp(toWei(kcoin)) == 0
	})

	return err
}

func (ctx *Context) TheBalanceIsAround(account string, kcoin int64) error {
	err := cluster.WaitFor(1*time.Second, 10*time.Second, func() bool {
		acct := common.HexToAddress(ctx.accountsCoinbase[account])
		balance, err := ctx.client.BalanceAt(context.Background(), acct, nil)
		if err != nil {
			return false
		}
		diff := balance.Sub(balance, toWei(kcoin))
		diff.Abs(diff)

		return diff.Cmp(big.NewInt(100000)) < 0
	})

	return err
}
