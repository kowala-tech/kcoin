package network_test

import (
	"math/big"
	"testing"

	nc "github.com/kowala-tech/kUSD/contracts/network"
	"github.com/kowala-tech/kUSD/params"
)

func TestOwnership(t *testing.T) {
	// create a new test
	mt, err := newMusdTest()
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
	err = commitTx(&tx{
		to:     other.addr,
		amount: new(big.Int).Mul(big.NewInt(10), new(big.Int).SetUint64(params.Ether)),
		data:   []byte{},
	}, mt.owner.key, mt.sim)
	if err != nil {
		t.Error(err)
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
