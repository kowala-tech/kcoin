package wal

type walError string

const (
	ErrNotFound walError = "there is no data got the block"
	ErrWALStorageFailed walError = "failed to ensure WAL directory is in place"
)

func (e walError) Error() string {
	return string(e)
}

func IsNotFound(err error) bool {
	return err == ErrNotFound
}