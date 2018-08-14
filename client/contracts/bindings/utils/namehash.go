package utils

//go:generate solc --allow-paths ., --abi --bin --libraries Strings:0x7d55e20244765F0Dda1aC0b91BA2BA1c5AA9D270 --overwrite -o build ../../truffle/contracts/utils/NameHash.sol
//go:generate ../../../build/bin/abigen -abi build/NameHash.abi -bin build/NameHash.bin -pkg utils -type NameHash -out ./gen_namehash.go
