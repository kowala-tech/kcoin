package accounts

import (
	"github.com/kowala-tech/kUSD/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewWalletAccount(t *testing.T) {
	address := common.Address{1}
	wallet := &MockWallet{}
	wallet.On("Contains", Account{Address: address}).Return(true)

	walletAccount, err := NewWalletAccount(wallet, address)
	assert.NoError(t, err)

	assert.Equal(t, Account{Address: address}, walletAccount.account)
}

func TestNewWalletAccountFailsIfAddressDoesntExistInWallet(t *testing.T) {
	address := common.Address{}
	wallet := &MockWallet{}
	wallet.On("Contains", Account{Address: address}).Return(false)

	_, err := NewWalletAccount(wallet, address)
	assert.Error(t, err, "invalid account address, doesnt exists in wallet")
}
