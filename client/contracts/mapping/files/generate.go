package files

//go:generate solc --allow-paths ., --abi --bin --overwrite -o build github.com/kowala-tech/kcoin/client/contracts/=../../truffle/contracts openzeppelin-solidity/=../../truffle/node_modules/openzeppelin-solidity/ zos-lib/=../../truffle/node_modules/zos-lib/ ../../truffle/contracts/sysvars/SystemVars.sol
//go:generate solc --allow-paths .,  --combined-json bin-runtime,srcmap-runtime --overwrite -o build github.com/kowala-tech/kcoin/client/contracts/=../../truffle/contracts openzeppelin-solidity/=../../truffle/node_modules/openzeppelin-solidity/ zos-lib/=../../truffle/node_modules/zos-lib/ ../../truffle/contracts/sysvars/SystemVars.sol
//go:generate ../../../build/bin/abigen -abi build/SystemVars.abi -bin build/SystemVars.bin -srcmap build/combined.json -pkg files -type SystemVars -out ./gen_systemvars.go
//go:generate go-bindata -nometadata -o bind_contracts.go -pkg files SystemVars.sol
