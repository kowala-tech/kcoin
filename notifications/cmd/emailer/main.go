package main

import (
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/kowala-tech/kcoin/notifications/notifier"

	"github.com/sirupsen/logrus"

	"github.com/go-redis/redis"
	"github.com/yourheropaul/inj"

	"github.com/kowala-tech/kcoin/notifications/core"
	"github.com/kowala-tech/kcoin/notifications/environment"
	"github.com/kowala-tech/kcoin/notifications/keyvalue"
	"github.com/kowala-tech/kcoin/notifications/pubsub"
)

func main() {
	exitSignal := make(chan os.Signal)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)

	envReader := environment.NewReaderOs()
	redisAddr := envReader.Read("REDIS_ADDR")
	nsqAddr := envReader.Read("NSQ_ADDR")
	logLevelRaw := envReader.Read("LOG_LEVEL")
	smtpFrom := envReader.Read("SMTP_FROM")
	smtpHost := envReader.Read("SMTP_HOST")
	smtpPortRaw := envReader.Read("SMTP_PORT")
	smtpUsername := envReader.Read("SMTP_USERNAME")
	smtpPassword := envReader.Read("SMTP_PASSWORD")
	if logLevelRaw == "" {
		logLevelRaw = "info"
	}

	smtpPort, err := strconv.Atoi(smtpPortRaw)
	if err != nil {
		panic(err)
	}

	logLevel, err := logrus.ParseLevel(logLevelRaw)
	if err != nil {
		panic(err)
	}

	logger := logrus.New()
	logger.SetLevel(logLevel)
	logger.Out = os.Stdout

	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err = redisClient.Ping().Result()
	if err != nil {
		panic(err)
	}

	sub := pubsub.NewNSQSubscriber("transactions", "emailer", nsqAddr, logrus.NewEntry(logger))

	dialer := notifier.NewSMTPDialer(smtpHost, smtpPort, smtpUsername, smtpPassword)
	notif, err := notifier.NewGomailNewTransfer(logrus.NewEntry(logger), dialer)
	if err != nil {
		panic(err)
	}

	worker := core.NewEmailer(logrus.NewEntry(logger), smtpFrom)

	g := inj.NewGraph()
	g.Provide(
		worker,
		keyvalue.NewRedisNamespacedKeyValue(redisClient, "emails"),
		sub,
		notif,
	)

	if valid, errors := g.Assert(); !valid {
		panic(strings.Join(errors, ", "))
	}

	worker.Register()
	err = sub.Start()
	if err != nil {
		panic(err)
	}

	<-exitSignal
	sub.Stop()
	redisClient.Close()
}
