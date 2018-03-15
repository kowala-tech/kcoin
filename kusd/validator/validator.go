package validator

import (
	"errors"
	"github.com/kowala-tech/kUSD/accounts"
	"github.com/kowala-tech/kUSD/core"
	"github.com/kowala-tech/kUSD/core/state"
	"github.com/kowala-tech/kUSD/core/types"
	"github.com/kowala-tech/kUSD/kusddb"
	"math/big"
)

var (
	ErrCantVoteNotValidating             = errors.New("can't vote, not validating")
	ErrCantSetCoinbaseOnStartedValidator = errors.New("can't set coinbase, already started validating")
	ErrCantSetDepositOnStartedValidator  = errors.New("can't set deposit, already started validating")
	ErrCantAddProposalNotValidating      = errors.New("can't add proposal, not validating")
	ErrCantAddBlockFragmentNotValidating = errors.New("can't add block fragment, not validating")
)

// Backend wraps all methods required for mining.
type Backend interface {
	BlockChain() *core.BlockChain
	TxPool() *core.TxPool
	ChainDb() kusddb.Database
}

type Validator interface {
	Start() (Validator, error)
	Stop() (Validator, error)
	SetExtra(extra []byte) error
	Validating() bool
	SetCoinbase(walletAccount accounts.WalletAccount) error
	SetDeposit(deposit uint64) error
	Pending() (*types.Block, *state.StateDB)
	PendingBlock() *types.Block
	AddProposal(proposal *types.Proposal) error
	AddVote(vote *types.Vote) error
	AddBlockFragment(blockNumber *big.Int, round uint64, fragment *types.BlockFragment) error
}

func NewValidator(context *context) Validator {
	return newSyncing(context)
}
