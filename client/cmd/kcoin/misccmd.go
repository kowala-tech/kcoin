package main

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/kowala-tech/kcoin/client/cmd/utils"
	"github.com/kowala-tech/kcoin/client/knode/protocol"
	"github.com/kowala-tech/kcoin/client/params"
	"gopkg.in/urfave/cli.v1"
	"github.com/kowala-tech/kcoin/client/version"
)

var (
	versionCommand = cli.Command{
		Action:    utils.MigrateFlags(versionPrint),
		Name:      "version",
		Usage:     "Print version numbers",
		ArgsUsage: " ",
		Category:  "MISCELLANEOUS COMMANDS",
		Flags: []cli.Flag{
			utils.VersionRepository,
		},
		Description: `
The output of this command is supposed to be machine-readable.
`,
	}
	updateCommand = cli.Command{
		Action:    utils.MigrateFlags(latest),
		Name:      "update",
		Usage:     "Update binary to latest version",
		ArgsUsage: " ",
		Category:  "MISCELLANEOUS COMMANDS",
		Flags: []cli.Flag{
			utils.VersionRepository,
		},
		Description: `
This should update binary to latest version.
`,
	}
	licenseCommand = cli.Command{
		Action:    utils.MigrateFlags(license),
		Name:      "license",
		Usage:     "Display license information",
		ArgsUsage: " ",
		Category:  "MISCELLANEOUS COMMANDS",
	}
)

func versionPrint(ctx *cli.Context) error {
	fmt.Println(strings.Title(clientIdentifier))
	fmt.Println("Version:", params.Version)

	printLatestIfAvailable(ctx)

	if params.Commit != "" {
		fmt.Println("Git Commit:", params.Commit)
	}
	if params.BuildTime != "" {
		fmt.Println("BuildTime:", params.BuildTime)
	}
	fmt.Println("Architecture:", runtime.GOARCH)
	fmt.Println("Constants Versions:", protocol.Constants.Versions)
	fmt.Println("Go Version:", runtime.Version())
	fmt.Println("Operating System:", runtime.GOOS)
	return nil
}

func printLatestIfAvailable(ctx *cli.Context) {
	repository := ctx.GlobalString(utils.VersionRepository.Name)
	finder := version.NewFinder(repository)
	latest, err := finder.Latest(runtime.GOOS, runtime.GOARCH)
	if err == nil {
		fmt.Println("Latest Version Available:", latest.Semver().String())
	}
	// ignore error, we don't print latest version available
}

func license(_ *cli.Context) error {
	fmt.Println(`kcoin is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

kcoin is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with knode. If not, see <http://www.gnu.org/licenses/>.
`)
	return nil
}

func latest(ctx *cli.Context) error {
	repository := ctx.GlobalString(utils.VersionRepository.Name)

	updater, err := version.NewUpdater(repository)
	if err != nil {
		return err
	}

	return updater.Update()
}
