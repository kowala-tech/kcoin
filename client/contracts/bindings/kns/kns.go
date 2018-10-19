package kns

//go:generate solc --allow-paths ., --abi --bin --overwrite -o build zos-lib/=../../truffle/node_modules/zos-lib/ ../../truffle/contracts/kns/KNSRegistry.sol
//go:generate solc --allow-paths ., --combined-json bin-runtime,srcmap-runtime --overwrite -o build/registry-combined zos-lib/=../../truffle/node_modules/zos-lib/ ../../truffle/contracts/kns/KNSRegistry.sol
//go:generate ../../../build/bin/abigen -abi build/KNSRegistry.abi -bin build/KNSRegistry.bin -srcmap build/registry-combined/combined.json -pkg kns -type KNSRegistry -out ./gen_registry.go
//go:generate solc --allow-paths ., --abi --bin --overwrite -o build zos-lib/=../../truffle/node_modules/zos-lib/ ../../truffle/contracts/kns/FIFSRegistrar.sol
//go:generate solc --allow-paths ., --combined-json bin-runtime,srcmap-runtime --overwrite -o build/registrar-combined/combined.json zos-lib/=../../truffle/node_modules/zos-lib/ ../../truffle/contracts/kns/FIFSRegistrar.sol
//go:generate ../../../build/bin/abigen -abi build/FIFSRegistrar.abi -bin build/FIFSRegistrar.bin -srcmap build/registrar-combined/combined.json -pkg kns -type FIFSRegistrar -out ./gen_registrar.go
//go:generate solc --allow-paths ., --abi --bin --overwrite -o build zos-lib/=../../truffle/node_modules/zos-lib/ ../../truffle/contracts/kns/PublicResolver.sol
//go:generate solc --allow-paths ., --combined-json bin-runtime,srcmap-runtime --overwrite -o build/combined-resolver/combined.json zos-lib/=../../truffle/node_modules/zos-lib/ ../../truffle/contracts/kns/PublicResolver.sol
//go:generate ../../../build/bin/abigen -abi build/PublicResolver.abi -bin build/PublicResolver.bin -srcmap build/combined-resolver/combined.json -pkg kns -type PublicResolver -out ./gen_resolver.go
