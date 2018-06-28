package main

import (
	"os"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/go-redis/redis"
	"github.com/yourheropaul/inj"

	"github.com/kowala-tech/kcoin/notifications/blockchain"
	"github.com/kowala-tech/kcoin/notifications/core"
	"github.com/kowala-tech/kcoin/notifications/environment"
	"github.com/kowala-tech/kcoin/notifications/keyvalue"
	"github.com/kowala-tech/kcoin/notifications/persistence"
)

func main() {
	envReader := environment.NewReaderOs()
	rpcURI := envReader.Read("TESTNET_RPC_ADDR")
	redisAddr := envReader.Read("REDIS_ADDR")
	pollingStr := envReader.Read("POLLING_INTERVAL")
	logLevelRaw := envReader.Read("LOG_LEVEL")
	if logLevelRaw == "" {
		logLevelRaw = "info"
	}

	logLevel, err := logrus.ParseLevel(logLevelRaw)
	if err != nil {
		panic(err)
	}

	logger := logrus.New()
	logger.SetLevel(logLevel)
	logger.Out = os.Stdout

	var pollingSeconds int
	if pollingStr == "" {
		pollingSeconds = 5
	} else {
		parsed, err := strconv.Atoi(pollingStr)
		if err != nil {
			panic(err)
		}
		pollingSeconds = parsed
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err = redisClient.Ping().Result()
	if err != nil {
		panic(err)
	}

	worker := core.NewDbTransactionsSynchronizer(logrus.NewEntry(logger))

	g := inj.NewGraph()
	g.Provide(
		worker,
		blockchain.NewKcoin(rpcURI, pollingSeconds, logrus.NewEntry(logger)),
		keyvalue.WrapKeyValue(keyvalue.NewRedisKeyValue(redisClient), "db_sync_latest_block"),
		persistence.NewRedisPersistence(redisClient),
	)

	if valid, errors := g.Assert(); !valid {
		panic(errors)
	}

	err = worker.Start()
	if err != nil {
		panic(err)
	}
}
