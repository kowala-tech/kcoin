package konsensus

import (
	"math/big"

	"github.com/kowala-tech/kcoin/client/common"
)

type PriceProvider interface {
	AveragePrice() (*big.Int, error)
	Submissions() ([]common.Address, error)
}

type SystemVarsReader interface {
	MintedAmount() (*big.Int, error)
	OracleDeduction(*big.Int) (*big.Int, error)
	OracleReward() (*big.Int, error)
}
