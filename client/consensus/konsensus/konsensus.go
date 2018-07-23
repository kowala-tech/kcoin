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

type Konsensus struct {
	System
	config    *params.KonsensusConfig // Consensus engine configuration parameters
	oracleMgr OracleMgr
	fakeMode  bool
}

func New(config *params.KonsensusConfig, oracleMgr OracleMgr, sys System) *Konsensus {
	return &Konsensus{
		config:    config,
		System:    sys,
		oracleMgr: oracleMgr,
	}
}

func NewFaker() *Konsensus {
	return &Konsensus{fakeMode: true}
}

func (ks *Konsensus) Author(header *types.Header) (common.Address, error) {
	return header.Coinbase, nil
}

func (ks *Konsensus) Finalize(chain consensus.ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, commit *types.Commit, receipts []*types.Receipt) (*types.Block, error) {
	mintedAmount, err := ks.MintedAmount()
	if err != nil {
		return nil, err
	}

	// oracle fund
	oracleDeduction, err := ks.OracleDeduction(mintedAmount)
	if err != nil {
		return nil, err
	}
	state.AddBalance(ks.Address(), oracleDeduction)

	// mining rewards
	mintedReward := new(big.Int).Sub(mintedAmount, oracleDeduction)
	state.AddBalance(header.Coinbase, mintedReward)

	// system updates and oracle rewards
	if OracleEpochEnd(header.Number) {

	}

	// Accumulate any block and uncle rewards and commit the final state root
	header.Root = state.IntermediateRoot(true)

	// Header seems complete, assemble into a block and return
	return types.NewBlock(header, txs, receipts, commit), nil
}

func OracleEpochEnd(blockNumber *big.Int) bool {
	_, mod := new(big.Int).DivMod(blockNumber, params.OracleEpochDuration, new(big.Int))
	return mod.Cmp(common.Big0) == 0
}

func (ks *Konsensus) VerifyHeader(chain consensus.ChainReader, header *types.Header, seal bool) error {
	return nil
}

func (ks *Konsensus) VerifyHeaders(chain consensus.ChainReader, headers []*types.Header, seals []bool) (chan<- struct{}, <-chan error) {
	// @TODO (rgeraldes) - temporary work around
	abort, results := make(chan struct{}), make(chan error, len(headers))
	for i := 0; i < len(headers); i++ {
		results <- nil
	}
	return abort, results
}

func (ks *Konsensus) VerifySeal(chain consensus.ChainReader, header *types.Header) error {
	return nil
}

func (ks *Konsensus) Prepare(chain consensus.ChainReader, header *types.Header) error {
	return nil
}

func (ks *Konsensus) Seal(chain consensus.ChainReader, block *types.Block, stop <-chan struct{}) (*types.Block, error) {
	return nil, nil
}

func (ks *Konsensus) APIs(chain consensus.ChainReader) []rpc.API {
	return nil
}
