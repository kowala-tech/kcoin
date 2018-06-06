package knode

import (
	"math/big"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"time"

	"github.com/kowala-tech/kcoin/common"
	"github.com/kowala-tech/kcoin/common/hexutil"
	"github.com/kowala-tech/kcoin/core"
	"github.com/kowala-tech/kcoin/knode/downloader"
	"github.com/kowala-tech/kcoin/knode/gasprice"
	"github.com/kowala-tech/kcoin/params"
)

// DefaultConfig contains default settings for use on the Kowala main net.
var DefaultConfig = Config{
	SyncMode:      downloader.FastSync,
	NetworkId:     params.MainnetChainConfig.ChainID.Uint64(),
	LightPeers:    100,
	DatabaseCache: 768,
	TrieCache:     256,
	TrieTimeout:   5 * time.Minute,
	GasPrice:      big.NewInt(1),

	TxPool: core.DefaultTxPoolConfig,
	GPO: gasprice.Config{
		Blocks:     20,
		Percentile: 60,
	},
}

//go:generate gencodec -type Config -field-override configMarshaling -formats toml -out gen_config.go

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

	// consensus validation-related options
	Coinbase  common.Address `toml:",omitempty"`
	Deposit   *big.Int       `toml:",omitempty"`
	ExtraData []byte         `toml:",omitempty"`
	GasPrice  *big.Int

	// Transaction pool options
	TxPool core.TxPoolConfig

	// Gas Price Oracle options
	GPO gasprice.Config

	// Enables tracking of SHA3 preimages in the VM
	EnablePreimageRecording bool

	// Miscellaneous options
	DocRoot string `toml:"-"`
}

type configMarshaling struct {
	ExtraData hexutil.Bytes
}
