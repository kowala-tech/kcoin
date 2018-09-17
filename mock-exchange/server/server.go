package server

import (
	"context"
	"fmt"
	"net/http"
)

type server struct {
	httpServ *http.Server
	config   *Config
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

type Server interface {
	Start() error
	Stop() error
}

func New(c *Config) (Server, error) {
	return &server{
		config: c,
	}, nil
}
