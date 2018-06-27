package core

import (
	"math/big"
	"testing"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/require"
	"github.com/yourheropaul/inj"

	"github.com/kowala-tech/kcoin/notifications/blockchain"
	"github.com/kowala-tech/kcoin/notifications/keyvalue"
	"github.com/kowala-tech/kcoin/notifications/protocolbuffer"
	"github.com/kowala-tech/kcoin/notifications/pubsub"
)

func setup_transactions_publisher(t *testing.T) (*TransactionsPublisher, *blockchain.BlockchainMock, *pubsub.PublisherMock, *keyvalue.ValueMock, chan *protocolbuffer.Transaction) {
	txChn := make(chan *protocolbuffer.Transaction)

	mockedBlockchain := &blockchain.BlockchainMock{
		OnBlockFunc: func(in1 blockchain.BlockHandler) error {
			go func() {
				for tx := range txChn {
					in1.HandleBlock(&blockchain.Block{
						Number:       big.NewInt(1),
						Transactions: []*protocolbuffer.Transaction{tx},
					})
				}
			}()
			return nil
		},
		SeekFunc: func(in1 *big.Int) error {
			return nil
		},
		StartFunc: func() error {
			return nil
		},
		StopFunc: func() {
		},
	}

	mockedPublisher := &pubsub.PublisherMock{
		PublishFunc: func(topic string, data []byte) error {
			return nil
		},
		StopFunc: func() {
		},
	}

	mockedValueStorage := &keyvalue.ValueMock{
		GetInt64Func: func() (int64, error) {
			return 0, nil
		},
		GetStringFunc: func() (string, error) {
			return "", nil
		},
		PutInt64Func: func(value int64) error {
			return nil
		},
		PutStringFunc: func(value string) error {
			return nil
		},
	}

	tp := NewTransactionsPublisher(logger)

	gr := inj.NewGraph()
	gr.Provide(
		tp,
		mockedBlockchain,
		mockedPublisher,
		mockedValueStorage,
	)

	valid, messages := gr.Assert()
	require.True(t, valid, messages)

	return tp, mockedBlockchain, mockedPublisher, mockedValueStorage, txChn
}

func TestTransactionsPublisher_SeeksToStoredValue(t *testing.T) {
	tp, mockedBlockchain, _, mockedValueStorage, _ := setup_transactions_publisher(t)

	seeked := make(chan bool)
	mockedValueStorage.GetInt64Func = func() (int64, error) {
		return 42, nil
	}
	mockedBlockchain.SeekFunc = func(in1 *big.Int) error {
		seeked <- true
		return nil
	}

	go tp.Start()
	defer tp.Stop()

	select {
	case <-seeked:
	case <-time.After(time.Second):
		t.Fatal("Timeout (1s). Transaction stream didn't seek.")
	}

	calls := mockedBlockchain.SeekCalls()
	require.Equal(t, len(calls), 1)
	require.Equal(t, calls[0].In1.Int64(), int64(42))
}

func TestTransactionsPublisher_DoesNotSeekIfThereIsNoValue(t *testing.T) {
	tp, mockedBlockchain, _, _, _ := setup_transactions_publisher(t)

	go tp.Start()
	defer tp.Stop()
	time.Sleep(10 * time.Millisecond) // Give some time to run through the seeking code

	calls := mockedBlockchain.SeekCalls()
	require.Equal(t, len(calls), 0)
}

func TestTransactionsPublisher_UpdatesValueStore(t *testing.T) {
	updated := make(chan bool)
	tp, _, _, mockedValueStorage, txChn := setup_transactions_publisher(t)

	mockedValueStorage.PutInt64Func = func(value int64) error {
		updated <- true
		return nil
	}

	go tp.Start()
	defer tp.Stop()

	transaction := &protocolbuffer.Transaction{
		To:     "abc",
		Amount: 42,
	}

	select {
	case txChn <- transaction:
	case <-time.After(time.Second):
		t.Fatal("Timeout (1s). Transaction not processed.")
	}

	select {
	case <-updated:
	case <-time.After(time.Second):
		t.Fatal("Timeout (1s). Value store not updated.")
	}

	calls := mockedValueStorage.PutInt64Calls()
	require.Equal(t, len(calls), 1)
}

func TestTransactionsPublisher_PublishesAllTransactions(t *testing.T) {
	published := make(chan bool)

	tp, _, mockedPublisher, _, txChn := setup_transactions_publisher(t)
	mockedPublisher.PublishFunc = func(topic string, data []byte) error {
		published <- true
		return nil
	}

	go tp.Start()
	defer tp.Stop()

	transaction := &protocolbuffer.Transaction{
		To:     "abc",
		Amount: 42,
	}

	select {
	case txChn <- transaction:
	case <-time.After(time.Second):
		t.Fatal("Timeout (1s). Transaction not processed.")
	}

	select {
	case <-published:
	case <-time.After(time.Second):
		t.Fatal("Timeout (1s). Transaction not published.")
	}

	calls := mockedPublisher.PublishCalls()
	require.Equal(t, len(calls), 1)
	require.Equal(t, calls[0].Topic, "transactions")

	data, err := proto.Marshal(transaction)
	require.NoError(t, err)
	require.Equal(t, calls[0].Data, data)
}
