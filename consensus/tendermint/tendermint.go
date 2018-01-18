package tendermint

import (
	"math/big"

	"github.com/kowala-tech/kUSD/common"
	"github.com/kowala-tech/kUSD/consensus"
	"github.com/kowala-tech/kUSD/core/state"
	"github.com/kowala-tech/kUSD/core/types"
	"github.com/kowala-tech/kUSD/log"
	"github.com/kowala-tech/kUSD/params"
	"github.com/kowala-tech/kUSD/rpc"
)

// Tenderming proof-of-stake protocol constants.
var (
	blockReward *big.Int = big.NewInt(5e+18) // Block reward in wei for successfully mining a block
)

type Tendermint struct {
	config   *params.TendermintConfig // Consensus engine configuration parameters
	fakeMode bool
}

func New(config *params.TendermintConfig) *Tendermint {
	return &Tendermint{config: config}
}

func NewFaker() *Tendermint {
	return &Tendermint{fakeMode: true}
}

func (tendermint *Tendermint) Author(header *types.Header) (common.Address, error) {
	return header.Coinbase, nil
}

func (tendermint *Tendermint) VerifyHeader(chain consensus.ChainReader, header *types.Header, seal bool) error {
	return nil
}

func (tendermint *Tendermint) VerifyHeaders(chain consensus.ChainReader, headers []*types.Header, seals []bool) (chan<- struct{}, <-chan error) {
	abort := make(chan struct{})
	errorsOut := make(chan error)

	// @NOTE (rgeraldes) - the following work around is mandatory
	// because the block insertion process will wait forever
	// until it gets the results
	// @TODO (rgeraldes) - temp work around
	go func() {
		errorsOut <- nil
	}()

	return abort, errorsOut
}

func (tendermint *Tendermint) VerifySeal(chain consensus.ChainReader, header *types.Header) error {
	return nil
}

func (tendermint *Tendermint) Prepare(chain consensus.ChainReader, header *types.Header) error {
	return nil
}

func (tendermint *Tendermint) Finalize(chain consensus.ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, receipts []*types.Receipt, commit *types.Commit) (*types.Block, error) {
	log.Info("Finalising the block")

	/*
		// distribute block reward
		if tendermint.config.Rewarded {
			// get signers addresses
			// @TODO (rgeraldes)
			if err := AccumulateRewards(state, header, nil); err != nil {
				return nil, err
			}
		}
	*/

	// Accumulate any block and uncle rewards and commit the final state root
	header.Root = state.IntermediateRoot(true)
	return types.NewBlock(header, txs, receipts, commit), nil
}

func AccumulateRewards(state *state.StateDB, header *types.Header, addrs []common.Address) error {
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

func (tendermint *Tendermint) Seal(chain consensus.ChainReader, block *types.Block, stop <-chan struct{}) (*types.Block, error) {
	return nil, nil
}

func (tendermint *Tendermint) APIs(chain consensus.ChainReader) []rpc.API {
	return nil
}
