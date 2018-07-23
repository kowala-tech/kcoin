package konsensus

import (
	"math/big"

	"github.com/kowala-tech/kcoin/client/common"
)

type OracleMgr interface {
	AveragePrice() (*big.Int, error)
	Submissions() ([]common.Address, error)
}

type System interface {
	MintedAmount() (*big.Int, error)
	OracleDeduction(*big.Int) (*big.Int, error)
	Address() common.Address
}
