package validator

import (
	"math/big"
	"github.com/kowala-tech/kUSD/common"
	"github.com/kowala-tech/kUSD/core/types"
	"github.com/kowala-tech/kUSD/core/state"
	"github.com/kowala-tech/kUSD/accounts"
	"github.com/kowala-tech/kUSD/log"
	"errors"
)

var (
	ErrCantStartValidatingSyncing = errors.New("can't start validating, syncing")
)

// awaitingSync state of the consensus validator when is waiting for sync before start
type awaitingSync struct {
	*context
	synced bool
}

func newAwaitingSync(context *context) *awaitingSync {
	return &awaitingSync{context: context, synced: false}
}

func (as *awaitingSync) sync() {
	if err := SyncWaiter(as.eventMux); err != nil {
		log.Warn("Failed to sync with network", "err", err)
		return
	}
	as.synced = true
}

func (as *awaitingSync) Start() (Validator, error) {
	if as.synced == true {
		return newValidating(as.context), nil
	}
	return as, ErrCantStartValidatingSyncing
}

func (as *awaitingSync) Stop() (Validator, error) {
	return newStopped(as.context), nil
}

func (as *awaitingSync) SetExtra(extra []byte) error {
	as.extra = extra
	return nil
}

func (as *awaitingSync) Validating() bool {
	return false
}

func (as *awaitingSync) SetCoinbase(address common.Address) error {
	newWalletAccount, err := accounts.NewWalletAccount(as.walletAccount, accounts.Account{Address: address})
	if err != nil {
		return err
	}
	as.walletAccount = newWalletAccount
	return nil
}

func (as *awaitingSync) SetDeposit(deposit uint64) {
	as.deposit = deposit
}

func (as *awaitingSync) Pending() (*types.Block, *state.StateDB) {
	return nil, nil
}

func (as *awaitingSync) PendingBlock() *types.Block {
	return nil
}

func (as *awaitingSync) AddProposal(proposal *types.Proposal) error {
	return ErrCantAddProposalNotValidating
}

func (as *awaitingSync) AddVote(vote *types.Vote) error {
	return ErrCantVoteNotValidating
}

func (as *awaitingSync) AddBlockFragment(blockNumber *big.Int, round uint64, fragment *types.BlockFragment) error {
	return ErrCantAddBlockFragmentNotValidating
}
