package params

// Proof-of-stake timeouts
const (
	ProposeDuration        uint64 = 500
	ProposeDeltaDuration   uint64 = 25
	PreVoteDuration        uint64 = 200
	PreVoteDeltaDuration   uint64 = 25
	PreCommitDuration      uint64 = 200
	PreCommitDeltaDuration uint64 = 25
	BlockTime              uint64 = 1000
)

const (
	// CallCreateDepth maximum depth of call/create stack
	CallCreateDepth uint64 = 1024

	// StackLimit is the maximum size of VM stack allowed.
	StackLimit uint64 = 1024
)
