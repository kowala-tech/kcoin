package validator

import (
	"context"
	"errors"

	"github.com/kowala-tech/kcoin/client/event"
	"github.com/kowala-tech/kcoin/client/knode/downloader"
	"github.com/kowala-tech/kcoin/client/log"
)

var (
	errFailedEventReceived      = errors.New("FailedEvent received")
	errFailedToReceiveDoneEvent = errors.New("failed to receive DoneEvent")
)

// SyncWaiter subscribes to DoneEvent and FailedEvent on a event.TypeMux
// and when one of those events is received call a user function.
// Example usage to delay Validator start after block sync has finished
func SyncWaiter(ctx context.Context, eventMux *event.TypeMux) error {
	events := eventMux.Subscribe(downloader.DoneEvent{}, downloader.FailedEvent{})
	defer events.Unsubscribe()

	log.Warn("!!!!! SyncWaiter started")
	defer func() {
		log.Warn("!!!!! SyncWaiter stopped")
	}()

	for {
		select {
		case ev, ok := <-events.Chan():
			if !ok {
				return errFailedToReceiveDoneEvent
			}

			switch ev.Data.(type) {
			case downloader.DoneEvent:
				log.Info("sync finished in SyncWaiter")
				return nil
			case downloader.FailedEvent:
				log.Info("failed to sync while SyncWaiter", "err", ev.Data.(downloader.FailedEvent).Err)
				return errFailedEventReceived
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	return nil
}
