package validator

import (
	"errors"
	"github.com/kowala-tech/kUSD/event"
	"github.com/kowala-tech/kUSD/kusd/downloader"
)

var (
	errFailedEventReceived      = errors.New("FailedEvent received")
	errFailedToReceiveDoneEvent = errors.New("failed to receive DoneEvent")
)

// SyncWaiter subscribes to DoneEvent and FailedEvent on a event.TypeMux
// and when one of those events is received call a user function.
// Example usage to delay Voter start after block sync has finished
func SyncWaiter(eventMux *event.TypeMux) error {
	events := eventMux.Subscribe(downloader.DoneEvent{}, downloader.FailedEvent{})
	defer events.Unsubscribe()
	for ev := range events.Chan() {
		switch ev.Data.(type) {
		case downloader.DoneEvent:
			return nil
		case downloader.FailedEvent:
			return errFailedEventReceived
		}
	}

	return errFailedToReceiveDoneEvent
}
