package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"unicode"

	"gopkg.in/urfave/cli.v1"

	"path/filepath"

	"github.com/kowala-tech/kcoin/client/cmd/utils"
	"github.com/kowala-tech/kcoin/client/knode"
	"github.com/kowala-tech/kcoin/client/node"
	"github.com/kowala-tech/kcoin/client/params"
	"github.com/kowala-tech/kcoin/client/stats"
	"github.com/naoina/toml"
)

var (
	dumpConfigCommand = cli.Command{
		Action:      utils.MigrateFlags(dumpConfig),
		Name:        "dumpconfig",
		Usage:       "Show configuration values",
		ArgsUsage:   "",
		Flags:       append(append(nodeFlags, rpcFlags...)),
		Category:    "MISCELLANEOUS COMMANDS",
		Description: `The dumpconfig command shows configuration values.`,
	}

	configFileFlag = cli.StringFlag{
		Name:  "config",
		Usage: "TOML configuration file",
	}
)

// These settings ensure that TOML keys use the same names as Go struct fields.
var tomlSettings = toml.Config{
	NormFieldName: func(rt reflect.Type, key string) string {
		return key
	},
	FieldToKey: func(rt reflect.Type, field string) string {
		return field
	},
	MissingField: func(rt reflect.Type, field string) error {
		link := ""
		if unicode.IsUpper(rune(rt.Name()[0])) && rt.PkgPath() != "main" {
			link = fmt.Sprintf(", see https://godoc.org/%s#%s for available fields", rt.PkgPath(), rt.Name())
		}
		return fmt.Errorf("field '%s' is not defined in %s%s", field, rt.String(), link)
	},
}

type kcoinConfig struct {
	Kowala knode.Config
	Node   node.Config
	Stats  stats.Config
}

func loadConfig(file string, cfg *kcoinConfig) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	err = tomlSettings.NewDecoder(bufio.NewReader(f)).Decode(cfg)
	// Add file name to errors that have a line number.
	if _, ok := err.(*toml.LineError); ok {
		err = errors.New(file + ", " + err.Error())
	}
	return err
}

func defaultNodeConfig() node.Config {
	cfg := node.DefaultConfig
	cfg.Name = clientIdentifier
	cfg.Version = params.VersionWithCommit
	cfg.HTTPModules = append(cfg.HTTPModules, "eth", "shh")
	cfg.WSModules = append(cfg.WSModules, "eth", "shh")
	cfg.IPCPath = fmt.Sprintf("%s.ipc", clientIdentifier)
	return cfg
}

func makeConfigNode(ctx *cli.Context) (*node.Node, kcoinConfig) {
	// Load defaults.
	cfg := kcoinConfig{
		Kowala: knode.DefaultConfig,
		Node:   defaultNodeConfig(),
	}

	// Load config file.
	if file := ctx.GlobalString(configFileFlag.Name); file != "" {
		if err := loadConfig(file, &cfg); err != nil {
			utils.Fatalf("%v", err)
		}
	}

	// Apply flags.
	utils.SetNodeConfig(ctx, &cfg.Node)
	cfg.Node.DataDir = getPathBasedOnCurrency(cfg)
	stack, err := node.New(&cfg.Node)
	if err != nil {
		utils.Fatalf("Failed to create the protocol stack: %v", err)
	}
	utils.SetKowalaConfig(ctx, stack, &cfg.Kowala)
	if ctx.GlobalIsSet(utils.KowalaStatsURLFlag.Name) {
		cfg.Stats.URL = ctx.GlobalString(utils.KowalaStatsURLFlag.Name)
	}

	return stack, cfg
}

func getPathBasedOnCurrency(cfg kcoinConfig) string {
	return filepath.Join(cfg.Node.DataDir, cfg.Kowala.Currency)
}

func makeFullNode(ctx *cli.Context) *node.Node {
	stack, cfg := makeConfigNode(ctx)

	utils.RegisterKowalaService(stack, &cfg.Kowala)

	// Add the Stats daemon if requested.
	statsURL := cfg.Stats.GetURL()
	if statsURL != "" {
		utils.RegisterKowalaStatsService(stack, statsURL)
	}

	return stack
}

// dumpConfig is the dumpconfig command.
func dumpConfig(ctx *cli.Context) error {
	_, cfg := makeConfigNode(ctx)
	comment := ""

	if cfg.Kowala.Genesis != nil {
		cfg.Kowala.Genesis = nil
		comment += "# Note: this config doesn't contain the genesis block.\n\n"
	}

	out, err := tomlSettings.Marshal(&cfg)
	if err != nil {
		return err
	}
	io.WriteString(os.Stdout, comment)
	os.Stdout.Write(out)
	return nil
}
