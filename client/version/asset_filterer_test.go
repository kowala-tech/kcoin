package version

import (
	"github.com/blang/semver"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlatform(t *testing.T) {
	assets := []Asset{
		NewAsset(semver.MustParse("1.0.0"), "linux", "arm64", "/"),
		NewAsset(semver.MustParse("1.0.0"), "linux", "amd64", "/"),
		NewAsset(semver.MustParse("1.0.0"), "darwin", "amd64", "/"),
		NewAsset(semver.MustParse("1.0.1"), "darwin", "amd64", "/"),
		NewAsset(semver.MustParse("1.0.2"), "darwin", "amd64", "/"),
	}

	filteredAssets := NewAssetFilterer(assets).by(platform("darwin", "amd64"))
	assert.Len(t, filteredAssets, 3)

	filteredAssets = NewAssetFilterer(assets).by(platform("darwin", "arm64"))
	assert.Len(t, filteredAssets, 0)

	filteredAssets = NewAssetFilterer(assets).by(platform("linux", "amd64"))
	assert.Len(t, filteredAssets, 1)

	filteredAssets = NewAssetFilterer(assets).by(platform("linux", "arm64"))
	assert.Len(t, filteredAssets, 1)

	filteredAssets = NewAssetFilterer(assets).by(platform("windows", "amd64"))
	assert.Len(t, filteredAssets, 0)
}

func TestPlatformMajor(t *testing.T) {
	assets := []Asset{
		NewAsset(semver.MustParse("1.0.0"), "darwin", "amd64", "/"),
		NewAsset(semver.MustParse("1.0.1"), "darwin", "amd64", "/"),
		NewAsset(semver.MustParse("1.0.2"), "darwin", "amd64", "/"),
		NewAsset(semver.MustParse("2.0.0"), "darwin", "amd64", "/"),
	}

	filteredAssets := NewAssetFilterer(assets).by(platformMajor("darwin", "amd64", 1))
	assert.Len(t, filteredAssets, 3)

	filteredAssets = NewAssetFilterer(assets).by(platformMajor("darwin", "amd64", 2))
	assert.Len(t, filteredAssets, 1)
}
