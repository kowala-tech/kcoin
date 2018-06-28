package core

import (
	"testing"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/kowala-tech/kcoin/notifications/keyvalue"
	"github.com/stretchr/testify/require"
	"github.com/yourheropaul/inj"

	"github.com/kowala-tech/kcoin/notifications/notifier"
	"github.com/kowala-tech/kcoin/notifications/protocolbuffer"
	"github.com/kowala-tech/kcoin/notifications/pubsub"
)

func setup_emailer(t *testing.T) (*Emailer, *notifier.NotifierMock, *pubsub.SubscriberMock, *keyvalue.KeyValueMock) {
	emailer := NewEmailer(logger, "from@test.com")
	notif := &notifier.NotifierMock{
		SendFunc: func(vars map[string]string) error {
			return nil
		},
	}
	subs := &pubsub.SubscriberMock{
		StartFunc: func() error {
			return nil
		},
		StopFunc: func() {
		},
	}
	kv := &keyvalue.KeyValueMock{}

	gr := inj.NewGraph()
	gr.Provide(
		emailer,
		notif,
		subs,
		kv,
	)

	valid, messages := gr.Assert()
	require.True(t, valid, messages)

	return emailer, notif, subs, kv
}

func TestEmailer_SendsEmailToRegisteredEmails(t *testing.T) {
	emailer, notif, subs, kv := setup_emailer(t)
	address := "0xabcd"

	kv.GetStringFunc = func(key string) (string, error) {
		return "to@test.com", nil
	}

	var handler pubsub.MessageHandler
	subs.AddHandlerFunc = func(in1 pubsub.MessageHandler) {
		handler = in1
	}

	emailer.Register()
	time.Sleep(10 * time.Millisecond)

	require.NotNil(t, handler)

	tx := &protocolbuffer.Transaction{
		Amount: 42,
		To:     address,
	}
	data, err := proto.Marshal(tx)
	require.NoError(t, err)
	handler.HandleMessage("transactions", data)

	require.Len(t, kv.GetStringCalls(), 1)
	require.Equal(t, kv.GetStringCalls()[0].Key, address)
	require.Len(t, notif.SendCalls(), 1)
	require.Equal(t, "to@test.com", notif.SendCalls()[0].Vars[notifier.EmailToKey])

}

func TestEmailer_DoesNotSendEmailsToNonRegisteredWallets(t *testing.T) {
	emailer, notif, subs, kv := setup_emailer(t)
	address := "0xabcd"

	kv.GetStringFunc = func(key string) (string, error) {
		return "", nil
	}

	var handler pubsub.MessageHandler
	subs.AddHandlerFunc = func(in1 pubsub.MessageHandler) {
		handler = in1
	}

	emailer.Register()
	time.Sleep(10 * time.Millisecond)

	require.NotNil(t, handler)

	tx := &protocolbuffer.Transaction{
		Amount: 42,
		To:     address,
	}
	data, err := proto.Marshal(tx)
	require.NoError(t, err)
	handler.HandleMessage("transactions", data)

	require.Len(t, kv.GetStringCalls(), 1)
	require.Equal(t, kv.GetStringCalls()[0].Key, address)
	require.Len(t, notif.SendCalls(), 0)
}
