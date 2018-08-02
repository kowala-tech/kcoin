package version

import (
	"errors"
	"github.com/blang/semver"
	"regexp"
)

// matches version (1.0.0), os (linux) and architecture (amd64)
// ex: kcoin-1.0.11-linux-amd64.zip -> version: 1.0.11, platform: linux-amd64
var re = regexp.MustCompile("^kcoin-(\\d+\\.\\d+\\.\\d+)-([\\w\\-\\d\\.]+)-(\\w+)(\\.exe)*\\.(zip|gz)$")

func filenameParser(filename string) (Asset, error) {
	if len(filename) < len("kcoin-1.0.1-os.zip") {
		return asset{}, errors.New("cant parse filename")
	}

	matches := re.FindAllStringSubmatch(filename, -1)
	if len(matches) < 1 {
		return asset{}, errors.New("cant parse filename")
	}
	if len(matches[0]) < 4 {
		return asset{}, errors.New("cant parse filename")
	}

	version, err := semver.Make(matches[0][1])
	if err != nil {
		return asset{}, err
	}

	return NewAsset(version, matches[0][2], matches[0][3], filename), nil
}
