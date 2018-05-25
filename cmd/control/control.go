package main

//go:generate go-bindata -nometadata -o website.go index.html
//go:generate gofmt -w -s website.go

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/kowala-tech/kcoin/core/types"
	"github.com/kowala-tech/kcoin/kcoinclient"

	"github.com/kowala-tech/kcoin/log"
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

// control represents a control backed
type control struct {
	index   []byte // Index page to serve up on the web
	client  *kcoinclient.Client
	conns   []*websocket.Conn
	lock    sync.RWMutex
	headers chan *types.Header
}

func newControl(ipcFile string, index []byte) (*control, error) {
	client, err := dial(ipcFile)
	if err != nil {
		return nil, err
	}

	return &control{
		client: client,
		index:  index,
	}, nil
}

// dial dials to the kcoin process. It retries every second up to ten times.
func dial(ipcFile string) (*kcoinclient.Client, error) {
	attempts := 10
	var err error
	var client *kcoinclient.Client
	for attempts > 0 {
		client, err = kcoinclient.Dial(ipcFile)
		if err == nil {
			return client, nil
		}
		attempts -= 1
		time.Sleep(1 * time.Second)
	}
	return nil, err
}

// close terminates the Kowala connection and tears down the faucet.
func (ctrl *control) close() error {
	close(ctrl.headers)
	return nil
}

// listenAndServe registers the HTTP handlers for the faucet and boots it up
// for service user funding requests.
func (ctrl *control) listenAndServe(port int) error {
	go ctrl.sendBlocksToEveryone()
	http.HandleFunc("/", ctrl.webHandler)
	http.Handle("/api", websocket.Handler(ctrl.apiHandler))

	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

// webHandler handles all non-api requests, simply flattening and returning the
// faucet website.
func (ctrl *control) webHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(ctrl.index)
}

func (ctrl *control) sendBlocksToEveryone() {
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
		ctrl.lock.RLock()
		for _, conn := range ctrl.conns {
			if err := send(conn, map[string]interface{}{
				"block": header.Number.Int64(),
			}, 3*time.Second); err != nil {
				log.Warn("Failed to send initial stats to client", "err", err)
			}
		}
		ctrl.lock.RUnlock()
	}
}

// apiHandler handles requests for kcoin grants and transaction statuses.
func (ctrl *control) apiHandler(conn *websocket.Conn) {
	defer conn.Close()
	ctrl.register(conn)
	defer ctrl.unregister(conn)

	for {
		data := make(map[string]interface{})
		if err := websocket.JSON.Receive(conn, &data); err != nil {
			return
		}
	}
}
func (ctrl *control) register(conn *websocket.Conn) {
	ctrl.lock.Lock()
	ctrl.conns = append(ctrl.conns, conn)
	ctrl.lock.Unlock()
}
func (ctrl *control) unregister(conn *websocket.Conn) {
	ctrl.lock.Lock()
	for i, c := range ctrl.conns {
		if c == conn {
			ctrl.conns = append(ctrl.conns[:i], ctrl.conns[i+1:]...)
			break
		}
	}
	ctrl.lock.Unlock()
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
