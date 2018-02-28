package light

import (
	"github.com/kowala-tech/kUSD/core"
	"github.com/kowala-tech/kUSD/event"
	"github.com/kowala-tech/kUSD/kusd"
	"github.com/kowala-tech/kUSD/kusddb"
	"github.com/kowala-tech/kUSD/log"
	"github.com/kowala-tech/kUSD/node"
	"github.com/kowala-tech/kUSD/params"
)

// LightKowala represents a the kowala light client service
type LightKowala struct {
	config *kusd.Config

	chainConfig *params.ChainConfig
	chainDB     kusddb.Database // Block chain database

	eventMux *event.TypeMux
}

func New(ctx *node.ServiceContext, config *kusd.Config) (*LightKowala, error) {
	chainDB, err := kusd.CreateDB(ctx, config, "lightchaindata")
	if err != nil {
		return nil, err
	}
	chainConfig, genesisHash, genesisErr := core.SetupGenesisBlock(chainDB, config.Genesis)
	if _, isCompat := genesisErr.(*params.ConfigCompatError); genesisErr != nil && !isCompat {
		return nil, genesisErr
	}
	log.Info("Initialised chain configuration", "config", chainConfig)

	light := &LightKowala{
		config:      config,
		chainConfig: chainConfig,
		chainDB:     chainDb,
		eventMux:    ctx.EventMux,
	}

}
