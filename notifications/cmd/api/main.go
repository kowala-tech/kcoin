package main

import (
	"fmt"
	"net"
	"os"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/go-redis/redis"
	"github.com/yourheropaul/inj"

	"github.com/kowala-tech/kcoin/notifications/core/api"
	"github.com/kowala-tech/kcoin/notifications/environment"
	"github.com/kowala-tech/kcoin/notifications/keyvalue"
	"github.com/kowala-tech/kcoin/notifications/persistence"
	"github.com/kowala-tech/kcoin/notifications/protocolbuffer"
)

func main() {
	envReader := environment.NewReaderOs()
	port := envReader.Read("PORT")
	redisAddr := envReader.Read("REDIS_ADDR")
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

	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err = redisClient.Ping().Result()
	if err != nil {
		panic(err)
	}

	log := logrus.NewEntry(logger)
	emailMappingServer := api.NewEmailMappingServer(log)
	transactionService := api.NewTransactionServiceServer(log)

	g := inj.NewGraph()
	g.Provide(
		emailMappingServer,
		transactionService,
		keyvalue.NewRedisNamespacedKeyValue(redisClient, "emails"),
		persistence.NewRedisPersistence(redisClient),
	)

	if valid, errors := g.Assert(); !valid {
		panic(errors)
	}

	addr := fmt.Sprintf(":%v", port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	protocolbuffer.RegisterEmailMappingServer(grpcServer, emailMappingServer)
	protocolbuffer.RegisterTransactionServiceServer(grpcServer, transactionService)
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
}
