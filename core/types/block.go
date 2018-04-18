// Package types contains data types related to Ethereum consensus.
package types

import (
	"fmt"
	"io"
	"math/big"
	"sort"
	"sync/atomic"
	"time"
	"unsafe"

	"github.com/kowala-tech/kcoin/common"
	"github.com/kowala-tech/kcoin/common/hexutil"
	"github.com/kowala-tech/kcoin/crypto/sha3"
	"github.com/kowala-tech/kcoin/rlp"
)

var (
	EmptyRootHash = DeriveSha(Transactions{})
)

//go:generate gencodec -type Header -field-override headerMarshalling -out gen_header_json.go
//go:generate gencodec -type Commit -out gen_commit_json.go

// Header represents a block header in the Ethereum blockchain.
type Header struct {
	ParentHash     common.Hash    `json:"parentHash"       gencodec:"required"`
	Coinbase       common.Address `json:"miner"            gencodec:"required"`
	Root           common.Hash    `json:"stateRoot"        gencodec:"required"`
	TxHash         common.Hash    `json:"transactionsRoot" gencodec:"required"`
	ReceiptHash    common.Hash    `json:"receiptsRoot"     gencodec:"required"`
	ValidatorsHash common.Hash    `json:"validators"       gencodec:"required"`
	LastCommitHash common.Hash    `json:"lastCommit"       gencodec:"required"`
	Bloom          Bloom          `json:"logsBloom"        gencodec:"required"`
	Number         *big.Int       `json:"number"           gencodec:"required"`
	GasLimit       uint64         `json:"gasLimit"         gencodec:"required"`
	GasUsed        uint64         `json:"gasUsed"          gencodec:"required"`
	Time           *big.Int       `json:"timestamp"        gencodec:"required"`
	Extra          []byte         `json:"extraData"        gencodec:"required"`
}

// field type overrides for gencodec
type headerMarshaling struct {
	Number   *hexutil.Big
	GasLimit hexutil.Uint64
	GasUsed  hexutil.Uint64
	Time     *hexutil.Big
	Extra    hexutil.Bytes
	Hash     common.Hash `json:"hash"` // adds call to Hash() in MarshalJSON
}

// Hash returns the block hash of the header, which is simply the keccak256 hash of its
// RLP encoding.
func (h *Header) Hash() common.Hash {
	return rlpHash(h)
}

// HashNoNonce returns the hash which is used as input for the proof-of-stake search.
func (h *Header) HashNoNonce() common.Hash {
	return rlpHash([]interface{}{
		h.ParentHash,
		h.Coinbase,
		h.Root,
		h.TxHash,
		h.ReceiptHash,
		h.ValidatorsHash,
		h.LastCommitHash,
		h.Bloom,
		h.Number,
		h.GasLimit,
		h.GasUsed,
		h.Time,
		h.Extra,
	})
}

// Size returns the approximate memory used by all internal contents. It is used
// to approximate and limit the memory consumption of various caches.
func (h *Header) Size() common.StorageSize {
	return common.StorageSize(unsafe.Sizeof(*h)) + common.StorageSize(len(h.Extra)+(h.Number.BitLen()+h.Time.BitLen())/8)
}

func rlpHash(x interface{}) (h common.Hash) {
	hw := sha3.NewKeccak256()
	rlp.Encode(hw, x)
	hw.Sum(h[:0])
	return h
}

// Commit contains the evidence that a block was committed by a set of validators
type Commit struct {
	// @NOTE (rgeraldes) - pre-commits are in order of address
	PreCommits     Votes `json:"votes"    gencodec:"required"`
	FirstPreCommit *Vote `json:"vote"     gencodec:"required"`
}

func (cmt *Commit) Commits() Votes {
	return cmt.PreCommits
}

func (cmt *Commit) First() *Vote {
	return cmt.FirstPreCommit
}

func (cmt *Commit) Hash() common.Hash {
	return rlpHash(cmt)
}

func (cmt *Commit) Round() uint64 {
	if len(cmt.PreCommits) == 0 {
		return 0
	}

	if cmt.First() == nil {
		return 0
	}

	return cmt.First().Round()
}

// Body is a simple (mutable, non-safe) data container for storing and moving
// a block's data contents (transactions) together.
type Body struct {
	LastCommit   *Commit
	Transactions []*Transaction
}

// Block represents an entire block in the Ethereum blockchain.
type Block struct {
	header       *Header
	lastCommit   *Commit
	transactions Transactions

	// caches
	hash atomic.Value
	size atomic.Value

	// These fields are used by package eth to track
	// inter-peer block relay.
	ReceivedAt   time.Time
	ReceivedFrom interface{}
}

// "external" block encoding. used for kcoin protocol, etc.
type extblock struct {
	Header     *Header
	LastCommit *Commit
	Txs        []*Transaction
}

// NewBlock creates a new block. The input data is copied,
// changes to header and to the field values will not affect the
// block.
//
// The values of TxHash, ReceiptHash and Bloom in header
// are ignored and set to values derived from the given txs
// and receipts.
func NewBlock(header *Header, txs []*Transaction, receipts []*Receipt, commit *Commit) *Block {
	b := &Block{header: CopyHeader(header), lastCommit: &Commit{PreCommits: Votes{}, FirstPreCommit: &Vote{}}}

	// TODO: panic if len(txs) != len(receipts)
	if len(txs) == 0 {
		b.header.TxHash = EmptyRootHash
	} else {
		b.header.TxHash = DeriveSha(Transactions(txs))
		b.transactions = make(Transactions, len(txs))
		copy(b.transactions, txs)
	}

	if len(receipts) == 0 {
		b.header.ReceiptHash = EmptyRootHash
	} else {
		b.header.ReceiptHash = DeriveSha(Receipts(receipts))
		b.header.Bloom = CreateBloom(receipts)
	}

	if commit != nil {
		lastCommit := CopyCommit(commit)
		b.header.LastCommitHash = lastCommit.Hash()
		b.lastCommit = lastCommit
	}

	return b
}

// NewBlockWithHeader creates a block with the given header data. The
// header data is copied, changes to header and to the field values
// will not affect the block.
func NewBlockWithHeader(header *Header) *Block {
	return &Block{header: CopyHeader(header)}
}

// CopyHeader creates a deep copy of a block header to prevent side effects from
// modifying a header variable.
func CopyHeader(h *Header) *Header {
	cpy := *h
	if cpy.Time = new(big.Int); h.Time != nil {
		cpy.Time.Set(h.Time)
	}
	if cpy.Number = new(big.Int); h.Number != nil {
		cpy.Number.Set(h.Number)
	}
	if len(h.Extra) > 0 {
		cpy.Extra = make([]byte, len(h.Extra))
		copy(cpy.Extra, h.Extra)
	}
	return &cpy
}

// CopyCommit creates a deep copy of a block commit to prevent side efects from
// modifying a header variable
func CopyCommit(commit *Commit) *Commit {
	cpy := *commit

	if len(commit.PreCommits) > 0 {
		cpy.PreCommits = make(Votes, len(commit.PreCommits))
		copy(cpy.PreCommits, commit.PreCommits)
	}

	if commit.FirstPreCommit != nil {
		cpy.FirstPreCommit = &(*commit.FirstPreCommit)
	}

	return &cpy
}

// DecodeRLP decodes the Kowala block
func (b *Block) DecodeRLP(s *rlp.Stream) error {
	var eb extblock
	_, size, _ := s.Kind()
	if err := s.Decode(&eb); err != nil {
		return err
	}
	b.header, b.lastCommit, b.transactions, b.lastCommit = eb.Header, eb.LastCommit, eb.Txs, eb.LastCommit
	b.size.Store(common.StorageSize(rlp.ListSize(size)))
	return nil
}

// EncodeRLP serializes b into the Ethereum RLP block format.
func (b *Block) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, extblock{
		Header:     b.header,
		LastCommit: b.lastCommit,
		Txs:        b.transactions,
	})
}

// TODO: copies

func (b *Block) Transactions() Transactions { return b.transactions }

func (b *Block) Transaction(hash common.Hash) *Transaction {
	for _, transaction := range b.transactions {
		if transaction.Hash() == hash {
			return transaction
		}
	}
	return nil
}

func (b *Block) LastCommit() *Commit { return b.lastCommit }

func (b *Block) Number() *big.Int { return new(big.Int).Set(b.header.Number) }
func (b *Block) GasLimit() uint64 { return b.header.GasLimit }
func (b *Block) GasUsed() uint64  { return b.header.GasUsed }
func (b *Block) Time() *big.Int   { return new(big.Int).Set(b.header.Time) }

func (b *Block) NumberU64() uint64           { return b.header.Number.Uint64() }
func (b *Block) Bloom() Bloom                { return b.header.Bloom }
func (b *Block) Coinbase() common.Address    { return b.header.Coinbase }
func (b *Block) Root() common.Hash           { return b.header.Root }
func (b *Block) ParentHash() common.Hash     { return b.header.ParentHash }
func (b *Block) TxHash() common.Hash         { return b.header.TxHash }
func (b *Block) ReceiptHash() common.Hash    { return b.header.ReceiptHash }
func (b *Block) LastCommitHash() common.Hash { return b.header.LastCommitHash }
func (b *Block) ValidatorsHash() common.Hash { return b.header.ValidatorsHash }
func (b *Block) Extra() []byte               { return common.CopyBytes(b.header.Extra) }

func (b *Block) Header() *Header { return CopyHeader(b.header) }

// Body returns the non-header content of the block.
func (b *Block) Body() *Body { return &Body{b.lastCommit, b.transactions} }

// @TODO (rgeraldes) - review
func (b *Block) HashNoNonce() common.Hash {
	return b.header.HashNoNonce()
}

// Size returns the true RLP encoded storage size of the block, either by encoding
// and returning it, or returning a previsouly cached value.
func (b *Block) Size() common.StorageSize {
	if size := b.size.Load(); size != nil {
		return size.(common.StorageSize)
	}
	c := writeCounter(0)
	rlp.Encode(&c, b)
	b.size.Store(common.StorageSize(c))
	return common.StorageSize(c)
}

type writeCounter common.StorageSize

func (c *writeCounter) Write(b []byte) (int, error) {
	*c += writeCounter(len(b))
	return len(b), nil
}

// WithSeal returns a new block with the data from b but the header replaced with
// the sealed one.
func (b *Block) WithSeal(header *Header) *Block {
	cpy := *header

	return &Block{
		header:       &cpy,
		lastCommit:   b.lastCommit,
		transactions: b.transactions,
	}
}

// WithBody returns a new block with the given transaction contents.
func (b *Block) WithBody(transactions []*Transaction, lastCommit *Commit) *Block {
	block := &Block{
		header:       CopyHeader(b.header),
		transactions: make([]*Transaction, len(transactions)),
		lastCommit:   &Commit{},
	}

	if lastCommit != nil {
		block.lastCommit = lastCommit
	}

	copy(block.transactions, transactions)
	return block
}

// Hash returns the keccak256 hash of b's header.
// The hash is computed on the first call and cached thereafter.
func (b *Block) Hash() common.Hash {
	if hash := b.hash.Load(); hash != nil {
		return hash.(common.Hash)
	}
	v := b.header.Hash()
	b.hash.Store(v)
	return v
}

func (b *Block) AsFragments(size int) (*BlockFragments, error) {
	rawBlock, err := rlp.EncodeToBytes(b)
	if err != nil {
		return &BlockFragments{}, err
	}
	return NewDataSetFromData(rawBlock, size), nil
}

func (b *Block) String() string {
	str := fmt.Sprintf(`Block(#%v): Size: %v {
ValidatorHash: %x
%v
Transactions:
%v
LastCommit:
%v
}
`, b.Number(), b.Size(), b.header.HashNoNonce(), b.header, b.transactions, b.lastCommit)
	return str
}

func (h *Header) String() string {
	return fmt.Sprintf(`Header(%x):
[
	ParentHash:	    %x
	Coinbase:	    %x
	Root:		    %x
	TxSha		    %x
	ReceiptSha:	    %x
	ValidatorsHash: %x
	LastCommitHash: %x
	Bloom:		    %x
	Number:		    %v
	GasLimit:	    %v
	GasUsed:	    %v
	Time:		    %v
	Extra:		    %s
]`, h.Hash(), h.ParentHash, h.Coinbase, h.Root, h.TxHash, h.ReceiptHash, h.ValidatorsHash, h.LastCommitHash, h.Bloom, h.Number, h.GasLimit, h.GasUsed, h.Time, h.Extra)
}

type Blocks []*Block

type BlockBy func(b1, b2 *Block) bool

func (self BlockBy) Sort(blocks Blocks) {
	bs := blockSorter{
		blocks: blocks,
		by:     self,
	}
	sort.Sort(bs)
}

type blockSorter struct {
	blocks Blocks
	by     func(b1, b2 *Block) bool
}

func (self blockSorter) Len() int { return len(self.blocks) }
func (self blockSorter) Swap(i, j int) {
	self.blocks[i], self.blocks[j] = self.blocks[j], self.blocks[i]
}
func (self blockSorter) Less(i, j int) bool { return self.by(self.blocks[i], self.blocks[j]) }

func Number(b1, b2 *Block) bool { return b1.header.Number.Cmp(b2.header.Number) < 0 }

type BlockFragment = Chunk
type BlockFragments = DataSet
