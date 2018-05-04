package nameservice

import (
	"math/big"

	"github.com/kowala-tech/kcoin/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/common"
	"github.com/kowala-tech/kcoin/params"
)

//go:generate solc --abi --bin --overwrite -o build github.com/kowala-tech/kcoin/contracts/=/usr/local/include/solidity/ contracts/SimpleNameService.sol
//go:generate abigen -abi build/SimpleNameService.abi -bin build/SimpleNameService.bin -pkg nameservice -type NameService -out ./gen_nameservice.go

// @TODO (rgeraldes) - include namespace call in ABIGEN
// create default namespace no need to create an instance per binding
// nameservice to use the proxy upgrade method and other contracts to use the name service upgrade methods

var mapChainIDToAddr = map[uint64]common.Address{
	params.TestnetChainConfig.ChainID.Uint64(): common.HexToAddress("0xfe9bed356e7bc4f7a8fc48cc19c958f4e640ac62"),
}

var service *NameService

// Init initializes the name service
func Init(contractBackend bind.ContractBackend, chainID *big.Int) error {
	nameservice, err := NewNameService(mapChainIDToAddr[chainID.Uint64()], contractBackend)
	if err != nil {
		return err
	}

	service = nameservice

	return nil
}

func lookup(domain string) (common.Address, error) {
	return service.Lookup(&bind.CallOpts{}, domain)
}
