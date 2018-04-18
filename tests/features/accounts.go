package features

import (
	"context"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/DATA-DOG/godog/gherkin"
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
	accountsData, err := parseAccountsDataTable(accountsDataTable)
	if err != nil {
		return err
	}

	for _, accountData := range accountsData {
		if err := ctx.createAccountWithFunds(accountData.AccountName, accountData.Funds); err != nil {
			return err
		}
	}

	return nil
}

func (ctx *Context) TheBalanceIsExactly(accountName string, kcoin int64) error {
	expected := toWei(kcoin)

	account := ctx.accounts[accountName]
	balance, err := ctx.client.BalanceAt(context.Background(), account.Address, nil)
	if err != nil {
		return err
	}
	if balance.Cmp(expected) != 0 {
		return fmt.Errorf("Balance expected to be %v but is %v", expected, balance)
	}
	return nil
}

func (ctx *Context) TheBalanceIsAround(accountName string, kcoin int64) error {
	expected := toWei(kcoin)

	account := ctx.accounts[accountName]
	balance, err := ctx.client.BalanceAt(context.Background(), account.Address, nil)
	if err != nil {
		return err
	}
	diff := &big.Int{}
	diff.Sub(balance, expected)
	diff.Abs(diff)

	if diff.Cmp(big.NewInt(100000)) >= 0 {
		return fmt.Errorf("Balance expected to be around %v but is %v", expected, balance)
	}
	return nil
}

func (ctx *Context) createAccountWithFunds(accountName string, kcoins int64) error {
	if _, ok := ctx.accounts[accountName]; ok {
		return fmt.Errorf("an account with this name already exists: %s", accountName)
	}
	account, err := ctx.AccountsStorage.NewAccount("test")
	if err != nil {
		return err
	}
	if err := ctx.AccountsStorage.Unlock(account, "test"); err != nil {
		return err
	}

	ctx.accounts[accountName] = account
	if _, err := ctx.sendFunds(ctx.seederAccount, account, kcoins); err != nil {
		return err
	}
	err = waitFor("account receives the balance", 1*time.Second, 10*time.Second, func() bool {
		account := ctx.accounts[accountName]
		balance, err := ctx.client.BalanceAt(context.Background(), account.Address, nil)
		if err != nil {
			return false
		}
		return balance.Cmp(toWei(kcoins)) == 0
	})
	if err != nil {
		return err
	}
	return nil
}
