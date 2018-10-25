package main

//go:generate go-bindata -nometadata -o website.go index.html
//go:generate gofmt -w -s website.go

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"html/template"
	"math/big"
	"net/http"
	"os"
	"reflect"
	"sync"
	"time"

	"github.com/kowala-tech/kcoin/client/common/hexutil"
	"github.com/kowala-tech/kcoin/client/core/types"
	"github.com/kowala-tech/kcoin/client/kcoinclient"
	"github.com/kowala-tech/kcoin/client/log"
	"github.com/kowala-tech/kcoin/client/rpc"
	"golang.org/x/net/websocket"
)

var (
	IPCFlag     = flag.String("ipc", "", "Path to the IPC file")
	apiPortFlag = flag.Int("apiport", 8080, "Listener port for the HTTP API connection")
	logFlag     = flag.Int("verbosity", 3, "Log level to use for the control panel service")
)

func main() {
	// Parse the flags and set up the logger to print everything requested
	flag.Parse()
	log.Root().SetHandler(log.LvlFilterHandler(log.Lvl(*logFlag), log.StreamHandler(os.Stderr, log.TerminalFormat(true))))

	// Load up and render the website
	tmpl := MustAsset("index.html")
	website := new(bytes.Buffer)
	err := template.Must(template.New("").Parse(string(tmpl))).Execute(website, map[string]interface{}{})
	if err != nil {
		log.Crit("Failed to render the template", "err", err)
	}

	// Assemble and start the control service
	control, err := newControl(*IPCFlag, website.Bytes())
	if err != nil {
		log.Crit("Failed to start control", "err", err)
	}
	defer control.close()

	if err := control.listenAndServe(*apiPortFlag); err != nil {
		log.Crit("Failed to launch the API", "err", err)
	}
}

type mintListItem struct {
	Amount    *big.Int `json:"amount"`
	Confirmed bool     `json:"confirmed"`
	Id        int      `json:"id"`
	To        string   `json:"to"`
}
type state struct {
	block    int64
	coinbase string
	accounts []string
	mintList []*mintListItem
}

// control represents a control backed
type control struct {
	state     *state
	stateLock sync.RWMutex
	index     []byte // Index page to serve up on the web
	rpcClient *rpc.Client
	client    *kcoinclient.Client
	conns     []*websocket.Conn
	connLock  sync.RWMutex
	headers   chan *types.Header
}

func newControl(ipcFile string, index []byte) (*control, error) {
	rpcClient, client, err := dial(ipcFile)
	if err != nil {
		return nil, err
	}

	return &control{
		rpcClient: rpcClient,
		client:    client,
		index:     index,
		state:     &state{},
	}, nil
}

// dial dials to the kcoin process. It retries every second up to ten times.
func dial(ipcFile string) (*rpc.Client, *kcoinclient.Client, error) {
	attempts := 10
	var err error
	var rpcClient *rpc.Client
	var client *kcoinclient.Client
	for attempts > 0 {
		rpcClient, err = rpc.Dial(ipcFile)
		if err == nil {
			client = kcoinclient.NewClient(rpcClient)
			return rpcClient, client, nil
		}
		attempts -= 1
		time.Sleep(1 * time.Second)
	}
	return nil, nil, err
}

// close terminates the Kowala connection and tears down the faucet.
func (ctrl *control) close() error {
	close(ctrl.headers)
	return nil
}

// listenAndServe registers the HTTP handlers for the faucet and boots it up
// for service user funding requests.
func (ctrl *control) listenAndServe(port int) error {
	go ctrl.syncBlock()
	go ctrl.syncAccounts()
	go ctrl.syncMintList()
	http.HandleFunc("/", ctrl.webHandler)
	http.Handle("/api", websocket.Handler(ctrl.apiHandler))

	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

// webHandler handles all non-api requests, simply flattening and returning the
// faucet website.
func (ctrl *control) webHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(ctrl.index)
}

func (ctrl *control) syncBlock() {
	ctrl.headers = make(chan *types.Header)
	sub, err := ctrl.client.SubscribeNewHead(context.Background(), ctrl.headers)
	if err != nil {
		log.Warn("error subscribing to new heads", "err", err)
		return
	}
	defer func() {
		sub.Unsubscribe()
	}()

	for header := range ctrl.headers {
		ctrl.stateLock.Lock()
		ctrl.state.block = header.Number.Int64()
		ctrl.stateLock.Unlock()
		ctrl.sendState()
	}
}

func (ctrl *control) syncAccounts() {
	lastCoinbase := ""
	lastAccounts := []string{}

	for {
		time.Sleep(time.Second)
		var coinbase string
		var accounts []string

		if err := ctrl.rpcClient.Call(&coinbase, "eth_coinbase"); err != nil {
			log.Warn("error fetching coinbase", "err", err)
			break
		}
		if err := ctrl.rpcClient.Call(&accounts, "eth_accounts"); err != nil {
			log.Warn("error fetching accounts", "err", err)
			break
		}

		sendState := false
		if coinbase != "" && lastCoinbase != coinbase {
			lastCoinbase = coinbase
			ctrl.stateLock.Lock()
			ctrl.state.coinbase = coinbase
			ctrl.stateLock.Unlock()
			sendState = true
		}

		if accounts != nil && !reflect.DeepEqual(accounts, lastAccounts) {
			lastAccounts = accounts
			ctrl.stateLock.Lock()
			ctrl.state.accounts = accounts
			ctrl.stateLock.Unlock()
			sendState = true
		}

		if sendState {
			ctrl.sendState()
		}
	}
}

func (ctrl *control) syncMintList() {
	for {
		time.Sleep(time.Second)

		var mintList []*mintListItem

		err := ctrl.rpcClient.Call(&mintList, "mtoken_mintList")
		if err != nil {
			log.Error(err.Error())
			break
		}
		ctrl.stateLock.Lock()
		ctrl.state.mintList = mintList
		ctrl.stateLock.Unlock()
		ctrl.sendState()
	}
}

// apiHandler handles requests for kcoin grants and transaction statuses.
func (ctrl *control) apiHandler(conn *websocket.Conn) {
	defer conn.Close()
	ctrl.register(conn)
	defer ctrl.unregister(conn)

	ctrl.sendStateToConn(conn)

	for {
		data := make(map[string]interface{})
		if err := websocket.JSON.Receive(conn, &data); err != nil {
			return
		}
		switch data["action"] {
		case "mint":
			governor := data["governor"].(string)
			mintAddress := data["mint_address"].(string)
			mintAmount, _ := big.NewInt(0).SetString(data["mint_amount"].(string), 10)
			mintUnit, _ := big.NewInt(0).SetString(data["mint_unit"].(string), 10)

			mintScale := big.NewInt(10)
			mintScale = mintScale.Exp(mintScale, mintUnit, nil)
			finalAmmount := mintAmount.Mul(mintAmount, mintScale)

			err := ctrl.rpcClient.Call(nil, "mtoken_mint", governor,
				mintAddress,
				hexutil.EncodeBig(finalAmmount),
			)
			if err != nil {
				log.Error(err.Error())
			}
		case "confirm_mint":
			governor := data["governor"].(string)
			id, _ := data["id"].(float64)

			bigId := big.NewInt(int64(id))
			err := ctrl.rpcClient.Call(nil, "mtoken_confirm", governor, hexutil.EncodeBig(bigId))
			if err != nil {
				log.Error(err.Error())
			}
		}
	}
}
func (ctrl *control) register(conn *websocket.Conn) {
	ctrl.connLock.Lock()
	ctrl.conns = append(ctrl.conns, conn)
	ctrl.connLock.Unlock()
}
func (ctrl *control) unregister(conn *websocket.Conn) {
	ctrl.connLock.Lock()
	for i, c := range ctrl.conns {
		if c == conn {
			ctrl.conns = append(ctrl.conns[:i], ctrl.conns[i+1:]...)
			break
		}
	}
	ctrl.connLock.Unlock()
}

func (ctrl *control) sendState() {
	ctrl.connLock.RLock()
	for _, conn := range ctrl.conns {
		ctrl.sendStateToConn(conn)
	}
	ctrl.connLock.RUnlock()
}

func (ctrl *control) sendStateToConn(conn *websocket.Conn) {
	ctrl.stateLock.RLock()
	if err := send(conn, map[string]interface{}{
		"coinbase": ctrl.state.coinbase,
		"block":    ctrl.state.block,
		"mintList": ctrl.state.mintList,
		"accounts": ctrl.state.accounts,
	}, 3*time.Second); err != nil {
		log.Warn("Failed to send state", "err", err)
	}
	ctrl.stateLock.RUnlock()
}

// sends transmits a data packet to the remote end of the websocket, but also
// setting a write deadline to prevent waiting forever on the node.
func send(conn *websocket.Conn, value interface{}, timeout time.Duration) error {
	if timeout == 0 {
		timeout = 60 * time.Second
	}
	conn.SetWriteDeadline(time.Now().Add(timeout))
	return websocket.JSON.Send(conn, value)
}
