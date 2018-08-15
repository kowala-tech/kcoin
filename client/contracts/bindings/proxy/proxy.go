package proxy

//go:generate solc --allow-paths ., --abi --bin --overwrite -o build openzeppelin-solidity/=../../truffle/node_modules/openzeppelin-solidity/ ../../truffle/node_modules/zos-lib/contracts/upgradeability/UpgradeabilityProxyFactory.sol
//go:generate ../../../build/bin/abigen -abi build/UpgradeabilityProxyFactory.abi -bin build/UpgradeabilityProxyFactory.bin -pkg proxy -type UpgradeabilityProxyFactory -out ./gen_manager.go
