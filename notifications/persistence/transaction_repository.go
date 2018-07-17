package persistence

import (
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/notifications/protocolbuffer"
)

//TransactionRepository is a repository that persist transactions.
type TransactionRepository interface {
	Save(tx *protocolbuffer.Transaction) error
	GetTxByHash(hash common.Hash) (*protocolbuffer.Transaction, error)
	GetTxsFromAccount(address common.Address) ([]*protocolbuffer.Transaction, error)
}
