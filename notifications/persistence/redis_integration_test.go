// +build integration

package persistence

import (
	"testing"

	"math/big"

	"time"

	"github.com/go-redis/redis"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/notifications/environment"
	"github.com/kowala-tech/kcoin/notifications/protocolbuffer"
	"github.com/stretchr/testify/assert"
)

const RedisServerEnvKey = "REDIS_ADDR"

func TestSaveAndGetTransaction(t *testing.T) {
	p := redisPersistence{
		client: getRedisClient(t),
	}

	hash := common.HexToHash("0x4e197959672274721d4d6565ae60bc54a97092c818612823d105a981122e09a5")
	address := common.HexToAddress("0xdbdfdbce9a34c3ac5546657f651146d88d1b639a")
	to := common.HexToAddress("0xdbdfdbce9a34c3ac5546657f651146d88d1b63bb")
	amount := big.NewInt(12345)

	tx := &protocolbuffer.Transaction{
		Hash:        hash.String(),
		Amount:      amount.Int64(),
		From:        address.String(),
		To:          to.String(),
		GasUsed:     1000,
		GasPrice:    2000,
		BlockHeight: 1050,
		Timestamp:   time.Now().Unix(),
	}

	t.Run("Get a unexisting transaction", func(t *testing.T) {
		tx, err := p.GetTxByHash(hash)
		if err != nil {
			t.Fatalf("Error getting transaction: %s", err)
		}

		assert.Nil(t, tx)
	})

	t.Run("Save a transaction", func(t *testing.T) {
		err := p.Save(tx)
		assert.NoError(t, err)
	})

	t.Run("Get the saved transaction", func(t *testing.T) {
		savedTx, err := p.GetTxByHash(hash)
		if err != nil {
			t.Fatalf("Error getting transaction: %s", err)
		}

		assert.Equal(t, tx, savedTx)
	})

	// Teardown
	assert.NoError(t, p.client.FlushAll().Err())
}

func TestGetTransactionsFromAccount(t *testing.T) {
	p := redisPersistence{
		client: getRedisClient(t),
	}

	targetAccount := common.HexToAddress("0xdbdfdbce9a34c3ac5546657f651146d88d1b639a")

	t.Run("Get transactions from account with no transactions", func(t *testing.T) {
		var expectedTransactions []*protocolbuffer.Transaction

		txs, err := p.GetTxsFromAccount(targetAccount)
		if err != nil {
			t.Fatalf("Error getting transactions by account: %s", err)
		}

		assert.Equal(t, expectedTransactions, txs)
	})

	hash := common.HexToHash("0x4e197959672274721d4d6565ae60bc54a97092c818612823d105a981122e09a5")
	account := common.HexToAddress("0xdbdfdbce9a34c3ac5546657f651146d88d1bcaca")
	amount := big.NewInt(12345)

	fromAccountTransaction := &protocolbuffer.Transaction{
		Hash:        hash.String(),
		Amount:      amount.Int64(),
		From:        targetAccount.String(),
		To:          account.String(),
		GasUsed:     1000,
		GasPrice:    2000,
		BlockHeight: 1050,
		Timestamp:   time.Now().Unix(),
	}

	t.Run("Get transactions with account with one from transaction", func(t *testing.T) {

		p.Save(fromAccountTransaction)

		expectedTransactions := []*protocolbuffer.Transaction{
			fromAccountTransaction,
		}

		txs, err := p.GetTxsFromAccount(targetAccount)
		if err != nil {
			t.Fatalf("Error getting transactions from account: %s", err)
		}

		assert.Equal(t, expectedTransactions, txs)
	})

	toHash := common.HexToHash("0x4e197959672274721d4d6565ae60bc54a97092c818612823d105a981122e0808")
	account2 := common.HexToAddress("0xdbdfdbce9a34c3ac5546657f651146d88d1bcaca")

	toAccountTransaction := &protocolbuffer.Transaction{
		Hash:        toHash.String(),
		Amount:      amount.Int64(),
		From:        account2.String(),
		To:          targetAccount.String(),
		GasUsed:     1000,
		GasPrice:    2000,
		BlockHeight: 1050,
		Timestamp:   time.Now().Unix(),
	}

	t.Run("Get transactions with 1 from 1 to", func(t *testing.T) {
		p.Save(toAccountTransaction)

		expectedTransactions := []*protocolbuffer.Transaction{
			fromAccountTransaction,
			toAccountTransaction,
		}

		txs, err := p.GetTxsFromAccount(targetAccount)
		if err != nil {
			t.Fatalf("Error getting transactions: %s", err)
		}

		assert.Equal(t, expectedTransactions, txs)
	})

	// Teardown
	assert.NoError(t, p.client.FlushAll().Err())
}

func getRedisClient(t *testing.T) *redis.Client {
	envReader := environment.NewReaderOs()
	redisAddr := envReader.Read(RedisServerEnvKey)
	if redisAddr == "" {
		t.Fatalf("Cannot execute tests with empty redis addr.")
	}

	client := redis.NewClient(
		&redis.Options{
			Addr: redisAddr,
		},
	)

	assert.NoError(t, client.Ping().Err())

	return client
}
