package validator

import (
	"math/big"
	"github.com/kowala-tech/kUSD/common"
	"github.com/kowala-tech/kUSD/accounts"
	"github.com/kowala-tech/kUSD/core/types"
	"github.com/kowala-tech/kUSD/core/state"
)

// syncing state of the consensus validator when just syncing and not starting
type syncing struct {
	*context
}

func newSyncing(context *context) *syncing {
	return &syncing{context}
}

func (sync *syncing) Start() (Validator, error) {
	// check if has finished syncing
	return newAwaitingSync(sync.context), nil
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

func (sync *syncing) SetCoinbase(address common.Address) error {
	newWalletAccount, err := accounts.NewWalletAccount(sync.walletAccount, accounts.Account{Address: address})
	if err != nil {
		return err
	}
	sync.walletAccount = newWalletAccount
	return nil
}

func (sync *syncing) SetDeposit(deposit uint64) {
	sync.deposit = deposit
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
