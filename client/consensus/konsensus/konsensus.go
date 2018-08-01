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

type System interface {
	Mint()
}

type Konsensus struct {
	config   *params.KonsensusConfig
	system   System
	//provider PriceProvider
	//reader   SystemVarsReader
}

func New(config *params.KonsensusConfig, provider PriceProvider, reader SystemVarsReader) *Konsensus {
	return &Konsensus{
		config:   config,
		system: NewSystem()
	}
}

func (ks *Konsensus) Author(header *types.Header) (common.Address, error) {
	return header.Coinbase, nil
}

func (ks *Konsensus) Finalize(chain consensus.ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, commit *types.Commit, receipts []*types.Receipt) (*types.Block, error) {
	ks.system.WithState(state)
		
	mintedAmount, err := system.MintedAmount()
	if err != nil {
		return err
	}

	// oracle fund
	oracleDeduction, err := system.OracleDeduction(mintedAmount)
	system.Mint(system.OracleFund(), oracleDeduction)

	// mining reward
	miningReward := new(big.Int).Sub(mintedAmount, oracleDeduction)
	system.Mint(validator, miningReward)

	// update price and reward oracles
	if oracleEpochEnd(blockNumber) {
		submissions, err := system.PriceProvider().Submissions()
		if err != nil {
			return err
		}
		if len(submissions) != 0 {
			// reward oracle
			oracleReward, err := system.OracleReward()
			if err != nil {
				return err
			}
			rewardPerOracle := new(big.Int).Div(oracleReward, new(big.Int).SetUint64(uint64(len(submissions))))
			for _, oracle := range submissions {
				// transfer reward from the oracle fund to the oracle
				system.Transfer(system.OracleFund(), oracle, rewardPerOracle)
			}

			// update price
			newPrice, err := system.PriceProvider().AveragePrice()
			if err != nil {
				return err
			}
			system.SetPrice(newPrice)
		}
	}

	// commit the final state root
	header.Root = state.IntermediateRoot(true)

	// Header seems complete, assemble into a block and return
	return types.NewBlock(header, txs, receipts, commit), nil
}

func oracleEpochEnd(blockNumber *big.Int) bool {
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
