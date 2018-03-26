package log

import (
	"github.com/dougEfresh/logzio-go"
	"fmt"
)

type logzioHandler struct {
	sender *logzio.LogzioSender
}

func NewLogzioHandler(token string) (*logzioHandler, error) {
	sender, err := logzio.New(token)
	if err != nil {
		return nil, err
	}

	return &logzioHandler{sender: sender}, nil
}

func (l *logzioHandler) Log(r *Record) error {
	msg := fmt.Sprintf("{ \"%s\": \"%s\"}", "message", r.Msg)
	return l.sender.Send([]byte(msg))
}
