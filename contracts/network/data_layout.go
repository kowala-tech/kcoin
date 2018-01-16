package network

//go:generate solc --abi --bin --overwrite -o build contracts/mUSD.sol
//go:generate abigen -abi build/mUSD.abi -bin build/mUSD.bin -pkg network -type MusdContract -out mUSD_generated.go
//go:generate solc --abi --bin --overwrite -o build contracts/network-stats.sol
//go:generate abigen -abi build/NetworkStats.abi -bin build/NetworkStats.bin -pkg network -type NetworkStatsContract -out network_stats_generated.go
//go:generate solc --abi --bin --overwrite -o build contracts/network-contracts-map.sol
//go:generate abigen -abi build/Contracts.abi -bin build/Contracts.bin -pkg network -type NetworkContractsMapContract -out network_contracts_map_generated.go
//go:generate solc --abi --bin --overwrite -o build contracts/price_oracle.sol
//go:generate abigen -abi build/PriceOracle.abi -bin build/PriceOracle.bin -pkg network -type PriceOracleContract -out price_oracle_generated.go

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

// Contracts data layout.
type Contracts struct {
	Ownable
	// mToken contract address.
	MToken common.Address
	// Oracle address.
	PriceOracle common.Address
	// Network stats address.
	NetworkStats common.Address
}

var mapAddress = common.HexToAddress("0x0c6e12469d65f6135f4b6b32148abbb25bed7349")

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

func (cm *NetworkContractsMap) SetMToken(sdb *state.StateDB, mt *MToken) error {
	return sdb.MarshalState(cm.MToken, mt)
}

// GetPriceOracle parses the PriceOracle contract local storage.
func (cm *NetworkContractsMap) GetPriceOracle(sdb *state.StateDB) (*PriceOracle, error) {
	r := &PriceOracle{}
	if err := sdb.UnmarshalState(cm.PriceOracle, r); err != nil {
		return nil, err
	}
	return r, nil
}

func (cm *NetworkContractsMap) SetPriceOracle(sdb *state.StateDB, po *PriceOracle) error {
	return sdb.MarshalState(cm.PriceOracle, po)
}

// GetNetworkStats parses the GetNetworkStats contract local storage.
func (cm *NetworkContractsMap) GetNetworkStats(sdb *state.StateDB) (*NetworkStats, error) {
	r := &NetworkStats{}
	if err := sdb.UnmarshalState(cm.NetworkStats, r); err != nil {
		return nil, err
	}
	return r, nil
}

func (cm *NetworkContractsMap) SetNetworkStats(sdb *state.StateDB, ns *NetworkStats) error {
	return sdb.MarshalState(cm.NetworkStats, ns)
}

// OracleAllowedAddress data layout.
type OracleAllowedAddress struct {
	Allowed bool
	Name    string
}

// PriceOracle data layout.
type PriceOracle struct {
	// Ownable fields.
	Ownable
	// Cryptocurrency name.
	CryptoName string
	// Cryptocurrency symbol.
	CryptoSymbol string
	// Cryptocurrency decimal places.
	CryptoDecimals uint8 // 3
	// Fiat name.
	FiatName string
	// Fiat symbol.
	FiatSymbol string
	// Fiat decimal places.
	FiatDecimals uint8 // 6
	// Volume for crypto.
	VolCrypto *big.Int
	// Volume for fiat.
	VolFiat *big.Int
}

// PriceForCrypto returns the price in fiat for cryptoAmount.
func (po *PriceOracle) PriceForCrypto(cryptoAmount *big.Int) *big.Int {
	r := new(big.Int).Mul(po.VolFiat, cryptoAmount)
	return r.Div(r, po.VolCrypto)
}

var big10 = big.NewInt(10)

func (po *PriceOracle) OneCrypto() *big.Int {
	return new(big.Int).Exp(big10, big.NewInt(int64(po.CryptoDecimals)), nil)
}

// PriceForOneCrypto returns the price in fiat for 1 crypto.
func (po *PriceOracle) PriceForOneCrypto() *big.Int {
	return po.PriceForCrypto(po.OneCrypto())
}

// PriceForFiat returns the price in crypto for fiatAmount.
func (po *PriceOracle) PriceForFiat(fiatAmount *big.Int) *big.Int {
	r := new(big.Int).Mul(po.VolCrypto, fiatAmount)
	return r.Div(r, po.VolFiat)
}

func (po *PriceOracle) OneFiat() *big.Int {
	return new(big.Int).Exp(big10, big.NewInt(int64(po.FiatDecimals)), nil)
}

// PriceForOneFiat returns the price in crypto for 1 fiat.
func (po *PriceOracle) PriceForOneFiat() *big.Int {
	return po.PriceForFiat(po.OneFiat())
}

// Network data layout.
type Network struct {
	// Total supply of wei. Must be updated every block.
	TotalSupplyWei *big.Int
	// Reward calculated for the last block. Must be updated every block.
	LastBlockReward *big.Int
	// Price established by the price oracle for the last block. Must be updated every block.
	LastPrice *big.Int
	// Tendermint validators.
	TendermintValidators []common.Address
}
