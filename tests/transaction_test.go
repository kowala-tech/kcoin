package tests

import (
	"math/big"
	"testing"

	"github.com/kowala-tech/kUSD/params"
)

func TestTransaction(t *testing.T) {
	t.Parallel()

	txt := new(testMatcher)
	txt.config(`^Homestead/`, params.ChainConfig{
		HomesteadBlock: big.NewInt(0),
	})
	txt.config(`^EIP155/`, params.ChainConfig{
		HomesteadBlock: big.NewInt(0),
		EIP150Block:    big.NewInt(0),
		EIP155Block:    big.NewInt(0),
		EIP158Block:    big.NewInt(0),
		ChainId:        big.NewInt(1),
	})
	txt.config(`^Metropolis/`, params.ChainConfig{
		HomesteadBlock:  big.NewInt(0),
		EIP150Block:     big.NewInt(0),
		EIP155Block:     big.NewInt(0),
		EIP158Block:     big.NewInt(0),
		MetropolisBlock: big.NewInt(0),
	})

	txt.walk(t, transactionTestDir, func(t *testing.T, name string, test *TransactionTest) {
		cfg := txt.findConfig(name)
		if err := txt.checkFailure(t, name, test.Run(cfg)); err != nil {
			t.Error(err)
		}
	})
}
