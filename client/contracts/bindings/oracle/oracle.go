package oracle

//go:generate solc --allow-paths ., --abi --bin --overwrite -o build github.com/kowala-tech/kcoin/client/contracts/=../../truffle/contracts openzeppelin-solidity/=../../truffle/node_modules/openzeppelin-solidity/  ../../truffle/contracts/oracle/OracleMgr.sol
//go:generate ../../../build/bin/abigen -abi build/OracleMgr.abi -bin build/OracleMgr.bin -pkg oracle -type OracleMgr -out ./gen_manager.go

var (
	errNoAddress = errors.New("there isn't an address for the provided chain ID")
)

var mapOracleMgrToAddr = map[uint64]common.Address{
	params.TestnetChainConfig.ChainID.Uint64(): common.HexToAddress("0x80eDa603028fe504B57D14d947c8087c1798D800"),
}

type Manager interface {
	Price() (*big.Int, error)
	OracleCount() (uint64, error)
}

func Instance(contractBackend bind.ContractBackend, chainID *big.Int) (*OracleMgr, error) {
	addr, ok := mapOracleMgrToAddr[chainID.Uint64()]
	if !ok {
		return nil, errNoAddress
	}

	return NewOracleMgr(addr, contractBackend)
}