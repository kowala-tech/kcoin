package types_test

import (
	"io/ioutil"
	"testing"

	"math/big"

	"github.com/kowala-tech/kcoin/accounts/keystore"
	"github.com/kowala-tech/kcoin/common"
	"github.com/kowala-tech/kcoin/core/types"
	"github.com/stretchr/testify/assert"
)

func TestFrom(t *testing.T) {
	tmpdir, _ := ioutil.TempDir("", "eth-keystore-test")
	accountsStorage := keystore.NewKeyStore(tmpdir, 2, 1)

	account, err := accountsStorage.NewAccount("test")
	if err != nil {
		t.Errorf("Error creating account. %s", err)
	}

	err = accountsStorage.Unlock(account, "test")
	if err != nil {
		t.Fatalf("Error unlocking account. %s", err)
	}

	tx := types.NewTransaction(
		0,
		common.HexToAddress("0xecf8f87f810ecf450940c9f60066b4a7a501d6a7"),
		big.NewInt(10),
		big.NewInt(10),
		big.NewInt(1),
		nil,
	)

	signedTx, err := accountsStorage.SignTx(account, tx, big.NewInt(1))
	if err != nil {
		t.Errorf("Error signing transaction. %s", err)
	}

	fromAddr, err := signedTx.From()
	if err != nil {
		t.Errorf("Error getting from address from signed tx. %s", err)

	}

	assert.Equal(t, &account.Address, fromAddr)
}
