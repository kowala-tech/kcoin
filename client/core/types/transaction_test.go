package types_test

import (
	"io/ioutil"
	"testing"

	"math/big"

	"github.com/kowala-tech/kcoin/client/accounts/keystore"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/core/types"
	"github.com/stretchr/testify/require"
)

func TestFrom(t *testing.T) {
	tmpdir, _ := ioutil.TempDir("", "eth-keystore-test")
	accountsStorage := keystore.NewKeyStore(tmpdir, 2, 1)

	account, err := accountsStorage.NewAccount("test")
	require.NoError(t, err, "Error creating account. %s", err)

	err = accountsStorage.Unlock(account, "test")
	require.NoError(t, err, "Error unlocking account. %s", err)

	tx := types.NewTransaction(
		0,
		common.HexToAddress("0xecf8f87f810ecf450940c9f60066b4a7a501d6a7"),
		big.NewInt(10),
		10,
		nil,
	)

	signedTx, err := accountsStorage.SignTx(account, tx, big.NewInt(1))
	require.NoError(t, err, "Error signing transaction. %s", err)

	fromAddr, err := signedTx.From()
	require.NoError(t, err, "Error getting from address from signed tx. %s", err)

	require.Equal(t, &account.Address, fromAddr)
}
