package utils

//go:generate solc --allow-paths ., --abi --bin --overwrite -o build ../../truffle/contracts/utils/Strings.sol
//go:generate ../../../build/bin/abigen -abi build/strings.abi -bin build/strings.bin -pkg utils -type strings -out ./gen_strings.go
//go:generate solc --allow-paths ., --abi --bin --libraries Strings:0x2d7465b88a0A5A1bBff2671C8ED78F7506465ddc --overwrite -o build ../../truffle/contracts/utils/NameHash.sol
//go:generate ../../../build/bin/abigen -abi build/NameHash.abi -bin build/NameHash.bin -pkg utils -type NameHash -out ./gen_namehash.go
