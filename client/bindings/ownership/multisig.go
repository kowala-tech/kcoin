package ownership

//go:generate solc --allow-paths . --abi --bin --overwrite -o build ../../contracts/ownership/contracts/MultiSigWallet.sol
//go:generate ../../build/bin/abigen -abi build/MultiSigWallet.abi -bin build/MultiSigWallet.bin -pkg ownership -type MultiSigWallet -out ./gen_multisig.go
