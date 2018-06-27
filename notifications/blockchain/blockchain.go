package blockchain

import (
	"math/big"

	"github.com/kowala-tech/kcoin/notifications/protocolbuffer"
)

type Block struct {
	Number       *big.Int
	Transactions []*protocolbuffer.Transaction
}

//go:generate moq -out block_handler_mock.go . BlockHandler
type BlockHandler interface {
	HandleBlock(*Block)
}

//go:generate moq -out blockchain_mock.go . Blockchain
type Blockchain interface {
	Start() error
	Stop()
	Seek(*big.Int) error
	OnBlock(BlockHandler) error
}
