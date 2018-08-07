package konsensus

import (
	"math/big"

	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/consensus"
	"github.com/kowala-tech/kcoin/client/core/state"
	"github.com/kowala-tech/kcoin/client/core/types"
	"github.com/kowala-tech/kcoin/client/params"
	"github.com/kowala-tech/kcoin/client/rpc"
)

var (
	AndromedaBlockReward *big.Int = new(big.Int).SetUint64(115740741e+5)
)

type Konsensus struct {
	config *params.KonsensusConfig
}

func New(config *params.KonsensusConfig) *Konsensus {
	return &Konsensus{config: config}
}

func (kss *Konsensus) Author(header *types.Header) (common.Address, error) {
	return header.Coinbase, nil
}

func (kss *Konsensus) VerifyHeader(chain consensus.ChainReader, header *types.Header, seal bool) error {
	return nil
}

func (kss *Konsensus) VerifyHeaders(chain consensus.ChainReader, headers []*types.Header, seals []bool) (chan<- struct{}, <-chan error) {
	// @TODO (rgeraldes) - temporary work around
	abort, results := make(chan struct{}), make(chan error, len(headers))
	for i := 0; i < len(headers); i++ {
		results <- nil
	}
	return abort, results
}

func (kss *Konsensus) VerifySeal(chain consensus.ChainReader, header *types.Header) error {
	return nil
}

func (kss *Konsensus) Prepare(chain consensus.ChainReader, header *types.Header) error {
	return nil
}

func (kss *Konsensus) Finalize(chain consensus.ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, commit *types.Commit, receipts []*types.Receipt) (*types.Block, error) {
	if err := AccumulateRewards(state, header); err != nil {
		return nil, err
	}

	// Accumulate any block and uncle rewards and commit the final state root
	header.Root = state.IntermediateRoot(true)

	// Header seems complete, assemble into a block and return
	return types.NewBlock(header, txs, receipts, commit), nil
}

func AccumulateRewards(state *state.StateDB, header *types.Header) error {
	blockReward := AndromedaBlockReward

	// accumulate the rewards for the validator
	reward := new(big.Int).Set(blockReward)
	state.AddBalance(header.Coinbase, reward)

	return nil
}

func (kss *Konsensus) Seal(chain consensus.ChainReader, block *types.Block, stop <-chan struct{}) (*types.Block, error) {
	return nil, nil
}

func (kss *Konsensus) APIs(chain consensus.ChainReader) []rpc.API {
	return nil
}
