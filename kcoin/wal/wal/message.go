package wal

import (
	"math/big"
	"fmt"
	"github.com/kowala-tech/kcoin/common"
	"github.com/kowala-tech/kcoin/crypto/sha3"
	"github.com/kowala-tech/kcoin/rlp"
)

type Message interface {
	Byte() []byte
	Code() byte
	Block() *big.Int
	Valid() bool
}

type code byte

func (c code) Code() byte {
	return byte(c)
}

type blockNumber struct {
	*big.Int
}

func (b blockNumber) Block() *big.Int {
	return b.Int
}

type peerID string

func (id peerID) PeerID() string {
	return string(id)
}

const (
	BlockStartCode = code(iota)
	BlockCommitCode
)

func hash(x interface{}) (h common.Hash) {
	hw := sha3.NewKeccak256()
	rlp.Encode(hw, x)
	hw.Sum(h[:0])
	return h
}

type blockStart struct {
	code
	blockNumber
	hash common.Hash
}

func BlockStart(blockHeight *big.Int) blockStart {
	b := blockStart{code: BlockStartCode, blockNumber: blockNumber{blockHeight}}
	b.hash = hash(b)
	return b
}

func (block blockStart) Byte() []byte {
	fmt.Println("##########################")
	return []byte("a new block started")
}


func (block blockStart) Valid() bool {
	if common.EmptyHash(block.hash) {
		return false
	}

	b := block
	b.hash = common.Hash{}
	hash := hash(b)

	return block.hash == hash
}

type blockCommit struct {
	code
	blockNumber
	hash common.Hash
}

func BlockCommit(blockHeight *big.Int) blockCommit {
	b := blockCommit{code: BlockCommitCode, blockNumber: blockNumber{blockHeight}}
	b.hash = hash(b)
	return b
}

func (block blockCommit) Byte() []byte {
	fmt.Println("%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%")
	return []byte("a new block commited")
}

func (block blockCommit) Valid() bool {
	if common.EmptyHash(block.hash) {
		return false
	}

	b := block
	b.hash = common.Hash{}
	hash := hash(b)

	return block.hash == hash
}