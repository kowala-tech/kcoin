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
	total int // number of data chunks that compose the content
	count int // number of current data chunks

	dataMu     sync.Mutex
	data       []*Chunk         // stores data chunks
	membership *common.BitArray // indicates whether a data unit is present or not
	root       common.Hash      // trie root hash
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
		total:      total,
		root:       common.Hash{},
		data:       chunks,
		membership: membership,
		count:      total,
	}
}

/*
func (ds *DataSet) Add(chunk *Chunk) (bool, error) {
	ds.dataMu.Lock()
	defer ds.dataMu.Unlock()

	index := chunk.Index

	// invalid index
	if index >= ds.total || index < 0 {
		return false, ErrInvalidIndex
	}

	// if part already exists, return false
	if ds.data[index] != nil {
		return false, nil
	}

	// check hash proof

	// add data chunk
	ds.data[index] = chunk
	ds.membership.Set(index)
	ds.count++

	return true, nil
}

func (ds *DataSet) GetChunk(index int) *Chunk {
	ds.dataMu.Lock()
	defer ds.dataMu.Unlock()
	return ds.data[index]
}

func (ds *DataSet) IsComplete() bool {
	return ds.total == ds.count
}

func (set *DataSet) Get(i int) *DataUnit {
	return set.units[i]
}

func (set *DataSet) Add(unit *DataUnit) (bool, error) {
	// index validation
	if unit.Index >= set.meta.nUnits {
		return false, ErrDataSetInvalidIndex
	}

	// has membership
	if set.units[unit.Index] != nil {
		return false, nil
	}

	// verify proof

	// add the unit to the set
	set.units[unit.Index] = unit
	set.membership.Set(unit.Index)
	set.nMembers++
}
*/
