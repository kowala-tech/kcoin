package ownership

//go:generate solc --abi --bin --overwrite -o build contracts/MultiSigWallet.sol
//go:generate abigen -abi build/MultiSigWallet.abi -bin build/MultiSigWallet.bin -pkg ownership -type MultiSigWallet -out ./gen_multisig.go
