package server

type server struct {
}

func (s *server) Start() error {
	panic("implement me")
}

func (s *server) Stop() error {
	panic("implement me")
}

type Server interface {
	Start() error
	Stop() error
}

func New() (Server, error) {
	return &server{}, nil
}
