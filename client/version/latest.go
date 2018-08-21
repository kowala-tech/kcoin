package version

import (
	"bufio"
	"net/http"

	"github.com/kowala-tech/kcoin/client/log"
	"github.com/pkg/errors"
)

const indexFile = "index.txt"

type Finder interface {
	All() ([]Asset, error)
	Latest(os, arch string) (Asset, error)
}

func NewFinder(repository string) *finder {
	return &finder{
		repository: repository,
	}
}

type finder struct {
	repository string
}

func (f *finder) All() ([]Asset, error) {
	response, err := http.Get(f.repository + "/" + indexFile)
	if err != nil {
		return []Asset{}, err
	}
	defer response.Body.Close()

	var assets []Asset
	scanner := bufio.NewScanner(response.Body)
	for scanner.Scan() {
		text := scanner.Text()
		version, err := filenameParser(text)
		if err != nil {
			// ignore error and continue to next filename
			log.Debug("could not parse filename", "err", err, "test", text)
			continue
		}
		assets = append(assets, version)
	}

	if err := scanner.Err(); err != nil {
		return []Asset{}, err
	}

	return assets, nil
}

func (f *finder) Latest(os, arch string) (Asset, error) {
	assets, err := f.All()
	if err != nil {
		return asset{}, err
	}
	var latest Asset
	for _, asset := range assets {
		// is right platform
		if asset.Arch() != arch || asset.Os() != os {
			continue
		}

		// first version that matches this platform
		if latest == nil {
			latest = asset
			continue
		}

		// is this new greater then the one we have
		if asset.Semver().GT(latest.Semver()) {
			latest = asset
		}
	}

	if latest == nil {
		return asset{}, errors.New("no version found")
	}

	return latest, nil
}
