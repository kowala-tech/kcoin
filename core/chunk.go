package core 

import (
	"sync"
	"bytes"

	"github.com/kowala-tech/kUSD/trie"
)

//go:generate gencodec -type unitdata -field-override unitdataMarshalling -out gen_dataunit_json.go

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
	index uint	       	`json:"index"  gencodec:"required"`
   	data []byte       	`json:"bytes"  gencodec:"required"`
	proof common.Hash  	`json:"proof"  gencodec:"required"`
	
	// caches
	hash atomic.Value
	size atomic.Value
}

// Hash returns the chunk hash, which is simply the keccak256 hash of its
// RLP encoding.
func (ck *Chunk) Hash() common.Hash {
	if hash := ck.hash.Load(); hash != nil {
		return hash.(common.Hash)
	}
	v := rlpHash(ck)
	ck.hash.Store(v)
	return v
}

// Size returns the chunk size.
func (ck *Chunk) Size() common.StorageSize {
	if size := b.size.Load(); size != nil {
		return size.(common.StorageSize)
	}
	c := writeCounter(0)
	rlp.Encode(&c, ck)
	ck.size.Store(common.StorageSize(c))
	return common.StorageSize(c)
}

func (ck *Chunk) Index() int			{ return ck.index }
func (ck *Chunk) Data() []byte			{ return ck.data  }
func (ck *Chunk) Proof() common.Hash 	{ return ck.proof }

// Metadata represents the content specifications
type Metadata struct {
	nchunks int			`json:"nchunks"  gencodec:"required"`
	root common.Hash	`json:"proof"  	 gencodec:"required"` // root hash of the trie
}

// DataSet represents content as a set of data chunks
type DataSet struct {
	total int			 	// number of data chunks that compose the content
	count int				// number of current data chunks

	dataMu sync.Mutex
	data []*Chunk	 		// stores data chunks
	membership BitArray  	// indicates whether a data unit is present or not
	root common.Hash		// trie root hash
}

func NewDataSetFromData(data []byte, chunkSize int) {
	total := (len(data) + chunkSize - 1) / chunkSize
	chunks := make([]*Chunk, total)
	membership := NewBitArray(total)
	for i := 0; i < total; i++ {
		chunk := &Chunk{
			index: i,
			// @TODO (rgeraldes) - min int
			data: data[i*chunkSize : min(len(data), (i+1)*dataUnitSize)]
			chunks[i] = chunk
			membership.Set(i)
		}
	}

	// merkle proof for each chunk
	trie := new(trie.Trie)
	trie.Update()
	root := trie.Hash()
}

func (ds *DataSet) Add(chunk *Chunk) (bool, error) {
	ds.dataMu.Lock()
	defer ds.dataMu.Unlock()

	index := chunk.index

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


func (set *DataSet) IsComplete() bool {
	return set.nMembers == set.meta.nUnits
}
