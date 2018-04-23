package wal

import (
	"path/filepath"
	"github.com/kowala-tech/kcoin/common/io"
	"github.com/kowala-tech/kcoin/log"
	"os"
	"fmt"
)

type baseWal struct {
	storage 	*os.File
}

func New(file string) (*baseWal, error) {
	err := io.EnsureDir(filepath.Dir(file), 0700)
	if err != nil {
		fmt.Println(err)
		log.Warn(ErrWALStorageFailed.Error(), "err", err)
		return nil, ErrWALStorageFailed
	}

	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!! NEW WAL", file)
	f, err := io.OpenFile(file)
	if err != nil {
		log.Warn(ErrWALStorageFailed.Error(), "err", err)
		return nil, ErrWALStorageFailed
	}

	return &baseWal{storage: f}, nil
}

func (wal baseWal) Start() error {
	return nil
}

func (wal baseWal) Stop() error {
	return nil
}

func (wal baseWal) Save(message Message) {
	n, err := wal.storage.Write(message.Byte())
	fmt.Println("!!!!!!!!!!!!!!! SAVE", n, err)
}

func (wal baseWal) Messages(height int64, options ...Options) (messages <-chan Message, err error) {
	return
}
