package notifier

import (
	"bytes"
	"html/template"
	"sync"

	"github.com/go-mail/mail"
	"github.com/gobuffalo/packr"
	"github.com/sirupsen/logrus"
)

//go:generate moq -out dialer_mock.go . Dialer
type Dialer interface {
	Dial() (mail.SendCloser, error)
}

type gomailNotifier struct {
	dialer       Dialer
	sender       mail.Sender
	subject      string
	htmlTemplate *template.Template

	logger *logrus.Entry
	mtx    sync.Mutex
}

func NewSMTPDialer(host string, port int, username, password string) Dialer {
	return mail.NewDialer(host, port, username, password)
}

func NewGomailNewTransfer(logger *logrus.Entry, dialer Dialer) (Notifier, error) {
	templates := packr.NewBox("./templates")
	raw := templates.String("new_transfer.html")
	tpl, err := template.New("new_transfer").Parse(raw)
	if err != nil {
		return nil, err
	}

	subject := "You've received kUSD!"
	return NewGomailTemplate(logger, dialer, subject, tpl), nil
}

func NewGomailTemplate(logger *logrus.Entry, dialer Dialer, subject string, htmlTemplate *template.Template) Notifier {
	return &gomailNotifier{
		dialer:       dialer,
		subject:      subject,
		logger:       logger.WithField("app", "notifier/gomail"),
		htmlTemplate: htmlTemplate.Option("missingkey=error"),
	}
}

func (notifier *gomailNotifier) Send(vars map[string]string) error {
	notifier.logger.Debug("Sending...")
	notifier.mtx.Lock()
	defer notifier.mtx.Unlock()

	if notifier.sender == nil {
		notifier.logger.Debug("Not connected yet, connecting...")
		err := notifier.connect()
		if err != nil {
			notifier.logger.WithError(err).Error("Error connecting")
			return err
		}
	}
	m := mail.NewMessage()
	m.SetHeader("From", vars[EmailFromKey])
	m.SetHeader("To", vars[EmailToKey])
	m.SetHeader("Subject", notifier.subject)

	var body bytes.Buffer
	err := notifier.htmlTemplate.Execute(&body, vars)
	if err != nil {
		notifier.logger.WithError(err).Error("Error executing the template")
		return err
	}

	m.SetBody("text/html", body.String())

	err = mail.Send(notifier.sender, m)
	if err == nil {
		// Sent
		return nil
	}

	notifier.logger.WithError(err).Debug("Received an error while sending. Reconnecting to retry.")
	// retry, maybe the connection timed up
	err = notifier.connect()
	if err != nil {
		notifier.logger.WithError(err).Error("Error connecting")
		return err
	}
	err = mail.Send(notifier.sender, m)
	if err != nil {
		notifier.logger.WithError(err).Error("Error sending e-mail")
		return err
	}

	// Send after reconnect
	return nil
}

func (notifier *gomailNotifier) connect() error {
	sender, err := notifier.dialer.Dial()
	if err != nil {
		return err
	}
	notifier.sender = sender
	return nil
}
