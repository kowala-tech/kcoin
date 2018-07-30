package konsensus

import (
	"reflect"
	"math/big"

	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/core/state"
	"github.com/kowala-tech/kcoin/client/crypto/sha3"
	"github.com/pkg/errors"
)

type system struct {
	DomainResolver
	Minter
	Pricing
	*state.StateDB
}

func NewSystem(resolver DomainResolver) *system {
	if resolver == nil {
		resolver = ResolverFunc(hardcodedResolver)
	}

	sys := &system{
		DomainResolver: resolver,
	}

	minter := sys.wrapSupplyMetrics()(MinterFunc(sys.AddBalance))
	sys.Minter = minter

	pricing := sys.wrapResetOracleMgr()(PricerFunc(sys.setPrice))
	sys.Pricing = pricing

	return sys
}

func (sys *system) WithState(state *state.StateDB) {
	sys.StateDB = state
}

func (sys *system) wrapSupplyMetrics() MinterMiddleware {
	return func(minter Minter) Minter {
		fn := func(account common.Address, amount *big.Int) {
			vars := &SystemVars{
				// @TODO (rgeraldes)
				//CurrencySupply: new(big.Int).Add(gov.GetState(gov.Address(), supplyIdx).Big(), amount), 
			}
			vars.MintedReward = vars.CurrencySupply
			sys.save(vars)

			minter.Mint(account, amount)
		}

		return MinterFunc(fn)
	}
}

func (sys *system) wrapResetOracleMgr() PricingMiddleware {
	return func(pricing Pricing) Pricing {
		fn := func(price *big.Int) {
			oracleMgr := &OracleMgr{
				AveragePrice: common.Big0
				
			}

			// reset hasSubmittedPrice per author
			keccak := sha3.NewKeccak256()
			participants, err := sys.provider.Submissions()
			if err != nil {
				return err
			}
			for _, oracle := range participants {
				// oracle contract key (oracleRegistry)
				keccak.Write(oracle.Bytes())
				keccak.Write(hasSubmittedPriceIdx.Bytes())
				key := common.BytesToHash(keccak.Sum(nil))
				keccak.Reset()

				// reset hasSubmittedPrice per submission
				sys.SetState(sys.provider.Address(), key, common.BytesToHash([]byte{0}))

				// reset submissions entry
				sys.SetState(sys.provider.Address(), key, common.BytesToHash([]byte{}))
			}

			pricing.SetPrice(price)
		}
	}	
}

func (sys *system) save(storage interface{}) error {
	addr, err := sys.Resolve(domain)
	if err != nil {
		return errors.Wrap(err, "write failed")
	}

	storageTyp := reflect.TypeOf(storage)
	storageVal := reflect.ValueOf(storage)

	for i := 0; i < storageT.NumField(); i ++ {
		// omit empty by default
		if fieldVal := storageVal.Field(i); fieldVal == reflect.Zero(fieldVal.Type()).Interface() {
			continue
		}

		sys.SetState(addr, common.BytesToHash(big.NewInt(i).Bytes()), common.BytesToHash(data))
	}
}

func (sys *system) setPrice(price *big.Int) error {
	return sys.save(vars{CurrencyPrice: price})
}

func (sys *system) Transfer(dest common.Address, src common.Address, amount *big.Int) {
	sys.AddBalance(dest, amount)
	sys.SubBalance(src, amount)
}