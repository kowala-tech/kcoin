package params

// These are network parameters that need to be constant between clients, but
// aren't necesarilly consensus related.
const (
	// BloomBitsBlocks is the number of blocks a single bloom bit section vector
	// contains.
	BloomBitsBlocks uint64 = 4096
	// MaximumExtraDataSize maximum size extra data may be after Genesis.
	MaximumExtraDataSize uint64 = 32
	MaxCodeSize                 = 24576
	EpochDuration        uint64 = 30000
	HDCoinType                  = 91927009 // Hierarchical deterministic wallet coin_type (SLIP-44)
)
