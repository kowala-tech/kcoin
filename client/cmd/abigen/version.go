package main

import "github.com/kowala-tech/kcoin/client/params"

var (
	// Git SHA1 commit hash of the release (set via linker flags)
	gitCommit = "0"
	// Git tag of the release (set via linker flags)
	gitTag = "0.0.0"
	// Build time in nonoseconds of the release (set via linker flags)
	buildTime = "0"
)

func init() {
	params.SetGitTagVersion(gitTag)
	params.SetBuildTime(buildTime)
	params.SetCommit(gitCommit)
}
