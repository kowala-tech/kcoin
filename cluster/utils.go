package cluster

import (
	"encoding/hex"
	"math/rand"
	"time"
)

const rootPath = ".."

var src = rand.New(rand.NewSource(time.Now().UnixNano()))

func randStringBytes(n int) string {
	b := make([]byte, n/2)

	if _, err := src.Read(b); err != nil {
		panic(err)
	}

	return hex.EncodeToString(b)[:n]
}
