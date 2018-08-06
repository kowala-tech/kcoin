package testfiles

//go:generate solc --abi --bin --overwrite -o build contracts/Compatible.sol
//go:generate solc --abi --bin --overwrite -o build contracts/Incompatible.sol
//go:generate solc --abi --bin --overwrite -o build contracts/TokenMock.sol

//go:generate ../../../../build/bin/abigen -abi build/Compatible.abi -bin build/Compatible.bin -pkg testfiles -type Compatible -out ./gen_compatible.go
//go:generate ../../../../build/bin/abigen -abi build/Incompatible.abi -bin build/Incompatible.bin -pkg testfiles -type Incompatible -out ./gen_incompatible.go
//go:generate ../../../../build/bin/abigen -abi build/TokenMock.abi -bin build/TokenMock.bin -pkg testfiles -type TokenMock -out ./gen_tokenmock.go

const CustomFallback = "test()"
