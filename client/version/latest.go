package version

import (
	"github.com/pkg/errors"
)

type Finder interface {
	All() ([]Asset, error)
	Latest(os, arch string) (Asset, error)
	LatestForMajor(os, arch string, major string) (Asset, error)
}

func NewFinder(repository AssetRepository) *finder {
	return &finder{
		repository: repository,
	}
}

type finder struct {
	repository AssetRepository
}

func (f *finder) All() ([]Asset, error) {
	return f.repository.All()
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
