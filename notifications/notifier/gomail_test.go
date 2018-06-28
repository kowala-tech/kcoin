package notifier

import (
	"bytes"
	"errors"
	"html/template"
	"io"
	"testing"

	"github.com/go-mail/mail"
	"github.com/stretchr/testify/require"
)

func getDialer(sender mail.SendCloser) *DialerMock {
	return &DialerMock{
		DialFunc: func() (mail.SendCloser, error) {
			return sender, nil
		},
	}
}

type sendCloserMock struct {
	sendFunc func(from string, to []string, msg io.WriterTo) error
}

func (sc *sendCloserMock) Send(from string, to []string, msg io.WriterTo) error {
	return sc.sendFunc(from, to, msg)
}

func (sc *sendCloserMock) Close() error {
	return nil
}

func getSendCloser(fn func(from string, to []string, msg io.WriterTo) error) *sendCloserMock {
	return &sendCloserMock{
		sendFunc: fn,
	}
}

func TestGomail_SendsNewTransferEmails(t *testing.T) {
	var receivedFrom string
	var receivedTo string
	var receivedMsg string

	sender := getSendCloser(func(from string, to []string, msg io.WriterTo) error {
		receivedFrom = from
		receivedTo = to[0]
		var b bytes.Buffer
		msg.WriteTo(&b)
		receivedMsg = b.String()
		return nil
	})

	notifier, err := NewGomailNewTransfer(logger, getDialer(sender))
	require.NoError(t, err)

	err = notifier.Send(map[string]string{
		EmailFromKey: "from@test.com",
		EmailToKey:   "to@test.com",
	})
	require.NoError(t, err)

	require.Equal(t, "from@test.com", receivedFrom)
	require.Equal(t, "to@test.com", receivedTo)
	require.Contains(t, receivedMsg, "received a transaction") // Some text in the template
	require.Contains(t, receivedMsg, "You've received kUSD!")  // Subject
}
func TestGomail_SendsEmails(t *testing.T) {
	var receivedFrom string
	var receivedTo string
	var receivedMsg string

	sender := getSendCloser(func(from string, to []string, msg io.WriterTo) error {
		receivedFrom = from
		receivedTo = to[0]
		var b bytes.Buffer
		msg.WriteTo(&b)
		receivedMsg = b.String()
		return nil
	})

	tpl, err := template.New("test").Parse("Hello {{.VAR1}}")
	require.NoError(t, err)

	notifier := NewGomailTemplate(logger, getDialer(sender), "Test email", tpl)
	err = notifier.Send(map[string]string{
		EmailFromKey: "from@test.com",
		EmailToKey:   "to@test.com",
		"VAR1":       "World",
	})
	require.NoError(t, err)

	require.Equal(t, "from@test.com", receivedFrom)
	require.Equal(t, "to@test.com", receivedTo)
	require.Contains(t, receivedMsg, "Hello World")
}

func TestGomail_RequiresFromVar(t *testing.T) {
	sender := getSendCloser(func(from string, to []string, msg io.WriterTo) error {
		return nil
	})

	tpl, err := template.New("test").Parse("Hello {{.VAR1}}")
	require.NoError(t, err)

	notifier := NewGomailTemplate(logger, getDialer(sender), "Test email", tpl)
	err = notifier.Send(map[string]string{
		// EmailFromKey: "from@test.com",
		EmailToKey: "to@test.com",
		"VAR1":     "World",
	})
	require.Error(t, err)
}

func TestGomail_RequiresToVar(t *testing.T) {
	sender := getSendCloser(func(from string, to []string, msg io.WriterTo) error {
		return nil
	})

	tpl, err := template.New("test").Parse("Hello {{.VAR1}}")
	require.NoError(t, err)

	notifier := NewGomailTemplate(logger, getDialer(sender), "Test email", tpl)
	err = notifier.Send(map[string]string{
		EmailFromKey: "from@test.com",
		// EmailToKey:   "to@test.com",
		"VAR1": "World",
	})
	require.Error(t, err)
}

func TestGomail_RequiresTemplateVars(t *testing.T) {
	sender := getSendCloser(func(from string, to []string, msg io.WriterTo) error {
		return nil
	})

	tpl, err := template.New("test").Parse("Hello {{.VAR1}}")
	require.NoError(t, err)

	notifier := NewGomailTemplate(logger, getDialer(sender), "Test email", tpl)
	err = notifier.Send(map[string]string{
		EmailFromKey: "from@test.com",
		EmailToKey:   "to@test.com",
		// "VAR1": "World",
	})
	require.Error(t, err)
}

func TestGomail_ReconnectsIfSenderFails(t *testing.T) {
	failed := false
	sender := getSendCloser(func(from string, to []string, msg io.WriterTo) error {
		if !failed {
			failed = true
			return errors.New("Failed sending fake mail")
		}
		return nil
	})
	dialer := getDialer(sender)
	tpl, err := template.New("test").Parse("Hello world")
	require.NoError(t, err)

	notifier := NewGomailTemplate(logger, dialer, "Test email", tpl)
	err = notifier.Send(map[string]string{
		EmailFromKey: "from@test.com",
		EmailToKey:   "to@test.com",
	})
	require.NoError(t, err)

	require.Len(t, dialer.DialCalls(), 2) //First one failed at sending
}
