package core

import "testing"

func TestWeSaveWhenWeReceiveABlock(t *testing.T) {
	worker := NewTransactionsPersistanceWorker(logger)
	worker.Start()
}