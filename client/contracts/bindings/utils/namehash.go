package utils

//go:generate solc --allow-paths ., --abi --bin --overwrite -o build ../../truffle/contracts/utils/Strings.sol
//go:generate solc --allow-paths ., --combined-json bin-runtime,srcmap-runtime --overwrite -o build/strings-combined ../../truffle/contracts/utils/Strings.sol
//go:generate ../../../build/bin/abigen -abi build/strings.abi -bin build/strings.bin -srcmap build/strings-combined/combined.json -pkg utils -type strings -out ./gen_strings.go
//go:generate solc --allow-paths ., --abi --bin --libraries Strings:0x2d7465b88a0A5A1bBff2671C8ED78F7506465ddc --overwrite -o build ../../truffle/contracts/utils/NameHash.sol
//go:generate solc --allow-paths ., --combined-json bin-runtime,srcmap-runtime --libraries Strings:0x2d7465b88a0A5A1bBff2671C8ED78F7506465ddc --overwrite -o build/namehash-combined ../../truffle/contracts/utils/NameHash.sol
//go:generate ../../../build/bin/abigen -abi build/NameHash.abi -bin build/NameHash.bin -srcmap build/namehash-combined/combined.json -pkg utils -type NameHash -out ./gen_namehash.go
//go:generate go-bindata -o bind_contracts.go -pkg utils ../../truffle/contracts/utils/NameHash.sol ../../truffle/contracts/utils/Strings.sol
