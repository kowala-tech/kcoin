package network_test

import (
	"math/big"
	"testing"

	"github.com/kowala-tech/kUSD/accounts/abi/bind/backends"
	nc "github.com/kowala-tech/kUSD/contracts/network"
	"github.com/kowala-tech/kUSD/core"
	"github.com/kowala-tech/kUSD/params"
)

func bigExp(b int, e int) *big.Int {
	return new(big.Int).Exp(big.NewInt(int64(b)), big.NewInt(int64(e)), nil)
}

func TestPriceOracle(t *testing.T) {
	// create contract(s) owner key
	owner, err := newKey()
	if err != nil {
		t.Error(err)
		return
	}
	// create a simulated backend
	sim := backends.NewSimulatedBackend(
		core.GenesisAlloc{
			owner.addr: core.GenesisAccount{
				Balance: new(big.Int).Mul(big.NewInt(100), new(big.Int).SetUint64(params.Ether)),
			},
		})
	oneCrypto := bigExp(10, 18)
	oneFiat := bigExp(10, 4)
	// deploy PriceOracle contract
	priceOracleAddr, _, priceOracle, err := nc.DeployPriceOracleContract(
		owner.auth, sim,
		"kUSD", "kUSD", 18, oneCrypto,
		"US Dollar", "USD", 4, oneFiat,
	)
	if err != nil {
		t.Error(err)
		return
	}
	sim.Commit()
	// get contract storage
	poData := nc.PriceOracle{}
	if err = unmarshalState(sim, priceOracleAddr, &poData); err != nil {
		t.Error(err)
		return
	}
	// check price for 1 crypto
	fiat, err := priceOracle.PriceForOneCrypto(nil)
	if err != nil {
		t.Error(err)
		return
	}
	if fiat.Cmp(oneFiat) != 0 {
		t.Errorf("got a bad value for the price of 1 crypto coin: exp: %s, got: %s", oneFiat, fiat)
		return
	}
	// check contract storage
	if fp := poData.PriceForOneCrypto(); fiat.Cmp(fp) != 0 {
		t.Errorf("the value calculated from the contract storage mismatch: got %v, exp: %v", fp, fiat)
		return
	}
	// check price for 1 fiat
	crypto, err := priceOracle.PriceForOneFiat(nil)
	if err != nil {
		t.Error(err)
		return
	}
	if crypto.Cmp(oneCrypto) != 0 {
		t.Errorf("got a bad value: exp: %s, got: %s", oneCrypto, crypto)
		return
	}
	// check the contract storage
	if cp := poData.PriceForOneFiat(); crypto.Cmp(cp) != 0 {
		t.Errorf("the value calculated from the contract storage mismatch: got %v, exp: %v", cp, crypto)
		return
	}
}
