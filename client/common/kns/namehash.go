package kns

import (
	"strings"

	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/crypto"
)

func NameHash(name string) common.Hash {
	if len(name) == 0 {
		return common.Hash{}
	}

	var node common.Hash

	labels := strings.Split(name, ".")
	for i := len(labels) - 1; i >= 0; i-- {
		labelSha := crypto.Keccak256Hash([]byte(labels[i]))
		node = crypto.Keccak256Hash(node.Bytes(), labelSha.Bytes())
	}

	return node
}
