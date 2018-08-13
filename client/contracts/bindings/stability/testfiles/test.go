package testfiles

//go:generate solc --abi --bin --overwrite -o build contracts/PriceProviderMock.sol

//go:generate ../../../../build/bin/abigen -abi build/PriceProviderMock.abi -bin build/PriceProviderMock.bin -pkg testfiles -type PriceProviderMock -out ./gen_priceprovidermock.go
