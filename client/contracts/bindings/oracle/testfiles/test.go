package testfiles

// We need a specific version of the Oracle build contract because we link libraries, and the library address is based on the deployment.
//go:generate solc --allow-paths ., --abi --bin --libraries Strings:0x878083F57956A74Df81798Dd7c0E8CCC08Ad65c7 --overwrite -o build ../../../truffle/contracts/utils/NameHash.sol
//go:generate ../../../../build/bin/abigen -abi build/NameHash.abi -bin build/NameHash.bin -pkg testfiles -type NameHash -out ./gen_namehash.go
//go:generate solc --allow-paths ., --abi --bin --overwrite --libraries NameHash:0x2b0D8ac41bD7aF24160c1F3C5430F496116b2292 -o build github.com/kowala-tech/kcoin/client/contracts/=../../truffle/contracts openzeppelin-solidity/=../../../truffle/node_modules/openzeppelin-solidity/  ../../../truffle/contracts/oracle/OracleMgr.sol
//go:generate ../../../../build/bin/abigen -abi build/OracleMgr.abi -bin build/OracleMgr.bin -pkg testfiles -type OracleMgr -out ./gen_manager.go

//go:generate solc --abi --bin --overwrite -o build contracts/ConsensusMock.sol
//go:generate ../../../../build/bin/abigen -abi build/ConsensusMock.abi -bin build/ConsensusMock.bin -pkg testfiles -type ConsensusMock -out ./gen_consensusmock.go
//go:generate solc --abi --bin --overwrite -o build contracts/DomainResolverMock.sol
//go:generate ../../../../build/bin/abigen -abi build/DomainResolverMock.abi -bin build/DomainResolverMock.bin -pkg testfiles -type DomainResolverMock -out ./gen_domainresolvermock.go
