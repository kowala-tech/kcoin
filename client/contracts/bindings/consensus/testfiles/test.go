package testfiles

//go:generate solc --abi --bin --overwrite -o build contracts/Compatible.sol
//go:generate solc --abi --bin --overwrite -o build contracts/Incompatible.sol
//go:generate solc --abi --bin --overwrite -o build contracts/TokenMock.sol

//go:generate ../../../../build/bin/abigen -abi build/Compatible.abi -bin build/Compatible.bin -pkg testfiles -type Compatible -out ./gen_compatible.go
//go:generate ../../../../build/bin/abigen -abi build/Incompatible.abi -bin build/Incompatible.bin -pkg testfiles -type Incompatible -out ./gen_incompatible.go
//go:generate ../../../../build/bin/abigen -abi build/TokenMock.abi -bin build/TokenMock.bin -pkg testfiles -type TokenMock -out ./gen_tokenmock.go
//go:generate solc --allow-paths ., --abi --bin --overwrite --libraries NameHash:0x2b0D8ac41bD7aF24160c1F3C5430F496116b2292 -o build github.com/kowala-tech/kcoin/client/contracts/=../../truffle/contracts openzeppelin-solidity/=../../../truffle/node_modules/openzeppelin-solidity/  ../../../truffle/contracts/consensus/mgr/ValidatorMgr.sol
//go:generate ../../../../build/bin/abigen -abi build/ValidatorMgr.abi -bin build/ValidatorMgr.bin -pkg testfiles -type ValidatorMgr -out ./gen_validator.go

const CustomFallback = "test()"
