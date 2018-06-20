package core

import (
	"math/big"

	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/core/types"
)

// TxPreEvent is posted when a transaction enters the transaction pool.
type TxPreEvent struct{ Tx *types.Transaction }

// NewVoteEvent is posted when a consensus validator votes.
type NewVoteEvent struct{ Vote *types.Vote }

// NewProposalEvent is posted when a consensus validator proposes a new block.
type NewProposalEvent struct{ Proposal *types.Proposal }

// NewBlockFragmentEvent is posted when a consensus validator broadcasts block fragments.
type NewBlockFragmentEvent struct {
	BlockNumber *big.Int
	Round       uint64
	Data        *types.BlockFragment
}

// NewMajorityEvent is posted when there's a majority during a sub election
type NewMajorityEvent struct {
	winner common.Hash
}

// PendingLogsEvent is posted pre mining and notifies of pending logs.
type PendingLogsEvent struct {
	Logs []*types.Log
}

// PendingStateEvent is posted pre mining and notifies of pending state changes.
type PendingStateEvent struct{}

// NewMinedBlockEvent is posted when a block has been imported.
type NewMinedBlockEvent struct{ Block *types.Block }

// RemovedTransactionEvent is posted when a reorg happens
type RemovedTransactionEvent struct{ Txs types.Transactions }

// RemovedLogsEvent is posted when a reorg happens
type RemovedLogsEvent struct{ Logs []*types.Log }

type ChainEvent struct {
	Block *types.Block
	Hash  common.Hash
	Logs  []*types.Log
}

type ChainSideEvent struct {
	Block *types.Block
}

type ChainHeadEvent struct{ Block *types.Block }
