package params

import "math/big"

// These are network parameters that need to be constant between clients, but
// aren't necesarilly consensus related.

var (
	// ComputeUnitPrice represents the network's fixed compute unit price (400 Gwei/Shannon)
	ComputeUnitPrice = new(big.Int).Mul(new(big.Int).SetUint64(400), new(big.Int).SetUint64(Shannon))
)

const (
	// ComputeCapacity represents the network's compute capacity per block
	ComputeCapacity uint64 = 4712388

	// MinComputeCapacity is the minimum the compute capacity that may ever be.
	MinComputeCapacity uint64 = 5000

	// HDCoinType hierarchical deterministic wallet coin_type (SLIP-44)
	HDCoinType = 91927009

	// BloomBitsBlocks is the number of blocks a single bloom bit section vector
	// contains.
	BloomBitsBlocks uint64 = 4096

	// StackLimit is the maximum size of VM stack allowed.
	StackLimit uint64 = 1024

	// MaximumExtraDataSize maximum size extra data may be after Genesis.
	MaximumExtraDataSize uint64 = 32

	MaxCodeSize = 24576

	EpochDuration uint64 = 30000
)
