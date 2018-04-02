package main

import (
	"gopkg.in/urfave/cli.v1"
	"fmt"
	"os"
	"github.com/prometheus/common/log"
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
