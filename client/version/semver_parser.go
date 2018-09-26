package version

import (
	"github.com/blang/semver"
	"strings"
)

// MakeSemver returns a new semver version instance personalised to ignore Pre tag
func MakeSemver(s string) (semver.Version, error) {
	if preIndex := strings.IndexRune(s, '-'); preIndex != -1 {
		s = s[:preIndex]
	}
	return semver.Make(s)
}
