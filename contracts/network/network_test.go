package network_test

import (
	"crypto/ecdsa"
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
	owner   *key
	sim     *backends.SimulatedBackend
	addr    common.Address
	network *network.NetworkContract
}

func NewNetworkTest(t *testing.T) (*NetworkTest, error) {
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
	addr, _, contract, err := network.DeployNetworkContract(owner.auth, sim)
	if err != nil {
		t.Error(err)
		return nil, err
	}
	sim.Commit()
	return &NetworkTest{
		owner:   owner,
		sim:     sim,
		addr:    addr,
		network: contract,
	}, nil
}

func TestNumberOfVoters(t *testing.T) {
	test, err := NewNetworkTest(t)
	if err != nil {
		t.Error(err)
		return
	}

	count, err := test.network.GetVoterCount(&bind.CallOpts{})
	if err != nil {
		t.Fatal(err)
	}

	if count.Cmp(big.NewInt(2)) != 0 {
		t.Error(count)
	}

}
