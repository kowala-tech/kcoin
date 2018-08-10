package params

import "fmt"

const (
	MultiSigDomain = iota
	OracleMgrDomain
	ValidatorMgrDomain
	MiningTokenDomain
)

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
	MultiSigDomain: KNSDomain{
		node: "multisig",
		tld:  "kowala",
	},
	OracleMgrDomain: KNSDomain{
		node: "oraclemgr",
		tld:  "kowala",
	},
	ValidatorMgrDomain: KNSDomain{
		node: "validatormgr",
		tld:  "kowala",
	},
	MiningTokenDomain: KNSDomain{
		node: "miningtoken",
		tld:  "kowala",
	},
}
