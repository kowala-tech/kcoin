package oracle

//go:generate solc --abi --bin --overwrite -o build github.com/kowala-tech/kcoin/contracts/=/usr/local/include/solidity/ contracts/OracleManager.sol
//go:generate abigen -abi build/OracleManager.abi -bin build/OracleManager.bin -pkg oracle -type OracleManager -out ./gen_manager.go
