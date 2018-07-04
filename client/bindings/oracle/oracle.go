package oracle

//go:generate solc --allow-paths ., --abi --bin --overwrite -o build github.com/kowala-tech/kcoin/client/contracts/=../../contracts openzeppelin-solidity/=../../node_modules/openzeppelin-solidity/  ../../contracts/oracle/contracts/OracleMgr.sol
//go:generate ../../build/bin/abigen -abi build/OracleMgr.abi -bin build/OracleMgr.bin -pkg oracle -type OracleMgr -out ./gen_manager.go
