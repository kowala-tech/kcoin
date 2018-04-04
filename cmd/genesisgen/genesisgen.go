package main

import (
	"fmt"
	"github.com/prometheus/common/log"
	"gopkg.in/urfave/cli.v1"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "genesisgen"
	app.Usage = "Generates a genesis file."

	app.Action = func(c *cli.Context) {
		fmt.Println("Hola!")
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
