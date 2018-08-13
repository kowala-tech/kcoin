package testfiles

//go:generate solc --abi --bin --overwrite -o build contracts/ConsensusMock.sol

//go:generate ../../../../build/bin/abigen -abi build/ConsensusMock.abi -bin build/ConsensusMock.bin -pkg testfiles -type ConsensusMock -out ./gen_consensusmock.go
