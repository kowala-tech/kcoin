package network_test

import (
	"fmt"
	"math/big"
	"math/rand"
	"testing"

	"github.com/kowala-tech/kUSD/common"
	"github.com/kowala-tech/kUSD/params"
)

func TestMintability(t *testing.T) {
	// create a new test
	mt, err := newMusdTest()
	if err != nil {
		t.Error(err)
		return
	}
	// generate a new address
	holderKey, err := newKey()
	if err != nil {
		t.Error(err)
		return
	}
	// mint 1 token
	const N_TOKENS = 1
	if _, err := mt.musd.Mint(mt.owner.auth, holderKey.addr, big.NewInt(N_TOKENS)); err != nil {
		t.Error(err)
		return
	}
	mt.sim.Commit()
	// check
	holderBal, err := mt.musd.AvailableTo(nil, holderKey.addr)
	if err != nil {
		t.Error(err)
		return
	}
	if holderBal.Uint64() != N_TOKENS {
		t.Errorf("expecting %d tokens from the contract call, got %d", N_TOKENS, holderBal)
		return
	}
	// mint every left token
	maxSupply, err := mt.musd.MaximumSupply(nil)
	if err != nil {
		t.Error(err)
		return
	}
	tokenSupply, err := mt.musd.TotalSupply(nil)
	if err != nil {
		t.Error(err)
		return
	}
	tokensLeft := new(big.Int).Sub(maxSupply, tokenSupply)
	if _, err = mt.musd.Mint(mt.owner.auth, holderKey.addr, tokensLeft); err != nil {
		t.Error(err)
		return
	}
	mt.sim.Commit()
	// check
	ts, err := mt.musd.TotalSupply(nil)
	if err != nil {
		t.Error(err)
		return
	}
	if maxSupply.Cmp(ts) != 0 {
		t.Errorf("couldn't mint all tokens left: only %v minted, expected %v", ts, maxSupply)
	}
	// try to mint another token
	if _, err = mt.musd.Mint(mt.owner.auth, holderKey.addr, common.Big1); err != nil {
		t.Error(err)
		return
	}
	mt.sim.Commit()
	// check
	nts, err := mt.musd.TotalSupply(nil)
	if err != nil {
		t.Error(err)
		return
	}
	if ts.Cmp(nts) != 0 {
		t.Errorf("got a different number of token than the expected: got %v, exp %v", nts, ts)
		return
	}
}

type addrAmount struct {
	addr   common.Address
	amount int
}

func TestDelegationRevocation(t *testing.T) {
	// create a new test
	mt, err := newMusdTest()
	if err != nil {
		t.Error(err)
		return
	}
	// generate holders addresses
	const N_HOLDERS = 128
	holders := make([]*key, 0, N_HOLDERS)
	for i := 0; i < N_HOLDERS; i++ {
		k, err := newKey()
		if err != nil {
			t.Error(err)
			return
		}
		holders = append(holders, k)
	}
	// delegate a random amount of tokens to random addresses
	const N_TOKENS = 1000
	available := make(map[common.Address]int, N_HOLDERS)
	bigNTokens := big.NewInt(N_TOKENS)
	delegations := make(map[common.Address]*addrAmount, N_HOLDERS)
	for _, k := range holders {
		// send kUSD to the holder address
		err := commitTx(&tx{
			to:     k.addr,
			amount: new(big.Int).SetUint64(params.Ether),
			data:   []byte{},
		}, mt.owner.key, mt.sim)
		if err != nil {
			t.Error(err)
			return
		}
		// mint tokens to the holder address
		if _, err := mt.musd.Mint(mt.owner.auth, k.addr, bigNTokens); err != nil {
			t.Error(err)
			return
		}
		mt.sim.Commit()
		nTokens := rand.Intn(N_TOKENS + 1)
		destAddr := holders[rand.Intn(N_HOLDERS)].addr
		available[k.addr] += N_TOKENS - nTokens
		available[destAddr] += nTokens
		aa := &addrAmount{addr: destAddr, amount: nTokens}
		delegations[k.addr] = aa
		// delegate random amount of tokens
		if _, err = mt.musd.Delegate(k.auth, aa.addr, big.NewInt(int64(aa.amount))); err != nil {
			t.Error(err)
			return
		}
		mt.sim.Commit()
	}
	// check available amounts
	for a, nt := range available {
		bigNt, err := mt.musd.AvailableTo(nil, a)
		if err != nil {
			t.Error(err)
			return
		}
		if big.NewInt(int64(nt)).Cmp(bigNt) != 0 {
			t.Errorf("got a different amount of available tokens for address %s. got: %v, exp: %v", a.Hex(), bigNt, nt)
			return
		}
	}
	// pick two addresses
	delegator := holders[0]
	delegate := holders[1]
	availableDelegate, err := mt.musd.AvailableTo(nil, delegate.addr)
	if err != nil {
		t.Error(err)
		return
	}
	// try to delegate more than the owned tokens
	if _, err = mt.musd.Delegate(delegator.auth, delegate.addr, big.NewInt(N_TOKENS+1)); err != nil {
		t.Error(err)
		return
	}
	mt.sim.Commit()
	newAvailableDelegate, err := mt.musd.AvailableTo(nil, delegate.addr)
	if err != nil {
		t.Error(err)
		return
	}
	if newAvailableDelegate.Cmp(availableDelegate) != 0 {
		t.Errorf("should not be able to delegate more tokens than then amount owned: exp: %v got: %v", availableDelegate, newAvailableDelegate)
		return
	}
	// calculate available tokens that can be delegated
	availableDelegator := new(big.Int).Sub(
		big.NewInt(N_TOKENS),
		big.NewInt(int64(delegations[delegator.addr].amount)),
	)
	// try to delegate more than the available tokens for delegation
	moreThanAvailable := new(big.Int).Add(availableDelegator, common.Big1)
	if _, err = mt.musd.Delegate(delegator.auth, delegate.addr, moreThanAvailable); err != nil {
		t.Error(err)
		return
	}
	mt.sim.Commit()
	newAvailableDelegate, err = mt.musd.AvailableTo(nil, delegate.addr)
	if err != nil {
		t.Error(err)
		return
	}
	if newAvailableDelegate.Cmp(availableDelegate) != 0 {
		t.Errorf("should not be able to delegate more tokens than the amount available for delegation: exp: %v got: %v", availableDelegate, newAvailableDelegate)
		return
	}
	// try to revoke more tokens than the delegated amount
	del := delegations[delegator.addr]
	for _, i := range holders {
		if i.addr == del.addr {
			delegate = i
		}
	}
	if availableDelegate, err = mt.musd.AvailableTo(nil, del.addr); err != nil {
		t.Error(err)
		return
	}
	if _, err = mt.musd.Revoke(delegator.auth, del.addr, new(big.Int).Add(availableDelegate, common.Big1)); err != nil {
		t.Error(err)
		return
	}
	mt.sim.Commit()
	if newAvailableDelegate, err = mt.musd.AvailableTo(nil, del.addr); err != nil {
		t.Error(err)
		return
	}
	if newAvailableDelegate.Cmp(availableDelegate) != 0 {
		t.Errorf("should not be able to revoke more tokens than the delegated amount: exp: %v got: %v", availableDelegate, newAvailableDelegate)
		return
	}
	if availableDelegate, err = mt.musd.AvailableTo(nil, del.addr); err != nil {
		t.Error(err)
		return
	}
	bigDel := big.NewInt(int64(del.amount))
	availableDelegate.Sub(availableDelegate, bigDel)
	// try to revoke delegation
	if _, err = mt.musd.Revoke(delegator.auth, del.addr, bigDel); err != nil {
		t.Error(err)
		return
	}
	mt.sim.Commit()
	if newAvailableDelegate, err = mt.musd.AvailableTo(nil, del.addr); err != nil {
		t.Error(err)
		return
	}
	if newAvailableDelegate.Cmp(availableDelegate) != 0 {
		t.Errorf("should have revoked: exp: %v got: %v", availableDelegate, newAvailableDelegate)
		return
	}
}

func TestAddressesChanges(t *testing.T) {
	// create a new test
	mt, err := newMusdTest()
	if err != nil {
		t.Error(err)
		return
	}
	// generate and seed address triplets
	const N_ADDRESSES = 2
	addresses := make([]*addressTriplet, 0, N_ADDRESSES)
	var at *addressTriplet
	for i := 0; i < N_ADDRESSES; i++ {
		if at, err = newAddressTriplet(); err != nil {
			t.Error(err)
			return
		}
		// send 1 coin to the management addr
		err = commitTx(&tx{
			to:     at.management.addr,
			amount: new(big.Int).Mul(big.NewInt(1), new(big.Int).SetUint64(params.Ether)),
		}, mt.owner.key, mt.sim)
		if err != nil {
			t.Error(err)
			return
		}
		// send 1 coin to the mining addr
		err = commitTx(&tx{
			to:     at.mining.addr,
			amount: new(big.Int).Mul(big.NewInt(1), new(big.Int).SetUint64(params.Ether)),
		}, mt.owner.key, mt.sim)
		if err != nil {
			t.Error(err)
			return
		}
		addresses = append(addresses, at)
	}
	holderA := addresses[0]
	holderB := addresses[1]
	// propose a mining address and try to accept with a different management address
	if _, err := mt.musd.ProposeMiningAddress(holderA.management.auth, holderA.mining.addr); err != nil {
		t.Error(err)
		return
	}
	mt.sim.Commit()
	if _, err := mt.musd.AcceptMiningAddress(holderA.mining.auth, holderB.management.addr); err != nil {
		t.Error(err)
		return
	}
	mt.sim.Commit()
	retAddr, err := mt.musd.MinersOwners(nil, holderA.mining.addr)
	if err != nil {
		t.Error(err)
		return
	}
	if retAddr == holderA.management.addr {
		t.Error("mining address should have not been accepted")
		return
	}
	// propose a mining address and accept it
	if _, err := mt.musd.ProposeMiningAddress(holderA.management.auth, holderA.mining.addr); err != nil {
		t.Error(err)
		return
	}
	mt.sim.Commit()
	if _, err = mt.musd.AcceptMiningAddress(holderA.mining.auth, holderA.management.addr); err != nil {
		t.Error(err)
		return
	}
	mt.sim.Commit()
	if retAddr, err = mt.musd.MinersOwners(nil, holderA.mining.addr); err != nil {
		t.Error(err)
		return
	}
	if retAddr != holderA.management.addr {
		t.Error("mining address should have been accepted")
		return
	}
	// propose a receiver address and try to accept from the wrong address
	if _, err := mt.musd.ProposeReceiverAddress(holderA.management.auth, holderA.receiver.addr); err != nil {
		t.Error(err)
		return
	}
	mt.sim.Commit()
	if _, err := mt.musd.AcceptReceiverAddress(holderA.mining.auth, holderB.receiver.addr); err != nil {
		t.Error(err)
		return
	}
	mt.sim.Commit()
	if retAddr, err = mt.musd.MinersReceivers(nil, holderA.mining.addr); err != nil {
		t.Error(err)
		return
	}
	if retAddr == holderA.receiver.addr {
		t.Error("receiver address should not have been accepted")
		return
	}
	// propose a receiver address and accept it
	if _, err := mt.musd.ProposeReceiverAddress(holderA.management.auth, holderA.receiver.addr); err != nil {
		t.Error(err)
		return
	}
	mt.sim.Commit()
	if _, err := mt.musd.AcceptReceiverAddress(holderA.mining.auth, holderA.receiver.addr); err != nil {
		t.Error(err)
		return
	}
	mt.sim.Commit()
	if retAddr, err = mt.musd.MinersReceivers(nil, holderA.mining.addr); err != nil {
		t.Error(err)
		return
	}
	if retAddr != holderA.receiver.addr {
		t.Error("receiver address should have been accepted")
		return
	}
}

func checkHolders(mt *musdTest, holders map[common.Address]*key) error {
	nHoldersBig, err := mt.musd.NumberTokenHolders(nil)
	if err != nil {
		return err
	}
	nHolders := int(nHoldersBig.Int64())
	if len(holders) != nHolders {
		return fmt.Errorf("got a different number of holders. got: %v, expected: %v", nHolders, len(holders))
	}
	for i := 0; i < nHolders; i++ {
		addr, err := mt.musd.TokenHolders(nil, big.NewInt(int64(i)))
		if err != nil {
			return err
		}
		if _, ok := holders[addr]; !ok {
			return fmt.Errorf("address missing")
		}
	}
	return nil
}

func TestTokenHoldersAddresses(t *testing.T) {
	// create a new test
	mt, err := newMusdTest()
	if err != nil {
		t.Error(err)
		return
	}
	// generate holders addresses, send 1 coin each and mint 100 tokens
	const N_HOLDERS = 10
	holders := make(map[common.Address]*key, N_HOLDERS)
	for i := 0; i < N_HOLDERS; i++ {
		k, err := newKey()
		if err != nil {
			t.Error(err)
			return
		}
		holders[k.addr] = k
		err = commitTx(&tx{
			to:     k.addr,
			amount: new(big.Int).Mul(big.NewInt(1), new(big.Int).SetUint64(params.Ether)),
		}, mt.owner.key, mt.sim)
		if err != nil {
			t.Error(err)
			return
		}
		if _, err = mt.musd.Mint(mt.owner.auth, k.addr, big.NewInt(100)); err != nil {
			t.Error(err)
			return
		}
		mt.sim.Commit()
	}
	// check holders
	if err := checkHolders(mt, holders); err != nil {
		t.Error(err)
		return
	}
	// transfer all tokens to some holder
	var k1, k2 *key
	for _, k := range holders {
		if k1 == nil {
			k1 = k
			continue
		}
		k2 = k
		break
	}
	if _, err = mt.musd.Transfer(k1.auth, k2.addr, big.NewInt(100)); err != nil {
		t.Error(err)
		return
	}
	mt.sim.Commit()
	delete(holders, k1.addr)
	// check holders
	if err := checkHolders(mt, holders); err != nil {
		t.Error(err)
		return
	}
}
