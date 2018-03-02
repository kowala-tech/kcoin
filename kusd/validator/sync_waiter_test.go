package validator

import (
	"github.com/kowala-tech/kUSD/event"
	"github.com/kowala-tech/kUSD/kusd/downloader"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSyncWaiterNoError(t *testing.T) {
	mux := new(event.TypeMux)
	defer mux.Stop()

	go func() {
		time.Sleep(time.Millisecond * 10)
		err := mux.Post(downloader.DoneEvent{})
		assert.NoError(t, err, "error posting event")
	}()

	err := SyncWaiter(mux)
	assert.NoError(t, err, "error from SyncWaiter")
}

func TestSyncWaiterReturnErrorOnFailedEvent(t *testing.T) {
	mux := new(event.TypeMux)
	defer mux.Stop()

	go func() {
		time.Sleep(time.Millisecond * 10)
		err := mux.Post(downloader.FailedEvent{})
		assert.NoError(t, err, "error posting event")
	}()

	err := SyncWaiter(mux)
	assert.Error(t, err, "error from SyncWaiter")
}

func TestSyncWaiterReturnsErrorOnClosedMutex(t *testing.T) {
	mux := new(event.TypeMux)

	go func() {
		time.Sleep(time.Millisecond * 10)
		mux.Stop()
	}()

	err := SyncWaiter(mux)
	assert.Error(t, err, "failed to receive DoneEvent")
}
