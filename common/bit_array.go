package common

import (
	"sync"

	"github.com/kowala-tech/kUSD/log"
)

const (
	word = 64  		// 64 bit word
	div = 6 		// number of right shifts necessary to divide by 64
	mod = word - 1
)

// BitArray represents an array of bits
type BitArray struct {
	bitsMu sync.Mutex
	bits   []uint64
	nbits  int
}

// NewBitArray returns a new array of bits
func NewBitArray(nbits int) *BitArray {
	if nbits == 0 {
		log.Error("Failed to create bit array")
		return nil
	}

	return &BitArray{
		nbits: nbits,
		bits: make([]uint64, nbits>>div),
	}
}

func (array *BitArray) Size() int { return array.nbits }


// Get checks whether the bit is set to one or not
func (array *BitArray) Get(i int) bool {
	if (i >= array.nbits || i < 0) {
		return false
	}

	return array.bits[i>>div] & (uint64(1) << (i & mod))
}

// Set sets the given bit to 1
func (array *BitArray) Set(i uint) bool {
	if (i >= array.nbits || i < 0) {
		return false
	}

	array.bits[i>>div] |= uint64(1) << (i & mod)
	return true
}

// Clear resets the bit array
func (array *BitArray) Clear() {
	for i := range f.bits {
		array.bits[i] = 0
	}
}
