package proxy

//go:generate solc --allow-paths ., --abi --bin --overwrite -o build openzeppelin-solidity/=../../truffle/node_modules/openzeppelin-solidity/ ../../truffle/node_modules/zos-lib/contracts/upgradeability/UpgradeabilityProxyFactory.sol
//go:generate solc --allow-paths ., --combined-json bin-runtime,srcmap-runtime --overwrite -o build/proxy-combined openzeppelin-solidity/=../../truffle/node_modules/openzeppelin-solidity/ ../../truffle/node_modules/zos-lib/contracts/upgradeability/UpgradeabilityProxyFactory.sol
//go:generate ../../../build/bin/abigen -abi build/UpgradeabilityProxyFactory.abi -bin build/UpgradeabilityProxyFactory.bin -srcmap build/proxy-combined/combined.json -pkg proxy -type UpgradeabilityProxyFactory -out ./gen_manager.go
