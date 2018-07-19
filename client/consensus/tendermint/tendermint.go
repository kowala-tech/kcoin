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

var (
	systemVariables = common.BytesToAddress([]byte{9})

	oracleFund              = common.BytesToAddress([]byte{10})
	oracleDeductionFraction = new(big.Int).SetUint64(4)
	andromedaOracleReward   = common.Big1

	// Some weird constants to avoid constant memory allocs for them.
	big100 = new(big.Int).SetUint64(100)
)

type PriceProvider interface {
	CurrentPrice() (*big.Int, error)
	PreviousPrice() (*big.Int, error)
	Submissions() ([]common.Address, error)
}

type Tendermint struct {
	priceProvider PriceProvider
	config        *params.TendermintConfig // Consensus engine configuration parameters
	fakeMode      bool
}

func New(config *params.TendermintConfig, priceProvider PriceProvider) *Tendermint {
	return &Tendermint{
		priceProvider: priceProvider,
		config:        config,
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

	// Accumulate any block and uncle rewards and commit the final state root
	header.Root = state.IntermediateRoot(true)

	// Header seems complete, assemble into a block and return
	return types.NewBlock(header, txs, receipts, commit), nil
}

func (tm *Tendermint) distributeRewards(state *state.StateDB, header *types.Header) error {
	currentPrice, err := tm.priceProvider.CurrentPrice()
	if err != nil {
		return err
	}

	prevPrice, err := tm.priceProvider.PreviousPrice()
	if err != nil {
		return err
	}

	prevSupply := state.GetState(systemVariables, common.BytesToHash([]byte{0}))
	prevMintedAmount := state.GetState(systemVariables, common.BytesToHash([]byte{1}))

	mintedAmount := mintedAmount(header.Number, currentPrice, prevPrice, prevSupply.Big(), prevMintedAmount.Big())
	oracleDeduction := new(big.Int).Div(new(big.Int).Mul(oracleDeductionFraction, mintedAmount), big100)
	state.AddBalance(oracleFund, oracleDeduction)

	mintedReward := new(big.Int).Sub(mintedAmount, oracleDeduction)
	state.Mint(header.Coinbase, mintedReward)

	if _, mod := new(big.Int).DivMod(header.Number, params.OracleEpochDuration, new(big.Int)); mod.Cmp(common.Big0) == 0 {
		submissions, err := tm.priceProvider.Submissions()
		if err != nil {
			return err
		}
		oracleReward := new(big.Int).Div(common.Min(andromedaOracleReward, state.GetBalance(oracleFund)), new(big.Int).SetUint64(len(submissions)))
		for _, author := range submissions {
			transfer(state, oracleFund, author, oracleReward)
		}

		// clean submissions state

	}

	return nil
}

func (tm *Tendermint) Seal(chain consensus.ChainReader, block *types.Block, stop <-chan struct{}) (*types.Block, error) {
	return nil, nil
}

func (tm *Tendermint) APIs(chain consensus.ChainReader) []rpc.API {
	return nil
}
