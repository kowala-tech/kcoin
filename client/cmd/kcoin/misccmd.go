package main

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/kowala-tech/kcoin/client/cmd/utils"
	"github.com/kowala-tech/kcoin/client/knode/protocol"
	"github.com/kowala-tech/kcoin/client/params"
	"gopkg.in/urfave/cli.v1"
	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
)

var (
	versionCommand = cli.Command{
		Action:    utils.MigrateFlags(version),
		Name:      "version",
		Usage:     "Print version numbers",
		ArgsUsage: " ",
		Category:  "MISCELLANEOUS COMMANDS",
		Description: `
The output of this command is supposed to be machine-readable.
`}

	updateCommand = cli.Command{
		Action:    utils.MigrateFlags(update),
		Name:      "update",
		Usage:     "Update client to latest version",
		ArgsUsage: " ",
		Category:  "MISCELLANEOUS COMMANDS",
		Description: `
This should update client to latest version.
`}

	licenseCommand = cli.Command{
		Action:    utils.MigrateFlags(license),
		Name:      "license",
		Usage:     "Display license information",
		ArgsUsage: " ",
		Category:  "MISCELLANEOUS COMMANDS",
	}
)

func version(ctx *cli.Context) error {
	fmt.Println(strings.Title(clientIdentifier))
	fmt.Println("Version:", params.Version)
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

func update(ctx *cli.Context) error {
	v := semver.MustParse(params.Version)
	latest, err := selfupdate.UpdateSelf(v, "kowala-tech/kcoin")
	if err != nil {
		fmt.Println("Binary update failed:", err)
		return err
	}
	if latest.Version.Equals(v) {
		// latest version is the same as current version. It means current binary is up to date.
		fmt.Println("Current binary is the latest version", params.Version)
	} else {
		fmt.Println("Successfully updated to version", latest.Version)
		fmt.Println("Release note:\n", latest.ReleaseNotes)
	}
	return nil
}
