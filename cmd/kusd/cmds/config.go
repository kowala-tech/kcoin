package cmd

import (
	"path/filepath"

	"github.com/kowala-tech/kUSD/cmd/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	defaultConfigFile = "kusd.toml"
)

// flags
var (
	out string
)

// ConfigCmd contains the config related commands
var ConfigCmd = &cobra.Command{
	Use:  "config",
	Args: cobra.NoArgs,
	RunE: config,
}

func init() {
	ConfigCmd.AddCommand(generateConfigCmd)
	//ConfigCmd.AddCommand(listConfigCmd)

	// flags
	ConfigCmd.Flags().StringVar(&out, "out", filepath.Join(utils.HomeDir(), defaultConfigFile), "path/to/file")

	// bash completion
	ConfigCmd.Flags().SetAnnotation("out", cobra.BashCompSubdirsInDir, []string{})
}

// generateConfig
func config(cmd *cobra.Command, args []string) error {
	var dir string

	if out != "" {
		return viper.WriteConfigAs(out)
	}

	viper.Debug()
	return nil
}
