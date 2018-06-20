package oracle

//go:generate solc --abi --bin --overwrite -o build github.com/kowala-tech/kcoin/client/contracts/=/usr/local/include/solidity/ contracts/OracleMgr.sol
//go:generate abigen -abi build/OracleMgr.abi -bin build/OracleMgr.bin -pkg oracle -type OracleMgr -out ./gen_manager.go
