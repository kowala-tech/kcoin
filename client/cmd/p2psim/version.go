package main

import "github.com/kowala-tech/kcoin/params"

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
