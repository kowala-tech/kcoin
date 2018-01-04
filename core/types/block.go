package types

import (
	"fmt"
	"io"
	"math/big"
	"sort"
	"sync/atomic"
	"time"

	"github.com/kowala-tech/kUSD/common"
	"github.com/kowala-tech/kUSD/common/hexutil"
	"github.com/kowala-tech/kUSD/crypto/sha3"
	"github.com/kowala-tech/kUSD/rlp"
)

var (
	EmptyRootHash = DeriveSha(Transactions{})
)

//go:generate gencodec -type Header -field-override headerMarshaling -out gen_header_json.go

// Header represents a block header in the Kowala blockchain.
type Header struct {
	Number         *big.Int    `json:"number"           gencodec:"required"`
	Time           *big.Int    `json:"timestamp"        gencodec:"required"`
	GasLimit       *big.Int    `json:"gasLimit"         gencodec:"required"`
	GasUsed        *big.Int    `json:"gasUsed"          gencodec:"required"`
	ParentHash     common.Hash `json:"parentHash"       gencodec:"required"`
	Root           common.Hash `json:"stateRoot"        gencodec:"required"`
	TxHash         common.Hash `json:"transactionsRoot" gencodec:"required"`
	ReceiptHash    common.Hash `json:"receiptsRoot"     gencodec:"required"`
	LastCommitHash common.Hash `json:"lastCommit"		gencodec:"required"`
	ValidatorsHash common.Hash `json:"validators"   	gencodec:"required"`
	Bloom          Bloom       `json:"logsBloom"        gencodec:"required"`
}

// headerMarshaling - field type overrides for gencodec.
type headerMarshaling struct {
	Number   *hexutil.Big
	GasLimit *hexutil.Big
	GasUsed  *hexutil.Big
	Time     *hexutil.Big
	Hash     common.Hash `json:"hash"` // adds call to Hash() in MarshalJSON
}

// Hash returns the block hash of the header, which is simply the keccak256 hash of its
// RLP encoding.
func (h *Header) Hash() common.Hash {
	return rlpHash(h)
}

func rlpHash(x interface{}) (h common.Hash) {
	hw := sha3.NewKeccak256()
	rlp.Encode(hw, x)
	hw.Sum(h[:0])
	return h
}

//go:generate gencodec -type Commit -out gen_commit_json.go

// Commit contains the evidence that a block was committed by a set of validators
type Commit struct {
	// @NOTE (rgeraldes) - pre-commits are in order of address
	preCommits Votes `json:"votes"	gencodec:"required"`
}

func (c *Commit) PreCommits() Votes {
	return c.preCommits
}

func (c *Commit) Hash() common.Hash {
	return rlpHash(c)
}

// Body is a simple (mutable, non-safe) data container for storing and moving
// a block's data contents (transactions) together.
type Body struct {
	Transactions []*Transaction
}

// Block represents an entire block in the KUSD blockchain.
type Block struct {
	header       *Header
	transactions Transactions
	lastCommit   *Commit

	// caches
	hash atomic.Value
	size atomic.Value

	// These fields are used by package kusd to track
	// inter-peer block relay.
	ReceivedAt   time.Time
	ReceivedFrom interface{}
}

// "external" block encoding. used for kusd protocol, etc.
type extblock struct {
	Header *Header
	Txs    []*Transaction
}

// NewBlock creates a new block. The input data is copied,
// changes to header and to the field values will not affect the
// block.
//
// The values of TxHash, ReceiptHash and Bloom in header
// are ignored and set to values derived from the given txs and receipts.
func NewBlock(header *Header, txs []*Transaction, receipts []*Receipt, commit *Commit) *Block {
	b := &Block{header: CopyHeader(header)}

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
		b.header.LastCommitHash = commit.Hash()
		copy(b.lastCommit, commit)
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
	if cpy.GasLimit = new(big.Int); h.GasLimit != nil {
		cpy.GasLimit.Set(h.GasLimit)
	}
	if cpy.GasUsed = new(big.Int); h.GasUsed != nil {
		cpy.GasUsed.Set(h.GasUsed)
	}
	return &cpy
}

// DecodeRLP decodes the KUSD RLP block format
func (b *Block) DecodeRLP(s *rlp.Stream) error {
	var eb extblock
	_, size, _ := s.Kind()
	if err := s.Decode(&eb); err != nil {
		return err
	}
	b.header, b.transactions = eb.Header, eb.Txs
	b.size.Store(common.StorageSize(rlp.ListSize(size)))
	return nil
}

// EncodeRLP serializes b into the KUSD RLP block format.
func (b *Block) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, extblock{
		Header: b.header,
		Txs:    b.transactions,
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

func (b *Block) Number() *big.Int   { return new(big.Int).Set(b.header.Number) }
func (b *Block) GasLimit() *big.Int { return new(big.Int).Set(b.header.GasLimit) }
func (b *Block) GasUsed() *big.Int  { return new(big.Int).Set(b.header.GasUsed) }
func (b *Block) Time() *big.Int     { return new(big.Int).Set(b.header.Time) }

func (b *Block) NumberU64() uint64          { return b.header.Number.Uint64() }
func (b *Block) Bloom() Bloom               { return b.header.Bloom }
func (b *Block) Coinbase() common.Address   { return b.header.Coinbase }
func (b *Block) Root() common.Hash          { return b.header.Root }
func (b *Block) ParentHash() common.Hash    { return b.header.ParentHash }
func (b *Block) TxHash() common.Hash        { return b.header.TxHash }
func (b *Block) ReceiptHash() common.Hash   { return b.header.ReceiptHash }
func (b *Block) ValidatorHash() common.Hash { return b.header.ValidatorHash }

func (b *Block) Header() *Header { return CopyHeader(b.header) }

// Body returns the non-header content of the block.
func (b *Block) Body() *Body { return &Body{b.transactions} }

func (b *Block) Commit() *Commit { return b.lastCommit }

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
		transactions: b.transactions,
	}
}

// WithBody returns a new block with the given transactions.
func (b *Block) WithBody(transactions []*Transaction) *Block {
	block := &Block{
		header:       CopyHeader(b.header),
		transactions: make([]*Transaction, len(transactions)),
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

func (b *Block) AsDataChunks(chunkSize int) (BlockChunks, error) {
	raw, err := rlp.EncodeToBytes(b)
	if err != nil {
		return nil, err
	}
	return NewDataSetFromData(raw, chunkSize), nil
}

func (b *Block) String() string {
	str := fmt.Sprintf(`Block(#%v): Size: %v {
%v
Transactions:
%v
}
`, b.Number(), b.Size(), b.header, b.transactions)
	return str
}

func (h *Header) String() string {
	return fmt.Sprintf(`Header(%x):
[
	ParentHash:	    %x
	Root:		    %x
	TxSha		    %x
	ReceiptSha:	    %x
	ValidatorSha:	%x
	LastCommitHash: %x
	Bloom:		    %x
	Number:		    %v
	GasLimit:	    %v
	GasUsed:	    %v
	Time:		    %v
]`, h.Hash(), h.ParentHash, h.Root, h.TxHash, h.ReceiptHash, h.ValidatorHash, h.LastCommitHash, h.Bloom, h.Number, h.GasLimit, h.GasUsed, h.Time)
}

type BlockChunks *DataSet

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
