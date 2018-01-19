package types

import (
	"bytes"
	"errors"
	"sync"
	"sync/atomic"

	"github.com/kowala-tech/kUSD/common"
	"github.com/kowala-tech/kUSD/common/hexutil"
	"github.com/kowala-tech/kUSD/rlp"
)

// @TODO (rgeraldes) - review uint64/int

//go:generate gencodec -type Chunk -field-override chunkMarshalling -out gen_chunk_json.go

var (
	ErrInvalidIndex = errors.New("invalid index")
)

// @TODO (rgeraldes) - move to another place
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// Chunk represents a fragment of information
type Chunk struct {
	Index uint64      `json:"index"  gencodec:"required"`
	Data  []byte      `json:"bytes"  gencodec:"required"`
	Proof common.Hash `json:"proof"  gencodec:"required"`

	// caches
	hash atomic.Value
	size atomic.Value
}

type chunkMarshalling struct {
	Index hexutil.Uint64
	Data  hexutil.Bytes
}

// DataSet represents content as a set of data chunks
type DataSet struct {
	meta *Metadata

	count      int // number of current data chunks
	dataMu     sync.Mutex
	data       []*Chunk         // stores data chunks
	membership *common.BitArray // indicates whether a data unit is present or not

}

func NewDataSetFromMeta(meta *Metadata) *DataSet {	
	// @TODO (rgeraldes) - review (pointers)
	return &DataSet{ 
		meta: CopyMeta(meta),
		count: 0,
		membership: common.NewBitArray(uint64(meta.nchunks)),
		data: make([]*Chunk, meta.nchunks),
	}
}
	
func CopyMeta(meta *Metadata) *Metadata {
	cpy := *meta
	return &cpy
}

func NewDataSetFromData(data []byte, size int) *DataSet {
	total := (len(data) + size - 1) / size
	chunks := make([]*Chunk, total)
	membership := common.NewBitArray(uint64(total))
	for i := 0; i < total; i++ {
		chunk := &Chunk{
			Index: uint64(i),
			Data:  data[i*size : min(len(data), (i+1)*size)],
		}
		chunks[i] = chunk
		membership.Set(i)
	}

	// @TODO (rgeraldes)
	// compute merkle proofs
	//trie := new(trie.Trie)
	//trie.Update()
	//root := trie.Hash()

	return &DataSet{
		meta: &Metadata{
			nchunks: total,
			root:    common.Hash{},
		},
		data:       chunks,
		membership: membership,
		count:      total,
	}
}

// Metadata represents the content specifications
type Metadata struct {
	nchunks int         `json:"nchunks"  gencodec:"required"`
	root    common.Hash `json:"proof"  	 gencodec:"required"` // root hash of the trie
}

func (ds DataSet) Metadata() *Metadata {
	return ds.meta
}

func (ds DataSet) Size() int {
	return ds.meta.nchunks
}

func (ds DataSet) Get(i int) *Chunk {
	// @TODO (rgeraldes) - add logic to verify if the fragment
	// exists
	return ds.data[i]
}

func (ds DataSet) Add(chunk *Chunk) {
	// @TODO (rgeraldes) - validate index
	// @TODO (rgeraldes) - check hash proof
	ds.data[chunk.Index] = chunk
	// @TODO (rgeraldes) - review int vs uint64
	ds.membership.Set(int(chunk.Index))
	ds.count++
}

func (ds DataSet) HasAll() bool {
	return ds.count == ds.meta.nchunks
}

func (ds DataSet) Data() []byte {
	var buffer bytes.Buffer
	for _, chunk := range ds.data {
		buffer.Write(chunk.Data)
	}
	return buffer.Bytes()
}

func (ds DataSet) Assemble() (*Block, error) {
	var block Block
	if err := rlp.DecodeBytes(ds.Data(), &block); err != nil {
		return nil, err
	}
	return &block, nil
}