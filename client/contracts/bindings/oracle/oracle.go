package oracle

import (
	"math/big"

	"github.com/kowala-tech/kcoin/client/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/client/common/kns"
	"github.com/kowala-tech/kcoin/client/contracts/bindings"
	"github.com/kowala-tech/kcoin/client/log"
	"github.com/kowala-tech/kcoin/client/params"
)

//go:generate solc --allow-paths ., --abi --bin --overwrite --libraries NameHash:0x3b058a1a62E59D185618f64BeBBAF3C52bf099E0 -o build github.com/kowala-tech/kcoin/client/contracts/=../../truffle/contracts openzeppelin-solidity/=../../truffle/node_modules/openzeppelin-solidity/ zos-lib/=../../truffle/node_modules/zos-lib/ ../../truffle/contracts/oracle/OracleMgr.sol
//go:generate solc --allow-paths ., --combined-json bin-runtime,srcmap-runtime --overwrite --libraries NameHash:0x3b058a1a62E59D185618f64BeBBAF3C52bf099E0 -o build/oracle-combined github.com/kowala-tech/kcoin/client/contracts/=../../truffle/contracts openzeppelin-solidity/=../../truffle/node_modules/openzeppelin-solidity/ zos-lib/=../../truffle/node_modules/zos-lib/ ../../truffle/contracts/oracle/OracleMgr.sol
//go:generate ../../../build/bin/abigen -abi build/OracleMgr.abi -bin build/OracleMgr.bin -srcmap build/oracle-combined/combined.json -pkg oracle -type OracleMgr -out ./gen_oracle.go
//go:generate solc --allow-paths ., --abi --bin --overwrite -o build github.com/kowala-tech/kcoin/client/contracts/=../../truffle/contracts openzeppelin-solidity/=../../truffle/node_modules/openzeppelin-solidity/ zos-lib/=../../truffle/node_modules/zos-lib/ ../../truffle/contracts/oracle/ExchangeMgr.sol
//go:generate solc --allow-paths ., --combined-json bin-runtime,srcmap-runtime --overwrite -o build/exchange-combined github.com/kowala-tech/kcoin/client/contracts/=../../truffle/contracts openzeppelin-solidity/=../../truffle/node_modules/openzeppelin-solidity/ zos-lib/=../../truffle/node_modules/zos-lib/ ../../truffle/contracts/oracle/ExchangeMgr.sol
//go:generate ../../../build/bin/abigen -abi build/ExchangeMgr.abi -bin build/ExchangeMgr.bin -srcmap build/exchange-combined/combined.json -pkg oracle -type ExchangeMgr -out ./gen_exchange.go

type Manager struct {
	*OracleMgrSession
}

// @TODO(rgeraldes) - temporary method
func (mgr *Manager) Domain() string {
	return ""
}

// Bind returns a binding to the current oracle mgr
func Bind(contractBackend bind.ContractBackend, chainID *big.Int) (bindings.Binding, error) {
	addr, err := kns.GetAddressFromDomain(
		params.KNSDomains[params.OracleMgrDomain].FullDomain(),
		contractBackend,
	)
	if err != nil {
		log.Error("can't find Oracle for given Network", "chainID", chainID.String())
		return nil, bindings.ErrNoAddress
	}

	mgr, err := NewOracleMgr(addr, contractBackend)
	if err != nil {
		return nil, err
	}

	return &Manager{
		&OracleMgrSession{
			Contract: mgr,
			CallOpts: bind.CallOpts{},
		},
	}, nil
}
