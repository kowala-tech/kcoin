package network

import (
	"math/big"

	"github.com/kowala-tech/kUSD/common"
	"github.com/kowala-tech/kUSD/core/state"
)

// Contracts data layout.

// Ownable contract.
type Ownable struct {
	ContractOwner common.Address
}

// ERC20Simple data layout.
type ERC20Simple struct {
	// Token name.
	Name string
	// Token symbol.
	Symbol string
	// Number of decimals places
	Decimals uint8
}

// MToken contract layout.
type MToken struct {
	Ownable
	ERC20Simple
	// Owned tokens by each address.
	OwnedTokens *state.Mapping
	// Total supply of tokens.
	TotalTokens *big.Int
	// Maximum supply of tokens.
	MaximumTokens *big.Int
	// Amount of tokens hold by delegates.
	DelegatesTokens *state.Mapping
	// Amount of tokens delegated.
	DelegatedTokens *state.Mapping
	// Amount of tokens delegated ( tokenDelegations[delegate][delegator] ).
	TokenDelegations *state.Mapping
}

// BalanceOf returns the available balance of the address (delegations included).
func (m *MToken) BalanceOf(addr common.Address) (*big.Int, error) {
	r := new(big.Int)
	if err := m.OwnedTokens.Get(addr, &r); err != nil {
		return nil, err
	}
	delegatedTo := new(big.Int)
	if err := m.DelegatesTokens.Get(addr, &delegatedTo); err != nil {
		return nil, err
	}
	delegatedFrom := new(big.Int)
	if err := m.DelegatedTokens.Get(addr, &delegatedFrom); err != nil {
		return nil, err
	}
	r = r.Add(r, delegatedTo)
	return r.Sub(r, delegatedFrom), nil
}

// Delegated returns the amount of tokens delegated from fromAddr to toAddr.
func (m *MToken) Delegated(fromAddr, toAddr common.Address) (*big.Int, error) {
	r := new(big.Int)
	t := state.NewEmptyMapping()
	if err := m.TokenDelegations.Get(toAddr, t); err != nil {
		return nil, err
	}
	if err := t.Get(fromAddr, r); err != nil {
		return nil, err
	}
	return r, nil
}

// NetworkContractsMap data layout.
type NetworkContractsMap struct {
	Ownable
	// mToken contract address.
	MToken common.Address
	// Oracle address.
	PriceOracle common.Address
	// Network stats address.
	NetworkStats common.Address
}

var mapAddress = common.HexToAddress("0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff")

func GetContractsMap(sdb *state.StateDB) (*NetworkContractsMap, error) {
	r := &NetworkContractsMap{}
	if err := sdb.UnmarshalState(mapAddress, r); err != nil {
		return nil, err
	}
	return r, nil
}

// GetMToken parses the mToken contract local storage.
func (cm *NetworkContractsMap) GetMToken(sdb *state.StateDB) (*MToken, error) {
	r := &MToken{}
	if err := sdb.UnmarshalState(cm.MToken, r); err != nil {
		return nil, err
	}
	return r, nil
}

// GetPriceOracle parses the PriceOracle contract local storage.
func (cm *NetworkContractsMap) GetPriceOracle(sdb *state.StateDB) (*PriceOracle, error) {
	r := &PriceOracle{}
	if err := sdb.UnmarshalState(cm.PriceOracle, r); err != nil {
		return nil, err
	}
	return r, nil
}

// GetNetworkStats parses the GetNetworkStats contract local storage.
func (cm *NetworkContractsMap) GetNetworkStats(sdb *state.StateDB) (*NetworkStats, error) {
	r := &NetworkStats{}
	if err := sdb.UnmarshalState(cm.NetworkStats, r); err != nil {
		return nil, err
	}
	return r, nil
}

// PriceOracle data layout
type PriceOracle struct {
	Ownable
	// Cryptocurrency name.
	CryptoName string
	// Cryptocurrency symbol.
	CryptoSymbol string
	// Cryptocurrency decimal places.
	CryptoDecimals uint8
	// Fiat name.
	FiatName string
	// Fiat symbol.
	FiatSymbol string
	// Fiat decimal places.
	FiatDecimals uint8
	// Amounts of both currencies used to establish a relationship.
	CryptoAmount *big.Int
	FiatAmount   *big.Int
}

// PriceForFiat returns the amount of crypto that fiatAmount corresponds to.
func (po *PriceOracle) PriceForFiat(fiatAmount *big.Int) *big.Int {
	t := new(big.Int).Mul(fiatAmount, po.CryptoAmount)
	return t.Div(t, po.FiatAmount)
}

// PriceForCrypto returns the amount of fiat that cryptoAmount corresponds to.
func (po *PriceOracle) PriceForCrypto(cryptoAmount *big.Int) *big.Int {
	t := new(big.Int).Mul(cryptoAmount, po.FiatAmount)
	return t.Div(t, po.CryptoAmount)
}

// NetworkStats data layout.
type NetworkStats struct {
	Ownable
	// Total mined wei.
	TotalMinedWei *big.Int
}
