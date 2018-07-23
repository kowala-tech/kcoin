package tendermint

import (
	"math/big"

	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/consensus"
	"github.com/kowala-tech/kcoin/client/core/state"
	"github.com/kowala-tech/kcoin/client/core/types"
	"github.com/kowala-tech/kcoin/client/params"
	"github.com/kowala-tech/kcoin/client/rpc"
)

type Tendermint struct {
	system    System
	oracleMgr OracleMgr
	config    *params.TendermintConfig // Consensus engine configuration parameters
	fakeMode  bool
}

func New(config *params.TendermintConfig, mgr OracleMgr, system System) *Tendermint {
	return &Tendermint{
		oracleMgr: mgr,
		system:    system,
		config:    config,
	}
}

func NewFaker() *Tendermint {
	return &Tendermint{fakeMode: true}
}

func (tm *Tendermint) Author(header *types.Header) (common.Address, error) {
	return header.Coinbase, nil
}

func (tm *Tendermint) VerifyHeader(chain consensus.ChainReader, header *types.Header, seal bool) error {
	return nil
}

func (tm *Tendermint) VerifyHeaders(chain consensus.ChainReader, headers []*types.Header, seals []bool) (chan<- struct{}, <-chan error) {
	// @TODO (rgeraldes) - temporary work around
	abort, results := make(chan struct{}), make(chan error, len(headers))
	for i := 0; i < len(headers); i++ {
		results <- nil
	}
	return abort, results
}

func (tm *Tendermint) VerifySeal(chain consensus.ChainReader, header *types.Header) error {
	return nil
}

func (tm *Tendermint) Prepare(chain consensus.ChainReader, header *types.Header) error {
	return nil
}

func (tm *Tendermint) Finalize(chain consensus.ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, commit *types.Commit, receipts []*types.Receipt) (*types.Block, error) {
	if err := tm.distributeRewards(state, header); err != nil {
		return nil, err
	}

	// Accumulate any block and oracle rewards and commit the final state root
	header.Root = state.IntermediateRoot(true)

	// Header seems complete, assemble into a block and return
	return types.NewBlock(header, txs, receipts, commit), nil
}

func (tm *Tendermint) Seal(chain consensus.ChainReader, block *types.Block, stop <-chan struct{}) (*types.Block, error) {
	return nil, nil
}

func (tm *Tendermint) APIs(chain consensus.ChainReader) []rpc.API {
	return nil
}

func (tm *Tendermint) distributeRewards(state *state.StateDB, header *types.Header) error {
	mintedAmount, err := tm.system.MintedAmount()
	if err != nil {
		return err
	}

	// oracle fund
	oracleDeduction, err := tm.system.OracleDeduction(mintedAmount)
	if err != nil {
		return err
	}
	state.AddBalance(tm.system.Address(), oracleDeduction)

	// mining rewards
	mintedReward := new(big.Int).Sub(mintedAmount, oracleDeduction)
	state.Mint(header.Coinbase, mintedReward)

	// system updates and oracle rewards
	if isEpochEnding(header.Number) {
		// oracle rewards
		submissions, err := tm.oracleMgr.Submissions()
		if err != nil {
			return err
		}
		oracleReward, err := tm.system.OracleReward(mintedAmount)
		if err != nil {
			return err
		}
		rewardPerOracle := new(big.Int).Div(oracleReward, new(big.Int).SetUint64(uint64(len(submissions))))
		for _, author := range submissions {
			transfer(state, tm.system.Address(), author, rewardPerOracle)
		}

		// update system price with the oracle's average price

		// update average price of the oracle manager to 0

		// clean submissions array and set each on of hasSubmittedPrice to false

	}

	// update prevMintedAmount

	return nil
}

func isEpochEnding(blockNumber *big.Int) bool {
	_, mod := new(big.Int).DivMod(blockNumber, params.OracleEpochDuration, new(big.Int))
	return mod.Cmp(common.Big0) == 0
}

func transfer(state *state.StateDB, sender, recipient common.Address, amount *big.Int) {
	state.SubBalance(sender, amount)
	state.AddBalance(recipient, amount)
}
