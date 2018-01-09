package tendermint

import (
	"math/big"

	"github.com/kowala-tech/kUSD/common"
	"github.com/kowala-tech/kUSD/consensus"
	"github.com/kowala-tech/kUSD/core/state"
	"github.com/kowala-tech/kUSD/core/types"
	"github.com/kowala-tech/kUSD/rpc"
)

// Tenderming proof-of-stake protocol constants.
var (
	blockReward *big.Int = big.NewInt(5e+18) // Block reward in wei for successfully mining a block
)

type Tendermint struct {
	fakeMode bool
}

func New() *Tendermint {
	return &Tendermint{}
}

func NewFaker() *Tendermint {
	return &Tendermint{fakeMode: true}
}

func (tendermint *Tendermint) Author(header *types.Header) (common.Address, error) {
	return common.Address{}, nil
}

func (tendermint *Tendermint) VerifyHeader(chain consensus.ChainReader, header *types.Header, seal bool) error {
	return nil
}

func (tendermint *Tendermint) VerifyHeaders(chain consensus.ChainReader, headers []*types.Header, seals []bool) (chan<- struct{}, <-chan error) {
	return nil, nil
}

func (tendermint *Tendermint) VerifyUncles(chain consensus.ChainReader, block *types.Block) error {
	return nil
}

func (tendermint *Tendermint) VerifySeal(chain consensus.ChainReader, header *types.Header) error {
	return nil
}

func (tendermint *Tendermint) Prepare(chain consensus.ChainReader, header *types.Header) error {
	return nil
}

func (tendermint *Tendermint) Finalize(chain consensus.ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, receipts []*types.Receipt, commit *types.Commit) (*types.Block, error) {
	// Accumulate any block and uncle rewards and commit the final state root
	AccumulateRewards(state, header)
	header.Root = state.IntermediateRoot(true)

	// Header seems complete, assemble into a block and return
	return types.NewBlock(header, txs, receipts, commit), nil
}

func (tendermint *Tendermint) Seal(chain consensus.ChainReader, block *types.Block, stop <-chan struct{}) (*types.Block, error) {
	return nil, nil
}

func (tendermint *Tendermint) APIs(chain consensus.ChainReader) []rpc.API {
	return nil
}

// AccumulateRewards credits the validators of the given block with the validation
// reward. The total reward consists of the static block reward.
func AccumulateRewards(state *state.StateDB, header *types.Header) {
	// @TODO (rgeraldes) - call Helio's contract (rewards)

	/*

		reward := new(big.Int).Set(blockReward)
		r := new(big.Int)
		for _, uncle := range uncles {
			r.Add(uncle.Number, big8)
			r.Sub(r, header.Number)
			r.Mul(r, blockReward)
			r.Div(r, big8)
			state.AddBalance(uncle.Coinbase, r)

			r.Div(blockReward, big32)
			reward.Add(reward, r)
		}
		state.AddBalance(header.Coinbase, reward)

	*/
}
