package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Kowala network.
var MainnetBootnodes = []string{
	"enode://49bb43d743fef3f8b96c522316504c78a126dde087db975986202e098ee43bdaf550ca3df3a1512210c94b2a33c2821a7afb68fad025259ef24910005f3c0a50@54.197.186.25:32233",
	"enode://beb55d4909fc62b99505593e421aef2cfc596ea41757f4ee1524c418399e29876ef38f89e70b0308c3609839f1b4a28d4c375aba4573cf7c1885146328405760@52.8.177.228:32233",
	"enode://a8745bd94d63f85d2ee38bca0d6bf7672688c5b7cda5060ca7dd5decbbf179e7ad312e53e71f0f5b980037f5b6e6e8ea4a521b9b3cb1844f0654bba519eb8dbd@35.176.220.151:32233",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// test network.
var TestnetBootnodes = []string{
	"enode://a46e5885d3da52bf452fcea7269b91059d93819e7e906cb5d29575508a18a4009e5d688bf29abf8cdd55aad65a9ca4beab98de8c0a226d694c8140acdedf3a55@18.219.208.190:33447",
	"enode://9791d40d97585d7241edeedd2491187e82e3273414cf1bed2aca049e8b21d96930437efe3e591d6996ee303d3e2fe4338e53551bb6e50f05dfeeaccd660eac01@18.220.38.32:33447",
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{}
