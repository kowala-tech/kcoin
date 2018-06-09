package oracle

import (
	"math/big"

	"github.com/kowala-tech/kcoin/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/common"
	"github.com/kowala-tech/kcoin/contracts"
	"github.com/kowala-tech/kcoin/params"
)

//go:generate solc --abi --bin --overwrite -o build github.com/kowala-tech/kcoin/contracts/=/usr/local/include/solidity/ contracts/OracleMgr.sol
//go:generate abigen -abi build/OracleMgr.abi -bin build/OracleMgr.bin -pkg oracle -type OracleMgr -out ./gen_manager.go

var mapOracleMgrToAddr = map[uint64]common.Address{
	params.TestnetChainConfig.ChainID.Uint64(): common.HexToAddress("0x80eDa603028fe504B57D14d947c8087c1798D800"),
}

// PriceFeed delivers the current price on-demand
type PriceFeed interface {
	CurrentPrice() (*big.Int, error)
}

type priceFeed = OracleMgr

// LoadPriceFeed returns the price feed bindings
func LoadPriceFeed(contractBackend bind.ContractBackend, chainID *big.Int) (*priceFeed, error) {
	addr, ok := mapOracleMgrToAddr[chainID.Uint64()]
	if !ok {
		return nil, contracts.ErrNoAddress
	}

	return NewOracleMgr(addr, contractBackend)
}

// CurrentPrice returns the available kUSD price at the time of the request
func (feed *priceFeed) CurrentPrice() (*big.Int, error) {
	return feed.Price(&bind.CallOpts{})
}
