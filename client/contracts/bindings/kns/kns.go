package kns

//go:generate solc --allow-paths ., --abi --bin --overwrite -o build zos-lib/=../../truffle/node_modules/zos-lib/ ../../truffle/contracts/kns/KNSRegistry.sol
//go:generate ../../../build/bin/abigen -abi build/KNSRegistry.abi -bin build/KNSRegistry.bin -pkg kns -type KNSRegistry -out ./gen_registry.go
//go:generate solc --allow-paths ., --abi --bin --overwrite -o build zos-lib/=../../truffle/node_modules/zos-lib/ ../../truffle/contracts/kns/FIFSRegistrar.sol
//go:generate ../../../build/bin/abigen -abi build/FIFSRegistrar.abi -bin build/FIFSRegistrar.bin -pkg kns -type FIFSRegistrar -out ./gen_registrar.go
