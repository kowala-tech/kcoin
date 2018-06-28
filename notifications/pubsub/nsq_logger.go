package pubsub

import (
	"strings"

	nsq "github.com/nsqio/go-nsq"
	"github.com/sirupsen/logrus"
)

var (
	nsqDebugLevel = nsq.LogLevelDebug.String()
	nsqInfoLevel  = nsq.LogLevelInfo.String()
	nsqWarnLevel  = nsq.LogLevelWarning.String()
	nsqErrLevel   = nsq.LogLevelError.String()
)

type nsqLogger struct {
	logger *logrus.Entry
}

func NewNSQLogger(logger *logrus.Entry) (*nsqLogger, nsq.LogLevel) {
	var level nsq.LogLevel
	switch logger.Level {
	case logrus.DebugLevel:
		level = nsq.LogLevelDebug
	case logrus.InfoLevel:
		level = nsq.LogLevelInfo
	case logrus.WarnLevel:
		level = nsq.LogLevelWarning
	case logrus.ErrorLevel:
		level = nsq.LogLevelError
	}
	return &nsqLogger{logger}, level
}

func (n *nsqLogger) Output(_ int, s string) error {
	if len(s) > 3 {
		msg := strings.TrimSpace(s[3:])
		switch s[:3] {
		case nsqDebugLevel:
			n.logger.Debugln(msg)
		case nsqInfoLevel:
			n.logger.Infoln(msg)
		case nsqWarnLevel:
			n.logger.Warnln(msg)
		case nsqErrLevel:
			n.logger.Errorln(msg)
		default:
			n.logger.Infoln(msg)
		}
	}
	return nil
}
