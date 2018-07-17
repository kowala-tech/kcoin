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
	oracleEpoch = new(big.Int) = SetUint64(900)
	oracleDeductionFraction = new(big.Int).SetUint64(4)
)

type PriceProvider interface {
	Price(pendingState bool) (*big.Int, error)
}

type Currency interface {
	PrevSupply() (*big.Int, error)
	PrevMintedAmount() (*big.Int, error)
	Address() common.Hash
}

type Tendermint struct {
	PriceProvider
	Currency

	config   *params.TendermintConfig // Consensus engine configuration parameters
	fakeMode bool
}

func New(config *params.TendermintConfig, priceProvider PriceProvider, currency Currency) *Tendermint {
	return &Tendermint{
		PriceProvider: priceProvider,
		Currency:      currency,
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
	if err := tm.accumulateRewards(state, header); err != nil {
		return nil, err
	}

	// Accumulate any block and uncle rewards and commit the final state root
	header.Root = state.IntermediateRoot(true)

	// Header seems complete, assemble into a block and return
	return types.NewBlock(header, txs, receipts, commit), nil
}

func (tm *Tendermint) accumulateRewards(state *state.StateDB, header *types.Header) error {
	currentPrice, err := tm.Price(true)
	if err != nil {
		return err
	}
	prevPrice, err := tm.Price(false)
	if err != nil {
		return err
	}
	prevSupply, err := tm.PrevSupply()
	if err != nil {
		return err
	}
	prevMintedAmount, err := tm.PrevMintedAmount()
	if err != nil {
		return err
	}

	mintedAmount := mintedAmount(header.Number, currentPrice, prevPrice, prevSupply, prevMintedAmount)

	oracleDeduction := new(big.Int).Div(new(big.Int).Mul(oracleDeductionFraction, mintedAmount), new(big.Int).SetUint64(100))
	state.AddBalance(common.BytesToAddress([]byte{0}), oracleDeduction)

	mintedReward := new(big.Int).Sub(mintedReward, oracleDeduction)
	reward := new(big.Int).Set(mintedReward)
	mint(state, header.Coinbase, reward)

	// reward oracles every 900 blocks
	// @TODO (What if timeouts?)
	if new(big.Int).DivMod(header.Number, oracleEpoch) == 0 {
		// check the list of submissions
		oracleReward := common.Min(baseOracleReward, oracleFundBalance)
		baseOracleReward := common.Big1
		// @TODO (rgeraldes) - use a core account
		oracleFundBalance := state.GetBalance(common.BytesToAddress([]byte{0}))

		// divide reward by all the participants

		// reward each one?

		// clean state back to empty
	}

	/*
		// @TODO (hrosa): what to do with transactions fees ?
		contracts, err := network.GetContracts(state)
		if err != nil {
			return err
		}
		// get mToken contract data
		mt, err := contracts.GetMToken(state)
		if err != nil {
			return err
		}
		// gather how many tokens each address holds
		addrsTokens := make(map[common.Address]int64, len(addrs))
		var totalTokens int64
		for _, a := range addrs {
			b, err := mt.BalanceOf(a)
			if err != nil {
				return err
			}
			bi := b.Int64()
			totalTokens += bi
			addrsTokens[a] = bi
		}
		// @TODO (hrosa): remove. on the mainnet, tokens already exist
		if totalTokens == 0 {
			return nil
		}
		// calculate the block reward.
		reward, err := CalculateBlockReward(header.Number, state)
		if err != nil {
			return err
		}
		coins, coinsRem := new(big.Int).DivMod(reward, big.NewInt(1000000000000000000), new(big.Int))
		coinsRemStr := coinsRem.String()
		fmt.Printf(">>>> reward(%s): %s.%s\n", header.Number, coins.String(), strings.Repeat("0", len(coinsRemStr))+coinsRemStr)
		// calculate the reward per token.
		rewardPerToken, remReward := new(big.Int).DivMod(
			reward,
			big.NewInt(totalTokens),
			new(big.Int),
		)
		// distribute rewards
		for _, a := range addrs {
			bal, err := mt.BalanceOf(a)
			if err != nil {
				return err
			}
			bal.Mul(bal, rewardPerToken)
			state.AddBalance(a, bal)
		}
		reward.Sub(reward, remReward)
		// update network stats
		networkInfo, err := contracts.GetNetworkContract(state)
		if err != nil {
			return err
		}
		// @TODO (hrosa): should be using a state writer
		reward.Sub(reward, remReward)
		w := common.BytesToHash(networkInfo.TotalSupplyWei.Add(networkInfo.TotalSupplyWei, reward).Bytes())
		state.SetState(contracts.Network, common.BytesToHash([]byte{0}), w)
	*/

	return nil
}

func (tm *Tendermint) Seal(chain consensus.ChainReader, block *types.Block, stop <-chan struct{}) (*types.Block, error) {
	return nil, nil
}

func (tm *Tendermint) APIs(chain consensus.ChainReader) []rpc.API {
	return nil
}

func mint(state *state.StateDB, addr common.Address, amount *big.Int) {
	state.Mint(addr, amount)
}
