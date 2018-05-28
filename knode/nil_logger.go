package knode

import (
	"context"

	"github.com/kowala-tech/kcoin"
	"github.com/kowala-tech/kcoin/core/types"
)

type NilLogFilter struct {
	nilLogger
}

func (f NilLogFilter) FilterLogs(ctx context.Context, query kowala.FilterQuery) ([]types.Log, error) {
	return []types.Log{}, nil
}

func (f NilLogFilter) SubscribeFilterLogs(ctx context.Context, query kowala.FilterQuery, ch chan<- types.Log) (kowala.Subscription, error) {
	return f.nilLogger, nil
}

type nilLogger <-chan error

func newNilLogger() nilLogger {
	return make(<-chan error, 1)
}

// Unsubscribe cancels the sending of events to the data channel
// and closes the error channel.
func (l nilLogger) Unsubscribe() {}

// Err returns the subscription error channel. The error channel receives
// a value if there is an issue with the subscription (e.g. the network connection
// delivering the events has been closed). Only one value will ever be sent.
// The error channel is closed by Unsubscribe.
func (l nilLogger) Err() <-chan error {
	return l
}
