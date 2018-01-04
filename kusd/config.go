package kusd

import (
	"math/big"

	"github.com/kowala-tech/kUSD/common"
	"github.com/kowala-tech/kUSD/core"
	"github.com/kowala-tech/kUSD/kusd/downloader"
	"github.com/kowala-tech/kUSD/kusd/gasprice"
	"github.com/kowala-tech/kUSD/params"
)

// DefaultConfig contains default settings for use on the Kowala main net.
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

type Config struct {
	// The genesis block, which is inserted if the database is empty.
	// If nil, the Kowala main net block is used.
	Genesis *core.Genesis `mapstructure:"genesis"`

	// Protocol options
	NetworkID uint64              `mapstructure:"networkid"` // Network ID to use for selecting peers to connect to
	SyncMode  downloader.SyncMode `mapstructure:"syncmode"`

	// Database options
	DatabaseHandles int `mapstructure:"dbhandles"`
	DatabaseCache   int `mapstructure:"dbcache"`

	// Validator-related options
	Coinbase common.Address `mapstructure:"coinbase"`
	GasPrice *big.Int       `mapstructure:"gasprice"`

	// Transaction pool options
	TxPool core.TxPoolConfig `mapstructure:"txpool"`

	// Gas Price Oracle options
	GPO gasprice.Config `mapstructure:"gpo"`

	// Enables tracking of SHA3 preimages in the VM
	EnablePreimageRecording bool `mapstructure:"preimage"`

	// Miscellaneous options
	DocRoot string `mapstructure:"docroot"`
}
