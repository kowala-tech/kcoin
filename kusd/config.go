package kusd

import (
	"math/big"
	"os"
	"os/user"

	"github.com/kowala-tech/kUSD/common"
	"github.com/kowala-tech/kUSD/core"
	"github.com/kowala-tech/kUSD/eth/downloader"
	"github.com/kowala-tech/kUSD/eth/gasprice"
	"github.com/kowala-tech/kUSD/params"
)

// DefaultConfig contains default settings for use on the KUSD main net.
var DefaultConfig = Config{
	SyncMode:      downloader.FastSync,
	NetworkID:     1,
	DatabaseCache: 128,
	GasPrice:      big.NewInt(18 * params.Shannon),

	TxPool: core.DefaultTxPoolConfig,
	GPO: gasprice.Config{
		Blocks:     10,
		Percentile: 50,
	},
}

func init() {
	home := os.Getenv("HOME")
	if home == "" {
		if user, err := user.Current(); err == nil {
			home = user.HomeDir
		}
	}
}

//go:generate gencodec -type Config -formats toml -out gen_config.go

type Config struct {
	// The genesis block, which is inserted if the database is empty.
	// If nil, the Ethereum main net block is used.
	Genesis *core.Genesis `toml:",omitempty"`

	// Protocol options
	NetworkID uint64 // Network ID to use for selecting peers to connect to
	SyncMode  downloader.SyncMode

	MaxPeers int `toml:"-"` // Maximum number of global peers

	// Database option
	//@TODO(rgeraldes) - analyze in the future
	//SkipBcVersionCheck bool `toml:"-"`
	DatabaseHandles int `toml:"-"`
	DatabaseCache   int

	// Validator-related options
	Coinbase common.Address `toml:",omitempty"`
	GasPrice *big.Int

	// Transaction pool options
	TxPool core.TxPoolConfig

	// Gas Price Oracle options
	GPO gasprice.Config

	// Enables tracking of SHA3 preimages in the VM
	EnablePreimageRecording bool

	// Miscellaneous options
	DocRoot string `toml:"-"`
}

// @NOTE(rgeraldes) - removed the gencodec overrides struct
