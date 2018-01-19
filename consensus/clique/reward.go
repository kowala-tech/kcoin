package clique

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/kowala-tech/kUSD/common"
	nc "github.com/kowala-tech/kUSD/contracts/network"
	"github.com/kowala-tech/kUSD/core/state"
	"github.com/kowala-tech/kUSD/core/types"
	"github.com/kowala-tech/kUSD/params"
)

var (
	big42kUSD = new(big.Int).Mul(big.NewInt(42), new(big.Int).SetUint64(params.Ether))
	big82kUSD = new(big.Int).Mul(big.NewInt(82), new(big.Int).SetUint64(params.Ether))
	big10e14  = big.NewInt(100000000000000)
	big1k     = big.NewInt(1000)
	big100    = big.NewInt(100)
	big101    = big.NewInt(101)
)

func (c *Clique) accumulateRewards(sdb *state.StateDB, header *types.Header, addrs []common.Address) error {
	// TODO(hrosa): what to do with transactions fees ?
	// get contracts map
	cMap, err := nc.GetContractsMap(sdb)
	if err != nil {
		return err
	}
	// get mToken contract data
	mt, err := cMap.GetMToken(sdb)
	if err != nil {
		return err
	}
	po, err := cMap.GetPriceOracle(sdb)
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
	// TODO(hrosa): remove. on the mainnet, tokens already exist
	if totalTokens == 0 {
		panic("need tokens")
	}
	// calculate the block reward.
	reward, err := calculateBlockReward(header.Number, sdb)
	if err != nil {
		return err
	}
	// calculate the reward per token.
	rewardPerToken := new(big.Int).Div(
		reward,
		mt.MaximumTokens,
	)
	// calculate total reward
	realReward := new(big.Int).Mul(rewardPerToken, big.NewInt(totalTokens))

	coins, coinsRem := new(big.Int).DivMod(realReward, big.NewInt(1000000000000000000), new(big.Int))
	coinsRemStr := coinsRem.String()
	coinsPerToken, coinsPerTokenRem := new(big.Int).DivMod(rewardPerToken, big.NewInt(1000000000000000000), new(big.Int))
	coinsPerTokenRemStr := coinsPerTokenRem.String()
	fmt.Printf(
		">>>> reward(%s): %s.%s (%s.%s per token, %d tokens) == %v\n>>>>\n",
		header.Number,
		coins.String(),
		strings.Repeat("0", 18-len(coinsRemStr))+coinsRemStr,
		coinsPerToken.String(),
		strings.Repeat("0", 18-len(coinsPerTokenRemStr))+coinsPerTokenRemStr,
		totalTokens,
		rewardPerToken,
	)

	// distribute rewards
	for _, a := range addrs {
		bal, err := mt.BalanceOf(a)
		if err != nil {
			return err
		}
		bal.Mul(bal, rewardPerToken)
		sdb.AddBalance(a, bal)
	}
	// update network stats
	nStats, err := cMap.GetNetworkStats(sdb)
	if err != nil {
		return err
	}
	nStats.TotalSupplyWei.Add(nStats.TotalSupplyWei, realReward)
	nStats.LastBlockReward = reward
	nStats.LastPrice = po.PriceForOneCrypto()
	if err := cMap.SetNetworkStats(sdb, nStats); err != nil {
		return err
	}
	return nil
}

func calculateBlockReward(blockNumber *big.Int, sdb *state.StateDB) (*big.Int, error) {
	// block 0
	if blockNumber.Cmp(common.Big0) == 0 {
		return common.Big0, nil
	}
	// block 1
	if blockNumber.Cmp(common.Big1) == 0 {
		return big42kUSD, nil
	}
	// open contracts map
	cMap, err := nc.GetContractsMap(sdb)
	if err != nil {
		return nil, err
	}
	// get network stats (last price)
	nStats, err := cMap.GetNetworkStats(sdb)
	if err != nil {
		return nil, err
	}
	// get price oracle
	po, err := cMap.GetPriceOracle(sdb)
	if err != nil {
		return nil, err
	}
	// get current price
	curPrice := po.PriceForOneCrypto()
	// check price
	cmpCurLast := curPrice.Cmp(nStats.LastPrice)
	cmpLastOneFiat := nStats.LastPrice.Cmp(po.OneFiat())
	var r *big.Int
	// p(b-1) > 1
	if cmpCurLast >= 0 && cmpLastOneFiat > 0 {
		fmt.Println(">>>> reward: p(b) >= p(b - 1) > 1 == min(1.01 * reward(b - 1), cap(b))", curPrice, nStats.LastPrice, nStats.LastBlockReward, blockRewardCap(nStats.TotalSupplyWei))
		// min(1.01 * reward(b - 1), cap(b))
		r = bigMin(
			new(big.Int).Add(
				nStats.LastBlockReward,
				new(big.Int).Div(nStats.LastBlockReward, big100),
			),
			blockRewardCap(nStats.TotalSupplyWei))
	} else if cmpCurLast <= 0 && cmpLastOneFiat < 0 {
		fmt.Println(">>>> reward: p(b) <= p(b-1) < 1 == max(1/1.01 * reward(b - 1), 0.0001)", curPrice, nStats.LastPrice, nStats.LastBlockReward)
		// max(1/1.01 * reward(b - 1), 0.0001)
		r = bigMax(
			new(big.Int).Div(
				new(big.Int).Mul(
					nStats.LastBlockReward,
					big100,
				),
				big101,
			),
			big10e14,
		)
	}
	// otherwise => reward(b - 1)
	if r == nil {
		fmt.Println(">>>> reward: otherwise == reward(b - 1)", nStats.LastBlockReward)
		r = nStats.LastBlockReward // reward(b - 1)
	}
	return r, nil
}

func bigMax(b1, b2 *big.Int) *big.Int {
	fmt.Printf(">>>> bigMax(%s, %s) =", b1.String(), b2.String())
	if b1.Cmp(b2) < 0 {
		fmt.Printf("%s\n", b2.String())
		return b2
	}
	fmt.Printf("%s\n", b1.String())
	return b1
}

func bigMin(b1, b2 *big.Int) *big.Int {
	fmt.Printf(">>>> bigMin(%s, %s) =", b1.String(), b2.String())
	if b1.Cmp(b2) > 0 {
		fmt.Printf("%s\n", b2.String())
		return b2
	}
	fmt.Printf("%s\n", b1.String())
	return b1
}

func blockRewardCap(totalWei *big.Int) *big.Int {
	fmt.Printf(">>>> blockRewardCap(): ")
	return bigMax(new(big.Int).Div(totalWei, big1k), big82kUSD)
}

/*
Mechanism 1: Block Reward Algorithm

totalCoinSupply(b) refers to the total number of coins issued as of block number b.

The block reward cap, cap(b) is defined as:

cap(b) := max(0.0001 * totalCoinSupply(b - 1), 82)

The market price, is represented by p(b)

With all these concepts in place, we are ready to define the block reward, reward(b), as:

reward(0) => 0
reward(1) => 42
reward(b):
	p(b) >= p(b - 1) > 1 => min(1.01 * reward(b - 1), cap(b))
	p(b) < p(b - 1) < 1 => max(1/1.01 * reward(b - 1), 0.0001)
    otherwise => reward(b - 1)

The calculation of the block reward is split into four scenarios: initial, divergent-rising, divergent-falling and convergent.

During the initial scenario, which applies only to the first block, the block reward is set to the arbitrary value of 42.

Next, we consider the divergent-rising scenario, which occurs when, over the course of two consecutive blocks, the price of kUSD is over $1 and rising or flat. In this scenario, we set the block's reward to 1% more than the previous block's reward, subject to the block reward cap, which prevents prolonged periods of block reward increase from growing too quickly.

We initially hypothesized that, when large numbers of newly minted coins are earned by miners, a large portion of such coins will reach exchanges as market sell orders and drive down the price of kUSD. A detailed agent-based behavior model with multiple market scenarios supports this hypothesis (see Agent-Based Modeling below). For this reason, we posit that for the divergent-rising scenario, no further mechanism is needed to reduce the price to $1.

Analogously to the divergent-rising scenario, the divergent-falling scenario occurs when, over the course of two consecutive blocks, the market price of kUSD is under $1 and is falling or flat. When this happens, the divergent-falling portion of the block reward function states that we should set the block's reward to the the previous block reward divided by 1.01 (subject to a minimum of 0.0001 kUSD).

Repeated applications of the formula in the divergent-falling scenario during a prolonged drop in price can lower the block reward to nearly zero. For example, in this scenario, reducing the block reward from 1 kUSD to 0.0001 kUSD takes only 925 blocks (3.9 hours at 15 seconds per block). However, even a near-zero block reward may not be sufficient to raise the price of the coin if there is a large drop in coin demand during the same period. In the section Mechanism 2: Stability Fee below, we will address this insufficiency by introducing a way to materially reduce the total coin supply.

Finally, the convergent scenario occurs whenever neither of the other three scenarios occur—that is, when b > 1 and the price for the current block is exactly at $1 or is closer to $1 than it was for the previous block. In this scenario, we consider that the previous block's reward is working well, so we set the current block's reward to the same value.
*/
