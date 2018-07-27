package konsensus

import (
	"math/big"

	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/core/state"
	"github.com/kowala-tech/kcoin/client/crypto/sha3"
)

var (
	prevPriceIdx         = common.BytesToHash([]byte{0})
	priceIdx             = common.BytesToHash([]byte{1})
	supplyIdx            = common.BytesToHash([]byte{2})
	mintedRewardIdx      = common.BytesToHash([]byte{3})
	avgPriceIdx          = common.BytesToHash([]byte{6})
	hasSubmittedPriceIdx = common.BytesToHash([]byte{2})
)

type system struct {
	*state.StateDB
	SystemVarsReader
	provider PriceProvider
}

func Sys(state *state.StateDB, reader SystemVarsReader, provider PriceProvider) System {
	return &system{
		StateDB:          state,
		SystemVarsReader: reader,
		provider:         provider,
	}
}

func (sys *system) SetPrice(price *big.Int) error {
	// update prev price
	sys.SetState(sys.Address(), prevPriceIdx, sys.GetState(sys.Address(), priceIdx))
	// update current price
	sys.SetState(sys.Address(), priceIdx, common.BytesToHash(price.Bytes()))
	// reset average price
	sys.SetState(sys.Address(), avgPriceIdx, common.BytesToHash(common.Big0.Bytes()))
	// reset oracles state - hasSubmittedPrice
	keccak := sha3.NewKeccak256()
	participants, err := sys.provider.Submissions()
	if err != nil {
		return err
	}
	for _, oracle := range participants {
		// oracle contract key (oracleRegistry)
		keccak.Reset()
		keccak.Write(oracle.Bytes())
		keccak.Write(hasSubmittedPriceIdx.Bytes())
		key := common.BytesToHash(keccak.Sum(nil))

		// reset hasSubmittedPrice per submission
		sys.SetState(sys.provider.Address(), key, common.BytesToHash([]byte{0}))
	}
	return nil
}

func (sys *system) Mint(addr common.Address, amount *big.Int) {
	// add balance
	sys.AddBalance(addr, amount)
	// increase supply
	sys.SetState(sys.Address(), supplyIdx, common.BytesToHash(new(big.Int).Add(sys.GetState(sys.Address(), supplyIdx).Big(), amount).Bytes()))
	// increase minted reward per block
	sys.SetState(sys.Address(), mintedRewardIdx, common.BytesToHash(amount.Bytes()))
}

func (sys *system) Transfer(dest common.Address, src common.Address, amount *big.Int) {
	sys.AddBalance(dest, amount)
	sys.SubBalance(src, amount)
}

func (sys *system) OracleFund() common.Address   { return sys.Address() }
func (sys *system) PriceProvider() PriceProvider { return sys.provider }
