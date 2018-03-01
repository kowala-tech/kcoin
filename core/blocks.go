package core

import "github.com/kowala-tech/kUSD/common"

// BadHashes represent a set of manually tracked bad hashes (usually hard forks)
var BadHashes = map[common.Hash]bool{}
