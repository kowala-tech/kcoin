// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/kowala-tech/kcoin/mock-exchange/server"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Launches a server as a mockup service.",
	Long: `For now it only fakes data comming from Exrates, in the future
	it will support other services.`,
	Run: func(cmd *cobra.Command, args []string) {
		s, err := server.New(server.DefaultConfig())
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
