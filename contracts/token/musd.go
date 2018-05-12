package token

import (
	"math/big"

	"github.com/kowala-tech/kcoin/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/common"
	"github.com/kowala-tech/kcoin/params"
)

//go:generate solc --abi --bin --overwrite -o build github.com/kowala-tech/kcoin/contracts/=/usr/local/include/solidity/ contracts/MiningToken.sol
//go:generate abigen -abi build/MiningToken.abi -bin build/MiningToken.bin -pkg token -type MiningToken -out ./gen_mtoken.go

var MapChainIDToAddr = map[uint64]common.Address{
	params.TestnetChainConfig.ChainID.Uint64(): common.HexToAddress("0xfe9bed356e7bc4f7a8fc48cc19c958f4e640ac62"),
}

type MUSD struct {
	*MiningToken
}

func Instance(contractBackend bind.ContractBackend, chainID *big.Int) (*MUSD, error) {
	token, err := NewMiningToken(MapChainIDToAddr[chainID.Uint64()], contractBackend)
	if err != nil {
		return nil, err
	}

	return &MUSD{MiningToken: token} , nil
}
