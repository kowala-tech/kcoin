package params

// These are network parameters that need to be constant between clients, but
// aren't necesarilly consensus related.
const (
	BloomBitsBlocks      uint64 = 4096 // number of blocks a single bloom bit section vector contains
	MaximumExtraDataSize uint64 = 32   // Maximum size extra data may be after Genesis.
	MaxCodeSize                 = 24576
	EpochDuration        uint64 = 30000
	HDCoinType                  = 91927009 // Hierarchical deterministic wallet coin_type (SLIP-44)
)
