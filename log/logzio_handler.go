package log

import (
	"github.com/dougEfresh/logzio-go"
)

type logzioHandler struct {
	sender *logzio.LogzioSender
	format Format
}

func NewLogzioHandler(token string) (*logzioHandler, error) {
	sender, err := logzio.New(token)
	if err != nil {
		return nil, err
	}
	return &logzioHandler{sender, JsonFormat()}, nil
}

func (l *logzioHandler) Log(r *Record) error {
	return l.sender.Send(l.format.Format(r))
}
