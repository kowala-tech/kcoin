package network_test

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"testing"

	"github.com/kowala-tech/kUSD/accounts/abi/bind"
	"github.com/kowala-tech/kUSD/accounts/abi/bind/backends"
	"github.com/kowala-tech/kUSD/common"
	"github.com/kowala-tech/kUSD/contracts/network"
	"github.com/kowala-tech/kUSD/core"
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

type NetworkTest struct {
	owner    *key
	sim      *backends.SimulatedBackend
	addr     common.Address
	contract *network.NetworkContract
}

func NewNetworkTest(t *testing.T) (*NetworkTest, error) {
	owner, err := newKey()
	if err != nil {
		t.Error(err)
		return nil, err
	}
	sim := backends.NewSimulatedBackend(core.GenesisAlloc{
		owner.addr: core.GenesisAccount{
			Balance: new(big.Int).Mul(big.NewInt(10), new(big.Int).SetUint64(params.Ether)),
		},
	})

	addr, _, contract, err := network.DeployNetworkContract(owner.auth, sim)
	if err != nil {
		t.Error(err)
		return nil, err
	}

	sim.Commit()

	return &NetworkTest{
		owner:    owner,
		sim:      sim,
		addr:     addr,
		contract: contract,
	}, nil
}

func TestNumberOfVoters(t *testing.T) {
	test, err := NewNetworkTest(t)
	if err != nil {
		t.Error(err)
		return
	}

	/*
		state, _ := test.sim.StateAt(test.sim.CurrentBlock().Root())
		code3 := state.GetCode(test.addr)
		fmt.Println("Current block:" + test.sim.CurrentBlock().Number().String())
		fmt.Println("code3: " + string(code3))

		fmt.Println("contract address" + test.addr.String())
		code2, err := test.sim.CodeAt(context.TODO(), test.addr, big.NewInt(0))
		code, err := test.sim.CodeAt(context.TODO(), test.addr, test.sim.CurrentBlock().Number())
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(string(code))
		if len(code) == 0 || len(code2) == 0 {
			fmt.Println("Size of the code:", 0)
		}
	*/

	count, err := test.contract.GetVoterCount(&bind.CallOpts{Pending: false, From: test.owner.addr})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("count " + count.String())

	if count.Cmp(big.NewInt(2)) != 0 {
		t.Error(count)
	}

}
