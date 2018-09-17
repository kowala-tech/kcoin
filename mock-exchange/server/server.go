package server

import "net/http"

type server struct {
	config *Config
}

func (s *server) Start() error {
	return http.ListenAndServe(":8080", nil)
}

func (s *server) Stop() error {
	panic("implement me")
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
