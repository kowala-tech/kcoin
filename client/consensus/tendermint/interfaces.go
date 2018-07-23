package tendermint

import (
	"math/big"

	"github.com/kowala-tech/kcoin/client/common"
)

type OracleMgr interface {
	AveragePrice() (*big.Int, error)
	Submissions() ([]common.Address, error)
}

type System interface {
	Address() common.Address
	PreviousPrice() (*big.Int, error)
	MintedAmount() (*big.Int, error)
	OracleDeduction(*big.Int) (*big.Int, error)
	OracleReward(*big.Int) (*big.Int, error)
}
