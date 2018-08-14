package utils

//go:generate solc --allow-paths ., --abi --bin --overwrite -o build ../../truffle/contracts/utils/Strings.sol
//go:generate ../../../build/bin/abigen -abi build/Strings.abi -bin build/Strings.bin -pkg utils -type Strings -out ./gen_strings.go
