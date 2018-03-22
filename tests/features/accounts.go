package features

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/DATA-DOG/godog/gherkin"
	"github.com/kowala-tech/kUSD/cluster"
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

func (context *Context) IHaveTheFollowingAccounts(accountsDataTable *gherkin.DataTable) error {
	accounts, err := parseAccountsDataTable(accountsDataTable)
	if err != nil {
		return err
	}

	// Create an archive node for each account and send them funds
	for _, account := range accounts {
		nodeName, err := context.cluster.RunArchiveNode()
		if err != nil {
			return err
		}

		res, err := context.cluster.Exec(nodeName, `eth.coinbase`)
		if err != nil {
			return err
		}
		coinbaseQuotes := res.StdOut

		context.accountsNodeNames[account.AccountName] = nodeName
		context.accountsCoinbase[account.AccountName] = strings.TrimSpace(strings.Replace(coinbaseQuotes, `"`, "", 2))

		res, err = context.cluster.Exec(context.genesisValidatorName,
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
			balance, err := context.cluster.GetBalance(context.accountsNodeNames[account.AccountName])
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

func (context *Context) TheBalanceIsExactly(account string, kusd int64) error {
	err := cluster.WaitFor(1*time.Second, 10*time.Second, func() bool {
		balance, err := context.cluster.GetBalance(context.accountsNodeNames[account])
		if err != nil {
			return false
		}
		return balance.Cmp(toWei(kusd)) == 0
	})

	return err
}

func (context *Context) TheBalanceIsAround(account string, kusd int64) error {
	err := cluster.WaitFor(1*time.Second, 10*time.Second, func() bool {
		balance, err := context.cluster.GetBalance(context.accountsNodeNames[account])
		if err != nil {
			return false
		}
		diff := balance.Sub(balance, toWei(kusd))
		diff.Abs(diff)

		return diff.Cmp(big.NewInt(100000)) < 0
	})

	return err
}
