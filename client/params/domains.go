package params

import "fmt"

const (
	MultiSigDomain = iota
	OracleMgrDomain
	ValidatorMgrDomain
	MiningTokenDomain
	SystemVarsDomain
)

const KowalaTLD = "kowala"

type KNSDomain struct {
	node string
	tld  string
}

func (k KNSDomain) FullDomain() string {
	return fmt.Sprintf("%s.%s", k.node, k.tld)
}

func (k KNSDomain) Tld() string {
	return k.tld
}

func (k KNSDomain) Node() string {
	return k.node
}

var KNSDomains = map[int]KNSDomain{
	MultiSigDomain: {
		node: "multisig",
		tld:  KowalaTLD,
	},
	OracleMgrDomain: {
		node: "oraclemgr",
		tld:  KowalaTLD,
	},
	ValidatorMgrDomain: {
		node: "validatormgr",
		tld:  KowalaTLD,
	},
	MiningTokenDomain: {
		node: "miningtoken",
		tld:  KowalaTLD,
	},
	SystemVarsDomain: {
		node: "systemvars",
		tld:  KowalaTLD,
	},
}
