package ownership

import (
	"github.com/kowala-tech/kcoin/client/core/types"
)

//go:generate solc --allow-paths . --abi --bin --overwrite -o build ../../truffle/contracts/ownership/MultiSigWallet.sol
//go:generate ../../../build/bin/abigen -abi build/MultiSigWallet.abi -bin build/MultiSigWallet.bin -pkg ownership -type MultiSigWallet -out ./gen_multisig.go

// Proxy function for the unpack log in the contract
func (w *MultiSigWalletFilterer) UnpackLog(out interface{}, event string, log types.Log) error {
	return w.contract.UnpackLog(out, event, log)
}
