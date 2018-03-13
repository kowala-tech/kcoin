package network_test

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/kowala-tech/kUSD/accounts/abi/bind"
	"github.com/kowala-tech/kUSD/accounts/abi/bind/backends"
	"github.com/kowala-tech/kUSD/common"
	nc "github.com/kowala-tech/kUSD/contracts/network"
	"github.com/kowala-tech/kUSD/core"
	"github.com/kowala-tech/kUSD/core/types"
	"github.com/kowala-tech/kUSD/crypto"
	"github.com/kowala-tech/kUSD/params"
)

var gasPrice = new(big.Int).Mul(big.NewInt(5000000000), big.NewInt(16))

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
	auth := bind.NewKeyedTransactor(k)
	auth.GasLimit = big.NewInt(4000000)
	return &key{
		key:  k,
		addr: crypto.PubkeyToAddress(k.PublicKey),
		auth: auth,
	}, nil
}

func unmarshalState(sim *backends.SimulatedBackend, addr common.Address, v interface{}) error {
	sdb, err := sim.State()
	if err != nil {
		return err
	}
	return sdb.UnmarshalState(addr, v)
}

type musdTest struct {
	owner *key
	sim   *backends.SimulatedBackend
	addr  common.Address
	musd  *nc.MusdContract
}

func newMusdTest() (*musdTest, error) {
	owner, err := newKey()
	if err != nil {
		return nil, err
	}
	sim := backends.NewSimulatedBackend(core.GenesisAlloc{
		owner.addr: core.GenesisAccount{
			Balance: new(big.Int).Mul(big.NewInt(1000000), new(big.Int).SetUint64(params.Ether)),
		},
	})
	addr, _, musd, err := nc.DeployMusdContract(owner.auth, sim)
	if err != nil {
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

type addressTriplet struct {
	management *key
	mining     *key
	receiver   *key
}

func newAddressTriplet() (*addressTriplet, error) {
	t := make([]*key, 0, 3)
	for i := 0; i < 3; i++ {
		k, err := newKey()
		if err != nil {
			return nil, err
		}
		t = append(t, k)
	}
	return &addressTriplet{
		management: t[0],
		mining:     t[1],
		receiver:   t[2],
	}, nil
}

type tx struct {
	to     common.Address
	amount *big.Int
	data   []byte
}

func commitTx(tx *tx, key *ecdsa.PrivateKey, sim *backends.SimulatedBackend) error {
	gp, err := sim.SuggestGasPrice(nil)
	if err != nil {
		return err
	}
	nonce, err := sim.PendingNonceAt(nil, crypto.PubkeyToAddress(key.PublicKey))
	if err != nil {
		return err
	}
	t, err := types.SignTx(
		types.NewTransaction(
			nonce,
			tx.to,
			tx.amount,
			sim.CurrentBlock().GasLimit(),
			gp,
			tx.data,
		),
		types.UnprotectedSigner{},
		key,
	)
	if err != nil {
		return err
	}
	if err = sim.SendTransaction(nil, t); err != nil {
		return err
	}
	sim.Commit()
	return nil
}
