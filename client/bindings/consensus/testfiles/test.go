package testfiles

//go:generate solc --abi --bin --overwrite -o build contracts/Compatible.sol
//go:generate solc --abi --bin --overwrite -o build contracts/Incompatible.sol

//go:generate ../../../build/bin/abigen -abi build/Compatible.abi -bin build/Compatible.bin -pkg testfiles -type Compatible -out ./gen_compatible.go
//go:generate ../../../build/bin/abigen -abi build/Incompatible.abi -bin build/Incompatible.bin -pkg testfiles -type Incompatible -out ./gen_incompatible.go

const CustomFallback = "test()"
