package main

import (
	"github.com/kowala-tech/kcoin/client/params"
	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
	"fmt"
)

var (
	// Git SHA1 commit hash of the release (set via linker flags)
	gitCommit = ""
	// Git tag of the release (set via linker flags)
	gitTag = ""
	// Build time in nonoseconds of the release (set via linker flags)
	buildTime = ""
)

func init() {
	params.SetGitTagVersion(gitTag)
	params.SetBuildTime(buildTime)
	params.SetCommit(gitCommit)
}

func doSelfUpdate() {
	v := semver.MustParse(gitTag)
	latest, err := selfupdate.UpdateSelf(v, "kowala-tech/kcoin")
	if err != nil {
		fmt.Println("Binary update failed:", err)
		return
	}
	if latest.Version.Equals(v) {
		// latest version is the same as current version. It means current binary is up to date.
		fmt.Println("Current binary is the latest version", gitTag)
	} else {
		fmt.Println("Successfully updated to version", latest.Version)
		fmt.Println("Release note:\n", latest.ReleaseNotes)
	}
}
