package nameservice

//go:generate solc --abi --bin --overwrite -o build github.com/kowala-tech/kcoin/contracts/=/usr/local/include/solidity/ contracts/SimpleNameService.sol
//go:generate abigen -abi build/SimpleNameService.abi -bin build/SimpleNameService.bin -pkg nameservice -type NameService -out ./gen_nameservice.go

// @TODO (rgeraldes) - include namespace call in ABIGEN
// create default namespace no need to create an instance per binding
// nameservice to use the proxy upgrade method and other contracts to use the name service upgrade methods

var MapChainIDToAddr = map[uint64]common.Address{
	params.TestnetChainConfig.ChainID.Uint64(): common.HexToAddress("0xfe9bed356e7bc4f7a8fc48cc19c958f4e640ac62"),
}

var service *NameService

func Init(contractBackend bind.ContractBackend, chainID *big.Int) {
	service = NewNameService()
}

func lookup(domain string) (common.Address, error) {
	service.Lookup()
}