package oracle

//go:generate solc --allow-paths ., --abi --bin --overwrite -o build github.com/kowala-tech/kcoin/client/contracts/=/usr/local/include/solidity/ openzeppelin-solidity/=/usr/local/include/solidity/openzeppelin-solidity/  ../../contracts/oracle/contracts/OracleMgr.sol
//go:generate abigen -abi build/OracleMgr.abi -bin build/OracleMgr.bin -pkg oracle -type OracleMgr -out ./gen_manager.go
