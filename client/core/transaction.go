package core

import (
	"container/heap"
	"time"

	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/core/types"
)

// Transaction represents a tx verified by the tx pool
type Transaction struct {
	AddedAt time.Time
	*types.Transaction
}

// Transactions represents a set of txs verified by the tx pool
type Transactions []*Transaction

// TxByNonce implements the sort interface to allow sorting a list of transactions
// by their nonces. This is usually only useful for sorting transactions from a
// single account, otherwise a nonce comparison doesn't make much sense.
type TxByNonce Transactions

func (s TxByNonce) Len() int { return len(s) }
func (s TxByNonce) Less(i, j int) bool {
	return s[i].Nonce() < s[j].Nonce()
}
func (s TxByNonce) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// TxByTimestamp implements the heap interface, making it useful for all
// at once sortingas well as individually adding and removing elements.
type TxByTimestamp Transactions

func (s TxByTimestamp) Len() int           { return len(s) }
func (s TxByTimestamp) Less(i, j int) bool { return s[i].AddedAt.After(s[j].AddedAt) }
func (s TxByTimestamp) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func (s *TxByTimestamp) Push(x interface{}) {
	*s = append(*s, x.(*Transaction))
}

func (s *TxByTimestamp) Pop() interface{} {
	old := *s
	n := len(old)
	x := old[n-1]
	*s = old[0 : n-1]
	return x
}

// TransactionsByTimestampAndNonce represents a set of transactions that can return
// transactions in a timestamp and nonce-honouring order.
type TransactionsByTimestampAndNonce struct {
	txs    map[common.Address]Transactions // per account nonce-sorted list of transactions
	heads  TxByTimestamp                   // Next transaction for each unique account (timestamp heap)
	signer types.Signer                    // signer for the set of transactions
}

// NewTransactionsByTimestampAndNonce creates a transaction set that can retrieve
// timestamp sorted transactions in a nonce-honouring way.
func NewTransactionsByTimestampAndNonce(signer types.Signer, txs map[common.Address]Transactions) *TransactionsByTimestampAndNonce {
	heads := make(TxByTimestamp, 0, len(txs))
	for from, accTxs := range txs {
		heads = append(heads, accTxs[0])
		// Ensure the sender address is from the signer
		acc, _ := types.TxSender(signer, accTxs[0].Transaction)
		txs[acc] = accTxs[1:]
		if from != acc {
			delete(txs, from)
		}
	}
	heap.Init(&heads)

	return &TransactionsByTimestampAndNonce{
		txs:    txs,
		heads:  heads,
		signer: signer,
	}
}

// Peek returns the next transaction by price.
func (t *TransactionsByTimestampAndNonce) Peek() *types.Transaction {
	if len(t.heads) == 0 {
		return nil
	}
	return t.heads[0].Transaction
}

// Shift replaces the current best head with the next one from the same account.
func (t *TransactionsByTimestampAndNonce) Shift() {
	acc, _ := types.TxSender(t.signer, t.heads[0].Transaction)
	if txs, ok := t.txs[acc]; ok && len(txs) > 0 {
		t.heads[0], t.txs[acc] = txs[0], txs[1:]
		heap.Fix(&t.heads, 0)
	} else {
		heap.Pop(&t.heads)
	}
}

// Pop removes the best transaction, *not* replacing it with the next one from
// the same account. This should be used when a transaction cannot be executed
// and hence all subsequent ones should be discarded from the same account.
func (t *TransactionsByTimestampAndNonce) Pop() {
	heap.Pop(&t.heads)
}
