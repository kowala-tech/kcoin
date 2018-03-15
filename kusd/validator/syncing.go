package validator

import (
	"math/big"
	"github.com/kowala-tech/kUSD/accounts"
	"github.com/kowala-tech/kUSD/core/types"
	"github.com/kowala-tech/kUSD/core/state"
	"github.com/kowala-tech/kUSD/log"
	"errors"
)

var (
	ErrCantStartValidatingSyncing = errors.New("can't start validating, syncing")
)

// syncing state of the consensus validator when just syncing and not starting
type syncing struct {
	*context
	synced bool
	start  bool
}

func newSyncing(context *context) *syncing {
	syncing := &syncing{context: context, synced: false, start: false}
	go syncing.sync()
	return syncing
}

func (sync *syncing) sync() {
	if err := SyncWaiter(sync.eventMux); err != nil {
		log.Warn("Failed to sync with network", "err", err)
		return
	}
	sync.synced = true
}

func (sync *syncing) Start() (Validator, error) {
	sync.start = true
	if sync.synced == true {
		return newValidating(sync.context), nil
	}
	return sync, ErrCantStartValidatingSyncing
}

func (sync *syncing) Stop() (Validator, error) {
	return newStopped(sync.context), nil
}

func (sync *syncing) SetExtra(extra []byte) error {
	sync.extra = extra
	return nil
}

func (sync *syncing) Validating() bool {
	return false
}

func (sync *syncing) SetCoinbase(walletAccount accounts.WalletAccount) error {
	sync.walletAccount = walletAccount
	return nil
}

func (sync *syncing) SetDeposit(deposit uint64) error {
	sync.deposit = deposit
	return nil
}

func (sync *syncing) Pending() (*types.Block, *state.StateDB) {
	return nil, nil
}

func (sync *syncing) PendingBlock() *types.Block {
	return nil
}

func (sync *syncing) AddProposal(proposal *types.Proposal) error {
	return ErrCantAddProposalNotValidating
}

func (sync *syncing) AddVote(vote *types.Vote) error {
	return ErrCantVoteNotValidating
}

func (sync *syncing) AddBlockFragment(blockNumber *big.Int, round uint64, fragment *types.BlockFragment) error {
	return ErrCantAddBlockFragmentNotValidating
}
