package network_test

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"testing"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/kowala-tech/contracts"
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
	sim := backends.NewSimulatedBackend(core.GenesisAlloc{
		owner.addr: core.GenesisAccount{
			Balance: new(big.Int).Mul(big.NewInt(100), new(big.Int).SetUint64(params.Ether)),
		},
	})
	// deploy mUSD contract
	mtAddr, _, _, err := contracts.DeployMusdContract(owner.auth, sim)
	if err != nil {
		t.Error(err)
		return
	}
	sim.Commit()
	// deploy PriceOracle contract
	_, _, priceOracle, err := contracts.DeployPriceOracleContract(
		owner.auth, sim,
		"kUSD", "kUSD", 18,
		"US Dollar", "USD", 4,
		mtAddr,
	)
	if err != nil {
		t.Error(err)
		return
	}
	sim.Commit()
	// create exchange key
	exchange, err := newKey()
	if err != nil {
		t.Error(err)
		return
	}
	exchange2, err := newKey()
	if err != nil {
		t.Error(err)
		return
	}
	// allow exchanges addresses
	if _, err = priceOracle.AllowAddress(
		owner.auth,
		exchange.addr,
		"Test Exchange (good)",
	); err != nil {
		t.Error(err)
		return
	}
	if _, err = priceOracle.AllowAddress(
		owner.auth,
		exchange2.addr,
		"Test Exchange (to be disallowed)",
	); err != nil {
		t.Error(err)
		return
	}
	// get nonce
	nonce, err := sim.PendingNonceAt(context.TODO(), owner.addr)
	if err != nil {
		t.Error(err)
		return
	}
	// send some coins to the exchanges
	addrs := []common.Address{exchange.addr, exchange2.addr}
	for _, a := range addrs {
		val := new(big.Int).Mul(big.NewInt(10), new(big.Int).SetUint64(params.Ether))
		gasLimit, err := sim.EstimateGas(context.TODO(), ethereum.CallMsg{
			From:     owner.addr,
			To:       &a,
			Value:    val,
			GasPrice: gasPrice,
		})
		if err != nil {
			t.Error(err)
			return
		}
		tx, err := types.SignTx(
			types.NewTransaction(
				nonce,
				a,
				val,
				gasLimit,
				gasPrice,
				[]byte{},
			),
			types.HomesteadSigner{},
			owner.key,
		)
		if err != nil {
			t.Error(err)
			return
		}
		if err = sim.SendTransaction(nil, tx); err != nil {
			t.Error(err)
			return
		}
		nonce++
		sim.Commit()
	}
	// disallow second exchange
	if _, err = priceOracle.DisallowAddress(
		owner.auth,
		exchange2.addr,
	); err != nil {
		t.Error(err)
		return
	}
	sim.Commit()
	// register transactions
	oneCrypto := bigExp(10, 18)
	oneFiat := bigExp(10, 4)
	blk := new(big.Int).Add(sim.CurrentBlock().Number(), common.Big1)
	if _, err = priceOracle.RegisterTransaction(
		exchange.auth,
		oneCrypto,
		oneFiat,
		blk,
	); err != nil {
		t.Error(err)
		return
	}
	sim.Commit()
	// this transaction shouldn't be considered
	if _, err = priceOracle.RegisterTransaction(
		exchange2.auth,
		common.Big1,
		common.Big2,
		blk,
	); err != nil {
		t.Error(err)
		return
	}
	sim.Commit()
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
}
