package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/gorilla/mux"

	"github.com/magiconair/properties/assert"
)

func TestWeCanConnectToSelectedPort(t *testing.T) {
	conf := DefaultConfig()
	conf.listenPort = 8181

	s := createAndStartServer(t, conf, mux.NewRouter())
	defer s.Stop()

	res, err := http.Get(
		fmt.Sprintf("http://:%d", 8181),
	)
	if err != nil {
		t.Errorf("Error creating request: %s", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error reading body from response: %s", err)
	}

	assert.Equal(t, "404 page not found\n", string(body))
}

func TestWeCanFetchTheApiAndWeReceiveMockedData(t *testing.T) {
	s := createAndStartServer(t, DefaultConfig(), mux.NewRouter())
	defer s.Stop()
}

func createAndStartServer(t *testing.T, config *Config, router *mux.Router) Server {
	s, err := New(config, router)
	if err != nil {
		t.Errorf("Error creating server.")
	}

	go s.Start()

	time.Sleep(3 * time.Second)

	return s
}
