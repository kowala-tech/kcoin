package oracle

//go:generate solc --abi --bin --overwrite -o build github.com/kowala-tech/kcoin/contracts/=/usr/local/include/solidity/ contract/ValidatorManager.sol
//go:generate abigen -abi build/VMC.abi -bin build/VMC.bin -pkg consensus -type ValidatorManager -out ./gen_vmc.go
