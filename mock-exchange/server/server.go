package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Server interface {
	Start() error
	Stop() error
}

func New(c *Config, r *mux.Router) (Server, error) {
	return &server{
		config: c,
		router: r,
	}, nil
}

type server struct {
	httpServ *http.Server
	config   *Config
	router   *mux.Router
}

func (s *server) Start() error {
	s.httpServ = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.config.listenPort),
		Handler: nil,
	}

	return s.httpServ.ListenAndServe()
}

func (s *server) Stop() error {
	return s.httpServ.Shutdown(context.Background())
}
