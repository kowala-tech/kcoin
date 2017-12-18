package network_test

//go:generate solc --abi --bin --overwrite -o build contracts/mUSD.sol
//go:generate abigen -abi build/mUSD.abi -bin build/mUSD.bin -pkg network -type MusdContract -out mUSD_generated.go
//go:generate solc --abi --bin --overwrite -o build contracts/network-stats.sol
//go:generate abigen -abi build/NetworkStats.abi -bin build/NetworkStats.bin -pkg network -type NetworkStatsContract -out network_stats_generated.go
//go:generate solc --abi --bin --overwrite -o build contracts/network-contracts-map.sol
//go:generate abigen -abi build/NetworkContractsMap.abi -bin build/NetworkContractsMap.bin -pkg network -type NetworkContractsMapContract -out network_contracts_map_generated.go
//go:generate solc --abi --bin --overwrite -o build contracts/price-oracle.sol
//go:generate abigen -abi build/PriceOracle.abi -bin build/PriceOracle.bin -pkg network -type PriceOracleContract -out price_oracle_generated.go

import (
	"crypto/ecdsa"
	"math/big"
	"testing"

	"github.com/kowala-tech/kUSD/accounts/abi/bind"
	"github.com/kowala-tech/kUSD/accounts/abi/bind/backends"
	"github.com/kowala-tech/kUSD/common"
	nc "github.com/kowala-tech/kUSD/contracts/network"
	"github.com/kowala-tech/kUSD/core"
	"github.com/kowala-tech/kUSD/core/types"
	"github.com/kowala-tech/kUSD/crypto"
	"github.com/kowala-tech/kUSD/params"
)

type key struct {
	key  *ecdsa.PrivateKey
	addr common.Address
	auth *bind.TransactOpts
}

func newKey() (*key, error) {
	k, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}
	return &key{
		key:  k,
		addr: crypto.PubkeyToAddress(k.PublicKey),
		auth: bind.NewKeyedTransactor(k),
	}, nil
}

type musdTest struct {
	owner *key
	sim   *backends.SimulatedBackend
	addr  common.Address
	musd  *nc.MusdContract
}

func newMusdTest(t *testing.T) (*musdTest, error) {
	owner, err := newKey()
	if err != nil {
		t.Error(err)
		return nil, err
	}
	sim := backends.NewSimulatedBackend(core.GenesisAlloc{
		owner.addr: core.GenesisAccount{
			Balance: new(big.Int).Mul(big.NewInt(100), new(big.Int).SetUint64(params.Ether)),
		},
	})
	addr, _, musd, err := nc.DeployMusdContract(owner.auth, sim)
	if err != nil {
		t.Error(err)
		return nil, err
	}
	sim.Commit()
	return &musdTest{
		owner: owner,
		sim:   sim,
		addr:  addr,
		musd:  musd,
	}, nil
}

var gasPrice = new(big.Int).Mul(big.NewInt(2500000000), big.NewInt(16))

func unmarshalState(sim *backends.SimulatedBackend, addr common.Address, v interface{}) error {
	sdb, err := sim.State()
	if err != nil {
		return err
	}
	return sdb.UnmarshalState(addr, v)
}

func TestOwnership(t *testing.T) {
	// create a new test
	mt, err := newMusdTest(t)
	if err != nil {
		t.Error(err)
		return
	}
	// create a new key
	other, err := newKey()
	if err != nil {
		t.Error(err)
		return
	}
	// seed the new key with kUSD
	tx, err := types.SignTx(
		types.NewTransaction(
			1,
			other.addr,
			new(big.Int).Mul(big.NewInt(10), new(big.Int).SetUint64(params.Ether)),
			mt.sim.CurrentBlock().GasLimit(),
			gasPrice,
			[]byte{},
		),
		types.HomesteadSigner{},
		mt.owner.key,
	)
	if err != nil {
		t.Error(err)
		return
	}
	if err = mt.sim.SendTransaction(nil, tx); err != nil {
		t.Error(err)
		return
	}
	mt.sim.Commit()
	// try to mint
	if _, err = mt.musd.MintTokens(other.auth, other.addr, common.Big1); err != nil {
		t.Error(err)
		return
	}
	mt.sim.Commit()
	//check balance
	bal, err := mt.musd.BalanceOf(nil, other.addr)
	if err != nil {
		t.Error(err)
		return
	}
	if bal.Uint64() != 0 {
		t.Error("able to mint without permission")
		return
	}
	// try to transfer ownership from a account different than the owner
	if _, err = mt.musd.TransferOwnership(other.auth, other.addr); err != nil {
		t.Error(err)
		return
	}
	mt.sim.Commit()
	// check owner
	sdb, err := mt.sim.State()
	if err != nil {
		t.Error(err)
		return
	}
	co := &nc.Ownable{}
	if err = sdb.UnmarshalState(mt.addr, co); err != nil {
		t.Error(err)
		return
	}
	if co.ContractOwner != mt.owner.addr {
		t.Errorf("got a different owner than expected: exp: %v, got: %v", mt.owner.addr.Hex(), co.ContractOwner.Hex())
		return
	}
	// try again from the owner account
	if _, err = mt.musd.TransferOwnership(mt.owner.auth, other.addr); err != nil {
		t.Error(err)
		return
	}
	mt.sim.Commit()
	// check owner
	if err = sdb.UnmarshalState(mt.addr, co); err != nil {
		t.Error(err)
		return
	}
	if co.ContractOwner != mt.owner.addr {
		t.Errorf("got a different owner than expected: exp: %v, got: %v", mt.owner.addr.Hex(), co.ContractOwner.Hex())
		return
	}
}

func TestMintableDelegableERC20Simple(t *testing.T) {
	// create a new test
	mt, err := newMusdTest(t)
	if err != nil {
		t.Error(err)
		return
	}
	// mint 1000 tokens for the owner
	const N_TOKENS = 1000
	if _, err := mt.musd.MintTokens(mt.owner.auth, mt.owner.addr, big.NewInt(N_TOKENS)); err != nil {
		t.Error(err)
		return
	}
	mt.sim.Commit()
	// check by calling the contract
	ownerBal, err := mt.musd.BalanceOf(nil, mt.owner.addr)
	if err != nil {
		t.Error(err)
		return
	}
	if ownerBal.Uint64() != N_TOKENS {
		t.Errorf("expecting %d tokens from the contract call, got %d", N_TOKENS, ownerBal)
		return
	}
	// check by parsing the local storage
	mtc := &nc.MToken{}
	if err = unmarshalState(mt.sim, mt.addr, mtc); err != nil {
		t.Error(err)
		return
	}
	if ownerBal, err = mtc.BalanceOf(mt.owner.addr); err != nil {
		t.Error(err)
		return
	}
	if ownerBal.Uint64() != N_TOKENS {
		t.Errorf("expecting %d tokens from the contract local storage, got %d", N_TOKENS, ownerBal)
		return
	}
	// create 10 token holders keys
	tokenHolders := make([]*key, 0, 10)
	for i := 0; i < 10; i++ {
		k, err := newKey()
		if err != nil {
			t.Error(err)
			return
		}
		tokenHolders = append(tokenHolders, k)
		// mint i+1 tokens
		nTokens := big.NewInt(int64(i) + 1)
		if _, err := mt.musd.MintTokens(mt.owner.auth, k.addr, nTokens); err != nil {
			t.Error(err)
			return
		}
		// transfer i+1 tokens
		if _, err := mt.musd.Transfer(mt.owner.auth, k.addr, nTokens); err != nil {
			t.Error(err)
			return
		}
		// delegate i+1 tokens
		if _, err := mt.musd.Delegate(mt.owner.auth, k.addr, nTokens); err != nil {
			t.Error(err)
			return
		}
		mt.sim.Commit()
		// check balance
		bal, err := mt.musd.BalanceOf(nil, k.addr)
		if err != nil {
			t.Error(err)
			return
		}
		expTokens := new(big.Int).Mul(nTokens, big.NewInt(3))
		if expTokens.Cmp(bal) != 0 {
			t.Errorf("error calling the contract. expected %s tokens, got %s", expTokens, bal)
			return
		}
		// check local storage
		if err = unmarshalState(mt.sim, mt.addr, mtc); err != nil {
			t.Error(err)
			return
		}
		if bal, err = mtc.BalanceOf(k.addr); err != nil {
			t.Error(err)
			return
		}
		if expTokens.Cmp(bal) != 0 {
			t.Errorf("error parsing local storage. expected %s tokens, got %s", expTokens, bal)
			return
		}
	}
	// check balance
	bal, err := mt.musd.BalanceOf(nil, mt.owner.addr)
	if err != nil {
		t.Error(err)
		return
	}
	expTokens := uint64(890)
	if bal.Uint64() != expTokens {
		t.Errorf("error calling contract. expected %d tokens, got %s", expTokens, bal)
		return
	}
	// check local storage
	if err = unmarshalState(mt.sim, mt.addr, mtc); err != nil {
		t.Error(err)
		return
	}
	if bal, err = mtc.BalanceOf(mt.owner.addr); err != nil {
		t.Error(err)
		return
	}
	if bal.Uint64() != expTokens {
		t.Errorf("error parsing local storage. expected %d tokens, got %s", expTokens, bal)
		return
	}
	// revoke delegated tokens
	for i, k := range tokenHolders {
		if _, err = mt.musd.Revoke(mt.owner.auth, k.addr, big.NewInt(int64(i)+1)); err != nil {
			t.Error(err)
			return
		}
	}
	mt.sim.Commit()
	// check balance
	if bal, err = mt.musd.BalanceOf(nil, mt.owner.addr); err != nil {
		t.Error(err)
		return
	}
	expTokens += 55
	if bal.Uint64() != expTokens {
		t.Errorf("error calling contract: expecting a balance of %d tokens after revocations, got %s", expTokens, bal)
		return
	}
	// check local storage
	if err = unmarshalState(mt.sim, mt.addr, mtc); err != nil {
		t.Error(err)
		return
	}
	if bal, err = mtc.BalanceOf(mt.owner.addr); err != nil {
		t.Error(err)
		return
	}
	if bal.Uint64() != expTokens {
		t.Errorf("error parsing local storage: expecting a balance of %d tokens after revocations, got %s", expTokens, bal)
		return
	}
}

const maxTokens = 1073741824

func TestMaxTokens(t *testing.T) {
	// create a new test
	mt, err := newMusdTest(t)
	if err != nil {
		t.Error(err)
		return
	}
	// mint the maximum tokens
	if _, err = mt.musd.MintTokens(mt.owner.auth, mt.owner.addr, big.NewInt(maxTokens)); err != nil {
		t.Error(err)
		return
	}
	mt.sim.Commit()
	// check balance
	bal, err := mt.musd.BalanceOf(nil, mt.owner.addr)
	if err != nil {
		t.Error(err)
		return
	}
	if bal.Uint64() != maxTokens {
		t.Errorf("expecting %d tokens, got %s", maxTokens, bal)
		return
	}
	// mint another tokens
	if _, err = mt.musd.MintTokens(mt.owner.auth, mt.owner.addr, common.Big1); err != nil {
		t.Error(err)
		return
	}
	mt.sim.Commit()
	// check balance
	if bal, err = mt.musd.BalanceOf(nil, mt.owner.addr); err != nil {
		t.Error(err)
		return
	}
	if bal.Uint64() != maxTokens {
		t.Errorf("expecting %d tokens, got %s", maxTokens, bal)
		return
	}
}

func TestPriceOracle(t *testing.T) {
	owner, err := newKey()
	if err != nil {
		t.Error(err)
		return
	}
	sim := backends.NewSimulatedBackend(core.GenesisAlloc{
		owner.addr: core.GenesisAccount{
			Balance: new(big.Int).Mul(big.NewInt(100), new(big.Int).SetUint64(params.Ether)),
		},
	})
	ca := new(big.Int).SetUint64(params.Ether)
	fa := big.NewInt(1 * 10000)
	addr, _, priceOracle, err := nc.DeployPriceOracleContract(owner.auth, sim, ca, fa)
	if err != nil {
		t.Error(err)
		return
	}
	sim.Commit()
	cPrice, err := priceOracle.PriceForFiat(nil, common.Big1)
	if err != nil {
		t.Error(err)
		return
	}
	poData := &nc.PriceOracle{}
	if err = unmarshalState(sim, addr, poData); err != nil {
		t.Error(err)
		return
	}
	if ca.Cmp(poData.CryptoAmount) != 0 {
		t.Errorf("the contract state has a value different than the initialized: %s %s", ca, poData.CryptoAmount)
		return
	}
	if fa.Cmp(poData.FiatAmount) != 0 {
		t.Errorf("the contract state has a value different than the initialized: %s %s", fa, poData.FiatAmount)
		return
	}
	if lsPrice := poData.PriceForFiat(common.Big1); lsPrice.Cmp(cPrice) != 0 {
		t.Errorf("got different prices for fiat: %s %s", cPrice, lsPrice)
		return
	}
	p := new(big.Int).SetUint64(params.Ether)
	if cPrice, err = priceOracle.PriceForCrypto(nil, p); err != nil {
		t.Error(err)
		return
	}
	if lsPrice := poData.PriceForCrypto(p); lsPrice.Cmp(cPrice) != 0 {
		t.Errorf("got different prices for crypto: %s %s", cPrice, lsPrice)
		return
	}
}

func TestDataLayouts(t *testing.T) {
	// create a new test
	mt, err := newMusdTest(t)
	if err != nil {
		t.Error(err)
		return
	}
	// create holders keys and mint some tokens
	const N_HOLDERS = 1024
	holders := make([]*key, 0, N_HOLDERS)
	var (
		k          *key
		tokenCount int64
	)
	for i := 0; i < N_HOLDERS; i++ {
		if k, err = newKey(); err != nil {
			t.Error(err)
			return
		}
		holders = append(holders, k)
		n := int64(i + 1)
		tokenCount += n
		if _, err := mt.musd.MintTokens(mt.owner.auth, k.addr, big.NewInt(n)); err != nil {
			t.Error(err)
			return
		}
		mt.sim.Commit()
	}
	// verify data layout
	mToken := &nc.MToken{}
	statedb, err := mt.sim.State()
	if err != nil {
		t.Error(err)
		return
	}
	if err = statedb.UnmarshalState(mt.addr, mToken); err != nil {
		t.Error(err)
		return
	}
	// contract owner
	if mToken.ContractOwner != mt.owner.addr {
		t.Errorf("got a different contract owner than expected. got: %s, exp: %s", mToken.ContractOwner.Hex(), mt.owner.addr.Hex())
		return
	}
	// token name
	if mToken.Name != "mUSD" {
		t.Errorf("got a bad token name: %s", mToken.Name)
		return
	}
	// token symbol
	if mToken.Symbol != "mUSD" {
		t.Errorf("got a bad symbol: %s", mToken.Symbol)
		return
	}
	// maximum tokens
	maxTokens := big.NewInt(1073741824)
	if mToken.MaximumTokens.Cmp(maxTokens) != 0 {
		t.Errorf("got a different maximum of tokens: %s", mToken.MaximumTokens)
		return
	}
	// owned tokens
	for i, h := range holders {
		bal, err := mt.musd.BalanceOf(nil, h.addr)
		if err != nil {
			t.Error(err)
			return
		}
		expBal := big.NewInt(int64(i + 1))
		if expBal.Cmp(bal) != 0 {
			t.Errorf("got a different balance than the expected for address %s. got: %s, exp: %s", h.addr.Hex(), bal, expBal)
			return
		}
	}
	noOwned, err := newKey()
	if err != nil {
		t.Error(err)
		return
	}
	// mint new tokens for the owner
	if _, err = mt.musd.MintTokens(mt.owner.auth, mt.owner.addr, big.NewInt(N_HOLDERS*100)); err != nil {
		t.Error(err)
		return
	}

	_ = noOwned
	// // delegate some tokens
	// delegatedTokens := make(map[common.Address]*big.Int, N_HOLDERS)
	// delegatesTokens := make(map[common.Address]*big.Int, N_HOLDERS+1)
	// delegations := make(map[common.Address]map[common.Address]*big.Int, N_HOLDERS+1)
	// delegatesTokens[noOwned.addr] = big.NewInt(0)
	// delegations[noOwned.addr] = make(map[common.Address]*big.Int, N_HOLDERS)
	// for i, h := range holders {
	// 	n := big.NewInt(int64(i + 1))
	// 	if _, err = mt.musd.Delegate(h.auth, noOwned.addr, common.Big1); err != nil {
	// 		t.Error(err)
	// 		return
	// 	}
	// 	mt.sim.Commit()
	// 	tk := delegatesTokens[noOwned.addr]
	// 	tk.Add(tk, common.Big1)
	// 	delegatedTokens[h.addr] = common.Big1
	// 	delegations[noOwned.addr][h.addr] = common.Big1
	// }
}
