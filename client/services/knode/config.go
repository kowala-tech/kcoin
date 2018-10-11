package knode

import (
	"time"

	"github.com/kowala-tech/kcoin/client/core"
	"github.com/kowala-tech/kcoin/client/knode/currency"
	"github.com/kowala-tech/kcoin/client/knode/downloader"
	"github.com/kowala-tech/kcoin/client/params"
)

// DefaultConfig contains default settings for use on the Kowala main net.
var DefaultConfig = Config{
	SyncMode:      downloader.FastSync,
	NetworkId:     params.MainnetChainConfig.ChainID.Uint64(),
	LightPeers:    20,
	DatabaseCache: 128,
	TrieCache:     256,
	TrieTimeout:   60 * time.Minute,

	TxPool:   core.DefaultTxPoolConfig,
	Currency: currency.KUSD,
}

//go:generate gencodec -type Config -formats toml -out gen_config.go

type Config struct {
	// The genesis block, which is inserted if the database is empty.
	// If nil, the Ethereum main net block is used.
	Genesis *core.Genesis `toml:",omitempty"`

	// Protocol options
	NetworkId uint64 // Network ID to use for selecting peers to connect to
	SyncMode  downloader.SyncMode
	NoPruning bool

	// Light client options
	LightServ  int `toml:",omitempty"` // Maximum percentage of time allowed for serving LES requests
	LightPeers int `toml:",omitempty"` // Maximum number of LES client peers

	// Database options
	SkipBcVersionCheck bool `toml:"-"`
	DatabaseHandles    int  `toml:"-"`
	DatabaseCache      int
	TrieCache          int
	TrieTimeout        time.Duration

	// Transaction pool options
	TxPool core.TxPoolConfig

	// Enables tracking of SHA3 preimages in the VM
	EnablePreimageRecording bool

	// Miscellaneous options
	DocRoot string `toml:"-"`

	Currency string
}
