package cmd

import (
	"fmt"
	"os"

	"github.com/imdario/mergo"
	"github.com/kowala-tech/kUSD/cmd/utils"
	"github.com/kowala-tech/kUSD/config"
	"github.com/kowala-tech/kUSD/params"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfg       *config.Config // top level config file
	gitCommit = " "          // Git SHA1 commit hash of the release (set via linker flags)
)

// persistent flags
var (
	cfgFile   string
	bootNodes []string
)

// KowalaCmd represents the base command when called without any subcommands
var KowalaCmd = &cobra.Command{
	Use:     "kusd",
	Short:   "the kowala command line interface",
	Version: params.VersionWithCommit(gitCommit),

	// @NOTE (rgeraldes) - executed for all the commands under KowalaCmd (except version cmd)
	PersistentPreRunE: func(cmd *cobra.Command, args []string) {
		// skip if version command
		if cmd.Name() == VersionCmd.Name() {
			return
		}
		cmd.

			// initialize top level config
			config = loadConfig()
	},
	Args: cobra.NoArgs,
}

func init() {

	initPersistentFlags(KowalaCmd)
}

func initPersistentFlags(cmd *cobra.Command) {
	KowalaCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "/path/to/config.yml|json|toml)")
	KowalaCmd.PersistentFlags().StringArrayVar(&bootNodes, "bootnodes", nil, "Comma separated enode URLs for P2P discovery bootstrap (set v4+v5 instead for light servers)")

	// @TODO(rgeraldes) - handle errors
	// bash completion
	KowalaCmd.PersistentFlags().SetAnnotation("config", cobra.BashCompSubdirsInDir, []string{})
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := KowalaCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// addNodeFlags adds node related persistent flags
func addNodeFlags(cmd *cobra.Command) {

	/*
		KowalaCmd.Flags().AddFlag()
		// fast sync
		KowalaCmd.Flags().AddFlag()
		KowalaCmd.PersistentFlags().Bool("fast", true, "Enable fast syncing through state downloads")
		viper.BindPFlag("fast", KUSDCmd.PersistentFlags().Lookup("fast"))

		// network id
		KowalaCmd.PersistentFlags().Int("networkid", false, "Network identifier (integer, 1=Main, 10000=Test)")
		viper.BindPFlag("networkid", KUSDCmd.PersistentFlags().Lookup("networkid"))

		// node key file
		KowalaCmd.PersistentFlags().Int("nodekey", 0, "P2P node key file")

		/*
		// keystore dir
		KUSDCmd.PersistentFlags().String("keystore", "", "Directory for the keystore (default = inside the datadir)")
		viper.BindPFlag("keystore", KUSDCmd.PersistentFlags().Lookup("keystore"))
		KUSDCmd.Flags().SetAnnotation("keystore", cobra.BashCompSubdirsInDir, []string{})) // bash completion

		// light kdf
		KUSDCmd.PersistentFlags().Bool("lightkdf",false, "Reduce key-derivation RAM & CPU usage at some expense of KDF strength")
		viper.BindPFlag("lightkfg", KUSDCmd.PersistentFlags().Lookup("lightkdf"))

		// no usb
		KUSDCmd.PersistentFlags().Bool("nousb", false, "Disables monitoring for and managine USB hardware wallets")
		viper.BindPFlag("nousb", KUSDCmd.PersistentFlags().Lookup("nousb"))
	*/
}

// initConfig returns the top level configuration
func initConfig() *config.Config {
	// default config is the foundation
	finalCfg := config.DefaultConfig()

	if cfgFile != "" {
		// Use config file from the persistent flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".kusd" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".kusd")
	}

	// read in environment variables that match
	//viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	// unmarshall current configuration in viper
	var viperCfg cfg.Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	// Apply extra flags logic
	// @TODO(rgeraldes) -

	// load viper configuration on top of the defaults
	// @NOTE (rgeraldes) Unmarshal does not replace the values that already
	// exist, so we cannot Unmarshal on top of a default config struct
	// https://github.com/spf13/viper/issues/295
	if err := mergo.Merge(finalCfg, viperCfg); err != nil {
		utils.Fatalf("Failed merging configuration: %v", err)
	}

	return finalCfg
}
