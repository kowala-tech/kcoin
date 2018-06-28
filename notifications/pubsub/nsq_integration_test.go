// +build integration

package pubsub

import (
	"testing"
	"time"

	"github.com/kowala-tech/kcoin/notifications/environment"
	"github.com/stretchr/testify/require"
)

func TestNSQ_CanPublishAndSubscribe(t *testing.T) {
	data := []byte("Hello world")
	topic := "test-topic#ephemeral"
	channel := "test-channel"

	envReader := environment.NewReaderOs()
	nsqAddr := envReader.Read("NSQ_ADDR")

	publisher, err := NewNSQPublisher(nsqAddr, logger)
	require.Nil(t, err)
	defer publisher.Stop()

	subscriber := NewNSQSubscriber(topic, channel, nsqAddr, logger)
	defer subscriber.Stop()

	receivedCh := make(chan []byte)
	handler := &MessageHandlerMock{
		HandleMessageFunc: func(topic string, data []byte) error {
			receivedCh <- data
			return nil
		},
	}

	subscriber.AddHandler(handler)
	err = subscriber.Start()
	require.Nil(t, err)

	err = publisher.Publish(topic, data)
	require.Nil(t, err)

	select {
	case received := <-receivedCh:
		require.Equal(t, data, received)
	case <-time.After(2 * time.Second):
		t.Error("Timeout")
	}
}

func TestNSQ_CanStopSubscription(t *testing.T) {
	data := []byte("Hello world")
	topic := "test-topic#ephemeral"
	channel := "test-channel"

	envReader := environment.NewReaderOs()
	nsqAddr := envReader.Read("NSQ_ADDR")

	publisher, err := NewNSQPublisher(nsqAddr, logger)
	require.Nil(t, err)
	defer publisher.Stop()

	subscriber := NewNSQSubscriber(topic, channel, nsqAddr, logger)

	receivedCh := make(chan []byte)
	handler := &MessageHandlerMock{
		HandleMessageFunc: func(topic string, data []byte) error {
			receivedCh <- data
			return nil
		},
	}

	subscriber.AddHandler(handler)
	err = subscriber.Start()
	require.Nil(t, err)

	err = publisher.Publish(topic, data)
	require.Nil(t, err)

	select {
	case received := <-receivedCh:
		require.Equal(t, data, received)
	case <-time.After(2 * time.Second):
		t.Error("Timeout")
	}

	subscriber.Stop()
	err = publisher.Publish(topic, data)
	require.Nil(t, err)

	select {
	case received := <-receivedCh:
		require.Nil(t, received)
	case <-time.After(50 * time.Millisecond):
		// Good
	}
}
