package stability

//go:generate solc --allow-paths ., --abi --bin --overwrite -o build github.com/kowala-tech/kcoin/client/contracts/=../../truffle/contracts openzeppelin-solidity/=../../truffle/node_modules/openzeppelin-solidity/  ../../truffle/contracts/stability/Stability.sol
//go:generate ../../../build/bin/abigen -abi build/Stability.abi -bin build/Stability.bin -pkg stability -type Stability -out ./gen_stability.go
