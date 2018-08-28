package utils

//go:generate solc --allow-paths ., --abi --bin --overwrite -o build ../../truffle/contracts/utils/Strings.sol
//go:generate ../../../build/bin/abigen -abi build/strings.abi -bin build/strings.bin -pkg utils -type strings -out ./gen_strings.go
//go:generate solc --allow-paths ., --abi --bin --libraries Strings:0x71d9bfa0be4CCDF8Dbd08A4A6629F3f7BEADAC4e --overwrite -o build ../../truffle/contracts/utils/NameHash.sol
//go:generate ../../../build/bin/abigen -abi build/NameHash.abi -bin build/NameHash.bin -pkg utils -type NameHash -out ./gen_namehash.go
