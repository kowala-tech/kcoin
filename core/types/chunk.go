package types

import (
	"errors"
	"sync"
	"sync/atomic"

	"github.com/kowala-tech/kUSD/common"
	"github.com/kowala-tech/kUSD/common/hexutil"
)

// @TODO (rgeraldes) - review uint64/int

//go:generate gencodec -type Chunk -field-override chunkMarshalling -out gen_chunk_json.go

var (
	ErrInvalidIndex = errors.New("invalid index")
)

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

type chunkMarshaling struct {
	Index hexutil.Uint64
	Data  hexutil.Bytes
}

// DataSet represents content as a set of data chunks
type DataSet struct {
	header Metadata

	count      int // number of current data chunks
	dataMu     sync.Mutex
	data       []*Chunk         // stores data chunks
	membership *common.BitArray // indicates whether a data unit is present or not

}

func NewDataSetFromData(data []byte, size int) DataSet {
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

	return DataSet{
		header: Metadata{
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

func (ds DataSet) Metadata() Metadata {
	return ds.header
}
