package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/core"
	"github.com/kowala-tech/kUSD/log"
	"github.com/kowala-tech/kUSD/node"
	"github.com/spf13/cobra"
)

// InitCmd bootstraps and initializes a new genesis block
var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Bootstrap and initialize a new genesis block",
	Long: `
	The init command initializes a new genesis block and definition for the network.
	This is a destructive action and changes the network in which you will be
	participating.
	It expects the genesis file as argument.`,
	RunE: init,
	Args: cobra.ExactArgs(1),
}

// initGenesis will initialise the given JSON format genesis file and writes it as
// the zero'd block (i.e. genesis) or will fail hard if it can't succeed.
func init(cmd *cobra.Command, args []string) error {
	// Make sure we have a valid genesis JSON
	file, err := os.Open(args[0])
	if err != nil {
		return fmt.Errorf("Failed to read genesis file: %v", err)
	}
	defer file.Close()
	genesis := new(core.Genesis)
	if err := json.NewDecoder(file).Decode(genesis); err != nil {
		return fmt.Errorf("invalid genesis file: %v", err)
	}

	// Open and initialise full database
	node := node.New(cfg.Node)
	for _, name := range []string{"chaindata"} {
		chaindb, err := node.OpenDatabase(name, 0, 0)
		if err != nil {
			return fmt.Errof("Failed to open database: %v", err)
		}
		_, hash, err := core.SetupGenesisBlock(chaindb, genesis)
		if err != nil {
			return fmt.Errorf("Failed to write genesis block: %v", err)
		}
		log.Info("Successfully wrote genesis state", "database", name, "hash", hash)
	}
	return nil
}
