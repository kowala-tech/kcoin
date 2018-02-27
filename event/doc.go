/*
Package event deals with subscriptions to real-time events.

Receivers can be registered to handle events of certain type.
Example: Event that is dispatched when a block is inserted in the chain.

Subscription

Subscription represents a stream of events. The carrier of the events is
typically a channel, but isn't part of the interface.
A new subscription runs a producer function as a subscription in a new
goroutine. The channel given to the producer is closed when Unsubscribe is
called.

Subscription example: example_subscription_test.go

Feed

Feed implements one-to-many subscriptions where the carrier of events is a
channel. Values sent to a Feed are delivered to all subscribed channels
simultaneously.
Feeds can only be used with a single type. The type is determined by the first
Send or Subscribe operation. Subsequent calls to these methods panic if the type
does not match. The zero value Feed is ready to be used.

The transaction pool implements a feed to let the consumers know when a
transaction enters the pool(TxPreEvent).

Feed example: example_feed_test.go
*/

package event
