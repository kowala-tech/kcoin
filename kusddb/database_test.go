package kusddb

import (
	"os"
	"path/filepath"

	"github.com/kowala-tech/kUSD/common"
)

func newDb() *LDBDatabase {
	file := filepath.Join("/", "tmp", "ldbtesttmpfile")
	if common.FileExist(file) {
		os.RemoveAll(file)
	}
	db, _ := NewLDBDatabase(file, 0, 0)

	return db
}
