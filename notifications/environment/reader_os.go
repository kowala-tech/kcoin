package environment

import "os"

type readerOs struct{}

// NewReaderOs creates a new Reader that wraps the standard OS environment.
func NewReaderOs() Reader {
	return &readerOs{}
}

func (r *readerOs) Read(key string) string {
	return os.Getenv(key)
}
