package root

import (
	"net/http"
	"os"

	"fmt"

	"context"

	"time"

	"os/signal"

	"syscall"

	"github.com/dougEfresh/kitz"
	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"github.com/kowala-tech/kcoin/kcoinclient"
	"github.com/kowala-tech/notifications/protocolbuffer"
	"github.com/kowala-tech/wallet-backend/application/api"
	"github.com/kowala-tech/wallet-backend/application/command"
	"github.com/kowala-tech/wallet-backend/application/websocket"
	"github.com/kowala-tech/wallet-backend/blockchain"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

const (
	nodeEndpointConfigKey                = "node-endpoint"
	nodePortConfigKey                    = "node-port"
	nodeLogzTokenConfigKey               = "logz-token"
	nodeDefaultNotificationsRPCConfigKey = "notifications-endpoint"

	nodeDefaultEndpoint                 = "http://rpcnode.testnet.kowala.io:30503"
	nodeDefaultPort                     = "8080"
	nodeDefaultLogzToken                = ""
	nodeDefaultEndpointNotificationsRPC = "api:3000"
)

var rootCmd *cobra.Command
var logger log.Logger
var endChan chan bool

func init() {
	rootCmd = &cobra.Command{
		Use:   "wallet-backend",
		Short: "Wallet backend allows to talk with the blockchain",
		Run: func(cmd *cobra.Command, args []string) {
			logger.Log(
				"type",
				"info",
				"msg",
				fmt.Sprintf("Starting server at port %s", viper.GetString(nodePortConfigKey)),
			)

			go http.ListenAndServe(
				fmt.Sprintf(":%s", viper.GetString(nodePortConfigKey)),
				createRouter(logger),
			)

			<-endChan

			logger.Log("type", "info", "msg", "finishing process...")
			os.Exit(0)
		},
	}

	rootCmd.Flags().StringP(nodeEndpointConfigKey, "n", nodeDefaultEndpoint, "The kcoin node that the server will connect.")
	viper.BindPFlag(nodeEndpointConfigKey, rootCmd.Flags().Lookup(nodeEndpointConfigKey))
	viper.BindEnv(nodeEndpointConfigKey, "NODE_ENDPOINT")

	rootCmd.Flags().StringP(nodePortConfigKey, "p", nodeDefaultPort, "The backend listen port.")
	viper.BindPFlag(nodePortConfigKey, rootCmd.Flags().Lookup(nodePortConfigKey))
	viper.BindEnv(nodePortConfigKey, "NODE_PORT")

	rootCmd.Flags().StringP(nodeLogzTokenConfigKey, "l", nodeDefaultLogzToken, "The access token for logz.io service.")
	viper.BindPFlag(nodeLogzTokenConfigKey, rootCmd.Flags().Lookup(nodeLogzTokenConfigKey))
	viper.BindEnv(nodeLogzTokenConfigKey, "LOGZ_TOKEN")

	rootCmd.Flags().StringP(nodeDefaultNotificationsRPCConfigKey, "m", nodeDefaultEndpointNotificationsRPC, "The endpoint of the notifications microservice.")
	viper.BindPFlag(nodeDefaultNotificationsRPCConfigKey, rootCmd.Flags().Lookup(nodeDefaultNotificationsRPCConfigKey))
	viper.BindEnv(nodeDefaultNotificationsRPCConfigKey, "TX_MS_ENDPOINT")

	endChan = make(chan bool)
}

//Execute is the main entry to execute the wallet backend.
func Execute() {
	createLogger()
	go runSignalListener()

	if err := rootCmd.Execute(); err != nil {
		logger.Log(
			"type",
			"fatal",
			"msg",
			err,
		)
		os.Exit(1)
	}
}

func createLogger() {
	var err error

	if token := viper.GetString(nodeLogzTokenConfigKey); token != "" {
		logger, err = kitz.New(token)
		if err != nil {
			logger.Log(
				"type",
				"fatal",
				"msg",
				err,
			)
			os.Exit(1)
		}
	} else {
		logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	}
}

func runSignalListener() {
	logger.Log(
		"type",
		"info",
		"msg",
		"listening for signals.",
	)
	signalChan := make(chan os.Signal)
	signal.Notify(
		signalChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	for {
		s := <-signalChan
		logger.Log(
			"type",
			"alert",
			"msg",
			fmt.Sprintf("Signal received: %s", s.String()),
		)

		endChan <- true
	}
}

func createRouter(l log.Logger) http.Handler {
	r := mux.NewRouter()

	nodeConnection := createNodeConnection(viper.GetString(nodeEndpointConfigKey))
	notificationsConn := createTransactionServiceClient(viper.GetString(nodeDefaultNotificationsRPCConfigKey))

	// Websocket
	r.Methods("GET").PathPrefix("/ws").Handler(createWebsocketHandler(l, nodeConnection))

	r.Methods("GET").Path("/api/blockheight").Handler(createAPIBlockHeightHandler(l, nodeConnection))
	r.Methods("GET").Path("/api/balance/{account}").Handler(createAPIBalanceHandler(l, nodeConnection))

	getTransactionsHandler := createAPIGetTransactionsHandler(l, notificationsConn)
	r.Methods("GET").Path("/api/transactions/{account}").Handler(getTransactionsHandler)
	r.Methods("GET").Path("/api/transactions/{account}/from/{fromblock}/to/{toblock}").Handler(getTransactionsHandler)
	r.Methods("GET").Path("/api/broadcasttx/{rawtx}").Handler(createBroadcastTransactionHandler(l, nodeConnection))

	return r
}

func createAPIBlockHeightHandler(l log.Logger, nodeConnection blockchain.Client) http.Handler {
	return api.NewBlockHeightHandler(
		l,
		command.GetBlockHeightHandler{
			Client: nodeConnection,
		},
	)
}

func createAPIBalanceHandler(l log.Logger, client blockchain.Client) http.Handler {
	return api.NewBalanceHandler(
		l,
		command.GetBalanceHandler{
			Client: client,
		},
	)
}

func createAPIGetTransactionsHandler(l log.Logger, client protocolbuffer.TransactionServiceClient) http.Handler {
	return api.NewGetTransactionsHandler(
		l,
		command.GetTransactionsHandler{
			Client: client,
		},
	)
}

func createBroadcastTransactionHandler(l log.Logger, client blockchain.Client) http.Handler {
	return api.NewBroadcastTransactionHandler(
		l,
		command.BroadcastTransactionHandler{
			Client: client,
		},
	)
}

func createNodeConnection(endpoint string) blockchain.Client {
	client, err := kcoinclient.Dial(endpoint)
	if err != nil {
		panic("Error connection kcoin rpc: " + err.Error())
	}

	// handshake trick to avoid first connection delay
	go heartBeatNodeConnection(client)

	return client
}

func createTransactionServiceClient(endpoint string) protocolbuffer.TransactionServiceClient {
	conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	if err != nil {
		panic("Error connection to notifications endpoint: " + err.Error())
	}

	return protocolbuffer.NewTransactionServiceClient(conn)
}

func heartBeatNodeConnection(client blockchain.Client) {
	for {
		client.BlockNumber(context.Background())
		time.Sleep(10 * time.Second)
	}
}

func createWebsocketHandler(l log.Logger, client blockchain.Client) *websocket.Handler {
	wsHandler := &websocket.Handler{
		Logger: l,
		GetBlockCmd: command.GetBlockHeightHandler{
			Client: client,
		},
	}

	return wsHandler
}
