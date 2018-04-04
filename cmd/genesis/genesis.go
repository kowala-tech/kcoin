package main

import (
	"github.com/spf13/cobra"
	"fmt"
)

func main() {
	 cmd := &cobra.Command{
		Use:   "hugo",
		Short: "Hugo is a very fast static site generator",
		Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Holaj")
		},
	}

	cmd.Execute()
}
