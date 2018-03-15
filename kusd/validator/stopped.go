package validator

import (
	"math/big"
	"github.com/kowala-tech/kUSD/accounts"
	"github.com/kowala-tech/kUSD/core/types"
	"github.com/kowala-tech/kUSD/core/state"
)

// stopped state of the consensus validator when stop was issued
type stopped struct {
	*context
}

func newStopped(context *context) *stopped {
	return &stopped{context}
}

func (st *stopped) Start() (Validator, error) {
	return newValidating(st.context), nil
}

func (st *stopped) Stop() (Validator, error) {
	return st, nil
}

func (st *stopped) SetExtra(extra []byte) error {
	st.extra = extra
	return nil
}

func (st *stopped) Validating() bool {
	return false
}

func (st *stopped) SetCoinbase(walletAccount accounts.WalletAccount) error {
	st.walletAccount = walletAccount
	return nil
}

func (st *stopped) SetDeposit(deposit uint64) error {
	st.deposit = deposit
	return nil
}

func (st *stopped) Pending() (*types.Block, *state.StateDB) {
	return nil, nil
}

func (st *stopped) PendingBlock() *types.Block {
	return nil
}

func (st *stopped) AddProposal(proposal *types.Proposal) error {
	return ErrCantAddProposalNotValidating
}

func (st *stopped) AddVote(vote *types.Vote) error {
	return ErrCantVoteNotValidating
}

func (st *stopped) AddBlockFragment(blockNumber *big.Int, round uint64, fragment *types.BlockFragment) error {
	return ErrCantAddBlockFragmentNotValidating
}
