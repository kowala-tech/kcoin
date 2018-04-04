package main

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
	"github.com/spf13/viper"
)

var (
	FileConfig string
)

func main() {
	cmd := createCommand()

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func createCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "genesis",
		Short: "Generator of a genesis file.",
		Long:  `Generate a genesis.json file based on a config file or parameters.`,
		Run: func(cmd *cobra.Command, args []string) {
			if FileConfig == "" {
				fmt.Println("Params usage.")
			} else {
				viper.SetConfigFile(FileConfig)
			}
		},
	}

	cmd.Flags().StringVarP(&FileConfig, "config", "c", "", "Use to load configuration from config file.")

	return cmd
}
