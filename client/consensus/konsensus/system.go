package konsensus

import (
	"math/big"
	"reflect"

	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/core/state"
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

// Minter represts the person who mints money
type Minter interface {
	Mint(account common.Address, amount *big.Int)
}

type MinterFunc func(common.Address, *big.Int)

func (fn MinterFunc) Mint(account common.Address, amount *big.Int) {
	fn(account, amount)
}

type MinterMiddleware func(Minter) Minter

// DataMapper converts data between different systems (vm storage <> golang types)
type DataMapper interface {
	Update(attrs ...interface{})
	Get(out interface{})
}

type system struct {
	Minter
	DataMapper
}

func NewSystem() *system {
	sys := new(system)
	sys.Minter = sys.wrapSupplyMetrics()(MinterFunc(sys.AddBalance))

	return sys
}

func (sys *system) wrapSupplyMetrics() MinterMiddleware {
	return func(minter Minter) Minter {
		fn := func(account common.Address, amount *big.Int) {
			sysvars := sys.Get(core.SystemVars{})
			//sysvars.MintedReward 
			//vars.MintedReward = vars.CurrencySupply
			sys.Update(sysvars)

			minter.Mint(account, amount)
		}

		return MinterFunc(fn)
	}
}

func (sys *system) SetPrice(price *big.Int) {
	sysvars := sys.Get(core.SystemVars{})
	sysvars.Price = price
	sys.Update(sysvars)
}

func (sys *system) Transfer(dest common.Address, src common.Address, amount *big.Int) {
	sys.AddBalance(dest, amount)
	sys.SubBalance(src, amount)
}