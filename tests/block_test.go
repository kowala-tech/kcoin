package tests

import (
	"math/big"
	"testing"

	"github.com/kowala-tech/kUSD/params"
)

func TestBlockchain(t *testing.T) {
	t.Parallel()

	bt := new(testMatcher)
	// General state tests are 'exported' as blockchain tests, but we can run them natively.
	bt.skipLoad(`^GeneralStateTests/`)
	// Skip random failures due to selfish mining test.
	bt.skipLoad(`bcForkUncle\.json/ForkUncle`)
	bt.skipLoad(`^bcMultiChainTest\.json/ChainAtoChainB_blockorder`)
	bt.skipLoad(`^bcTotalDifficultyTest\.json/(lotsOfLeafs|lotsOfBranches|sideChainWithMoreTransactions)$`)
	bt.skipLoad(`^bcMultiChainTest\.json/CallContractFromNotBestBlock`)
	// Expected failures:
	bt.fails(`(?i)metropolis`, "metropolis is not supported yet")
	bt.fails(`^TestNetwork/bcTheDaoTest\.json/(DaoTransactions$|DaoTransactions_UncleExtradata$)`, "issue in test")

	bt.config(`^TestNetwork/`, params.ChainConfig{
		HomesteadBlock: big.NewInt(5),
		DAOForkBlock:   big.NewInt(8),
		DAOForkSupport: true,
		EIP150Block:    big.NewInt(10),
		EIP155Block:    big.NewInt(10),
		EIP158Block:    big.NewInt(14),
		// MetropolisBlock: big.NewInt(16),
	})
	bt.config(`^RandomTests/.*EIP150`, params.ChainConfig{
		HomesteadBlock: big.NewInt(0),
		EIP150Block:    big.NewInt(0),
	})
	bt.config(`^RandomTests/.*EIP158`, params.ChainConfig{
		HomesteadBlock: big.NewInt(0),
		EIP150Block:    big.NewInt(0),
		EIP155Block:    big.NewInt(0),
		EIP158Block:    big.NewInt(0),
	})
	bt.config(`^RandomTests/`, params.ChainConfig{
		HomesteadBlock: big.NewInt(0),
		EIP150Block:    big.NewInt(10),
	})
	bt.config(`^Homestead/`, params.ChainConfig{
		HomesteadBlock: big.NewInt(0),
	})
	bt.config(`^EIP150/`, params.ChainConfig{
		HomesteadBlock: big.NewInt(0),
		EIP150Block:    big.NewInt(0),
	})
	bt.config(`^[^/]+\.json`, params.ChainConfig{
		HomesteadBlock: big.NewInt(1000000),
	})

	bt.walk(t, blockTestDir, func(t *testing.T, name string, test *BlockTest) {
		cfg := bt.findConfig(name)
		if err := bt.checkFailure(t, name, test.Run(cfg)); err != nil {
			t.Error(err)
		}
	})
}
