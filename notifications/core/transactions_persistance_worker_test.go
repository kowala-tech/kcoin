package core

import (
	"testing"
	"github.com/kowala-tech/kcoin/notifications/persistence/mocks"
	"github.com/yourheropaul/inj"
)

func TestWeSaveWhenWeReceiveABlock(t *testing.T) {
	mockedPersistance := mocks.TransactionRepository{}

	worker := NewTransactionsPersistanceWorker(logger)

	gr := inj.NewGraph()
	gr.Inject(worker, mockedPersistance)
}