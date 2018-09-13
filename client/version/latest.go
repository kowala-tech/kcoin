package version

import (
	"bufio"
	"github.com/kowala-tech/kcoin/client/log"
	"github.com/pkg/errors"
	"net/http"
)

const indexFile = "index.txt"

type Finder interface {
	All() ([]Asset, error)
	Latest(os, arch string) (Asset, error)
	LatestForMajor(os, arch string, major string) (Asset, error)
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
		version, err := filenameParser(scanner.Text())
		if err != nil {
			// ignore error and continue to next filename
			log.Debug("could not parse filename", err)
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
	allAssets, err := f.All()
	if err != nil {
		return asset{}, err
	}

	assets := NewAssetFilterer(allAssets).by(platform(os, arch))

	return f.latest(assets)
}

func (f *finder) LatestForMajor(os, arch string, major uint64) (Asset, error) {
	allAssets, err := f.All()
	if err != nil {
		return asset{}, err
	}

	assets := NewAssetFilterer(allAssets).by(platformMajor(os, arch, major))

	return f.latest(assets)
}

func (f *finder) latest(assets []Asset) (Asset, error) {
	if len(assets) == 0 {
		return asset{}, errors.New("no version found")
	}

	latest := assets[0]
	for _, asset := range assets {
		if asset.Semver().GT(latest.Semver()) {
			latest = asset
		}
	}

	if latest == nil {
		return asset{}, errors.New("no version found")
	}

	return latest, nil
}
