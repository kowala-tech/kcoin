package cluster

import "encoding/hex"

const rootPath = ".."

func randStringBytes(n int) string {
	b := make([]byte, n/2)

	if _, err := src.Read(b); err != nil {
		panic(err)
	}

	return hex.EncodeToString(b)[:n]
}
