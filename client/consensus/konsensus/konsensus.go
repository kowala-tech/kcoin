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
	prevPriceIdx    = common.BytesToHash([]byte{0})
	priceIdx        = common.BytesToHash([]byte{1})
	mintedAmountIdx = common.BytesToHash([]byte{3})
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
	if !ks.fakeMode {
		mintedAmount, err := ks.MintedAmount()
		if err != nil {
			return nil, err
		}

		// oracle fund
		oracleDeduction, err := ks.OracleDeduction(mintedAmount)
		if err != nil {
			return nil, err
		}
		state.Mint(ks.Address(), oracleDeduction)

		// mining rewards
		mintedReward := new(big.Int).Sub(mintedAmount, oracleDeduction)
		state.Mint(header.Coinbase, mintedReward)

		if OracleEpochEnd(header.Number) {
			// oracle rewards
			submissions, err := ks.oracleMgr.Submissions()
			if err != nil {
				return nil, err
			}
			oracleReward, err := ks.OracleReward(mintedAmount)
			if err != nil {
				return nil, err
			}
			rewardPerOracle := new(big.Int).Div(oracleReward, new(big.Int).SetUint64(uint64(len(submissions))))
			for _, oracle := range submissions {
				transfer(state, ks.Address(), oracle, rewardPerOracle)
			}

			// update prev price and current price
			averagePrice, err := ks.oracleMgr.AveragePrice()
			if err != nil {
				return nil, err
			}
			currentPrice, err := ks.CurrencyPrice()
			if err != nil {
				return nil, err
			}
			state.SetState(ks.Address(), prevPriceIdx, common.BytesToHash(currentPrice.Bytes()))
			state.SetState(ks.Address(), priceIdx, common.BytesToHash(averagePrice.Bytes()))
		}

		// update minted amount
		state.SetState(ks.Address(), mintedAmountIdx, common.BytesToHash(mintedAmount.Bytes()))
	}

	// commit the final state root
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

func transfer(state *state.StateDB, sender, recipient common.Address, amount *big.Int) {
	state.SubBalance(sender, amount)
	state.AddBalance(recipient, amount)
}
