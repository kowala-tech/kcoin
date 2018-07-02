package persistence

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/gogo/protobuf/proto"
	"github.com/kowala-tech/kcoin/client/common"
	proto2 "github.com/kowala-tech/kcoin/notifications/protocolbuffer"
)

const TxKeyPrefix = "tx:"
const TxKeyFromPrefix = "txfrom:"
const TxKeyToPrefix = "txto:"

type redisPersistence struct {
	client *redis.Client
}

func NewRedisPersistence(client *redis.Client) TransactionRepository {
	return &redisPersistence{
		client: client,
	}
}

func (p *redisPersistence) Save(tx *proto2.Transaction) error {
	enc, err := proto.Marshal(tx)
	if err != nil {
		return err
	}

	pipeline := p.client.TxPipeline()

	pipeline.Set(
		getKeyFromTx(tx),
		enc,
		0,
	)

	pipeline.SAdd(
		fmt.Sprintf("%s%s", TxKeyFromPrefix, tx.GetFrom()),
		tx.GetHash(),
	)

	pipeline.SAdd(
		fmt.Sprintf("%s%s", TxKeyToPrefix, tx.GetTo()),
		tx.GetHash(),
	)

	_, err = pipeline.Exec()

	return err
}

func (p *redisPersistence) GetTxByHash(hash common.Hash) (*proto2.Transaction, error) {
	res, err := p.client.Get(getKeyFromTxHash(hash.String())).Bytes()
	if err == redis.Nil {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	var tx proto2.Transaction
	proto.Unmarshal(res, &tx)

	return &tx, nil
}

func (p *redisPersistence) GetTxsFromAccount(account common.Address) ([]*proto2.Transaction, error) {
	var txHashesFound []common.Hash

	txsHashesFromAccount, err := p.getTransactionHashesComingFromAccount(account)
	if err != nil {
		return nil, err
	}

	txHashesFound = append(txHashesFound, txsHashesFromAccount...)

	txHashesToAccount, err := p.getTransactionHashesSentToAccount(account)
	if err != nil {
		return nil, err
	}

	txHashesFound = append(txHashesFound, txHashesToAccount...)

	txs, err := p.getTransactionsByHashes(txHashesFound)
	if err != nil {
		return nil, err
	}

	return txs, nil
}

func (p *redisPersistence) getTransactionsByHashes(hashes []common.Hash) ([]*proto2.Transaction, error) {
	var txs []*proto2.Transaction
	for _, hash := range hashes {
		tx, err := p.GetTxByHash(hash)
		if err != nil {
			return nil, err
		}
		txs = append(txs, tx)
	}

	return txs, nil
}

func (p *redisPersistence) getTransactionHashesComingFromAccount(account common.Address) ([]common.Hash, error) {
	return p.getTransactionsOfType(TxKeyFromPrefix, account)
}

func (p *redisPersistence) getTransactionHashesSentToAccount(account common.Address) ([]common.Hash, error) {
	return p.getTransactionsOfType(TxKeyToPrefix, account)
}

func (p *redisPersistence) getTransactionsOfType(typeTransaction string, account common.Address) ([]common.Hash, error) {
	resp := p.client.SMembers(
		fmt.Sprintf("%s%s", typeTransaction, account.String()),
	)

	// Non existent key
	if resp.Err() == redis.Nil {
		return nil, nil
	}

	if resp.Err() != nil {
		return nil, resp.Err()
	}

	result, err := resp.Result()
	if err != nil {
		return nil, err
	}

	var hashes []common.Hash
	for _, hashStr := range result {
		hashes = append(hashes, common.HexToHash(hashStr))
	}

	return hashes, nil
}

func getKeyFromTx(tx *proto2.Transaction) string {
	return getKeyFromTxHash(tx.GetHash())
}

func getKeyFromTxHash(hash string) string {
	return fmt.Sprintf("%s%s", TxKeyPrefix, hash)
}
