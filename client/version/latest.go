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
	assets, err := f.assetsBy(platform(os, arch))
	if err != nil {
		return asset{}, err
	}

	return f.latest(assets)
}

func (f *finder) LatestForMajor(os, arch string, major uint64) (Asset, error) {
	assets, err := f.assetsBy(platformMajor(os, arch, major))
	if err != nil {
		return asset{}, err
	}

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

type assetFilterFunc func(asset Asset) bool

func platform(os, arch string) assetFilterFunc {
	return func(asset Asset) bool {
		return asset.Arch() == arch && asset.Os() == os
	}
}

func platformMajor(os, arch string, major uint64) assetFilterFunc {
	return func(asset Asset) bool {
		return asset.Semver().Major == major &&
			platform(os, arch)(asset)
	}
}

func (f *finder) assetsBy(allowFilter assetFilterFunc) ([]Asset, error) {
	assets, err := f.All()
	if err != nil {
		return nil, err
	}
	for _, asset := range assets {
		if allowFilter(asset) {
			assets = append(assets, asset)
		}
	}
	return assets, nil
}
