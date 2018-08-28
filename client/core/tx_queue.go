package core

import "github.com/kowala-tech/kcoin/client/core/types"

type queueNode struct {
	tx   *types.Transaction
	next *queueNode
}

type txQueue struct {
	head  *queueNode
	tail  *queueNode
	count int
}

func newTxQueue() *txQueue {
	return new(txQueue)
}

// Len return the number of transactions in the queue.
func (q *txQueue) Len() int {
	return count
}

// Push inserts a transaction at the end of the queue.
func (q *txQueue) Push(tx *types.Transaction) {
	node := &queueNode{tx: tx}

	if q.tail == nil {
		q.tail = node
		q.head = node
	} else {
		q.tail.next = node
		q.tail = node
	}

	q.count++
}

// Pop returns the oldest transaction in the queue.
func (q *txQueue) Pop() *types.Transaction {
	if q.head == nil {
		return nil
	}

	node := q.head
	q.head = node.next
	if q.head == nil {
		q.tail = nil
	}

	q.count--

	return node.tx
}
