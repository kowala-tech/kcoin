package accounts

import (
	"errors"
	"github.com/kowala-tech/kUSD/common"
)

// WalletAccount wraps wallet with a "default" account
type WalletAccount interface {
	Wallet
	Account() Account
}

type walletAccount struct {
	wallet  Wallet
	account Account
}

var (
	errInvalidAccountAddress = errors.New("invalid account address, doesnt exists in wallet")
)

// NewWalletAccount ensure that the address provided exists in the wallet
func NewWalletAccount(wallet Wallet, accountAddress common.Address) (*walletAccount, error) {
	account := Account{Address: accountAddress}
	if !wallet.Contains(account) {
		return nil, errInvalidAccountAddress
	}
	return &walletAccount{wallet: wallet, account: account}, nil
}

func (account *walletAccount) Account() Account {
	return account.account
}
