package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Kowala network.
var MainnetBootnodes = []string{}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// test network.
var TestnetBootnodes = []string{}

var DevnetBootnodes = []string{}

var DevnetDiscoveryV5Bootnodes = []string{
	"enode://dd38c33eff2ba2fbf152bc698d86fa5baa18b30973e45700c48cdcc8555f2d437160731138960bc46f42b26e363ee5f8f1daa592cafa852669f91ef201ea569d@35.178.226.105:32233",
}

// TestnetDiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var TestnetDiscoveryV5Bootnodes = []string{
	"enode://1ed36c00c77e7a2d824699baadaab9142818b8f1904d397c65f3d75793b47e99a251704c2ad66d874292516e8c6aac134184e8cc3897899ea776c5ed3ba75ed8@52.23.43.192:32233",
	"enode://ae54a310873b252daaa9203c1995dc9e89f8afb2c70b255c44d005cbb79714a553b101b913eff128b197b6e0e24ea4991ea9528fca94621f4438bf4aa211a08e@54.175.98.140:32233",
}

var MainnetDiscoveryV5Bootnodes = []string{
	"enode://49bb43d743fef3f8b96c522316504c78a126dde087db975986202e098ee43bdaf550ca3df3a1512210c94b2a33c2821a7afb68fad025259ef24910005f3c0a50@18.130.99.150:32233",
	"enode://beb55d4909fc62b99505593e421aef2cfc596ea41757f4ee1524c418399e29876ef38f89e70b0308c3609839f1b4a28d4c375aba4573cf7c1885146328405760@18.136.143.195:32233",
	"enode://a8745bd94d63f85d2ee38bca0d6bf7672688c5b7cda5060ca7dd5decbbf179e7ad312e53e71f0f5b980037f5b6e6e8ea4a521b9b3cb1844f0654bba519eb8dbd@54.176.194.8:32233",
}
