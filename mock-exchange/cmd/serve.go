package cmd

import (
	"fmt"
	"os"

	"github.com/kowala-tech/kcoin/mock-exchange/server"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Launches a server as a mockup service.",
	Long: `For now it only fakes data comming from Exrates, in the future
	it will support other services.`,
	Run: func(cmd *cobra.Command, args []string) {
		s, err := server.New(server.DefaultConfig(), server.GetRouter())
		if err != nil {
			fmt.Printf("Error creating s: %s", err)
			os.Exit(1)
		}

		err = s.Start()
		if err != nil {
			fmt.Printf("Error starting s: %s", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
