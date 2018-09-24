package version

import (
	"github.com/blang/semver"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAllEmpty(t *testing.T) {
	repository := NewMemoryAssetRepository(nil)
	finder := NewFinder(repository)

	assets, err := finder.All()
	require.Nil(t, err)

	assert.Equal(t, 0, len(assets))
}

func TestAll(t *testing.T) {
	repoAssets := []Asset{
		NewAsset(semver.MustParse("1.0.0"), "linux", "arm64", "/"),
		NewAsset(semver.MustParse("1.0.0"), "linux", "amd64", "/"),
	}
	repository := NewMemoryAssetRepository(repoAssets)
	finder := NewFinder(repository)

	assets, err := finder.All()
	require.Nil(t, err)

	assert.Equal(t, 2, len(assets))
}

func TestLatest(t *testing.T) {
	repoAssets := []Asset{
		NewAsset(semver.MustParse("0.0.1"), "linux", "arm64", "1"),
		NewAsset(semver.MustParse("1.0.1"), "linux", "arm64", "2"),
		NewAsset(semver.MustParse("1.0.2"), "linux", "arm64", "3"),
		NewAsset(semver.MustParse("1.0.1"), "darwin", "arm64", "4"),
		NewAsset(semver.MustParse("1.0.5"), "darwin", "arm64", "5"),
	}
	finder := NewFinder(NewMemoryAssetRepository(repoAssets))

	asset, err := finder.Latest("linux", "arm64")
	require.Nil(t, err)

	assert.Equal(t, "1.0.2", asset.Semver().String())
	assert.Equal(t, "linux", asset.Os())
	assert.Equal(t, "arm64", asset.Arch())
	assert.Equal(t, "3", asset.Path())
}

func TestLatestMajor(t *testing.T) {
	repoAssets := []Asset{
		NewAsset(semver.MustParse("0.0.1"), "linux", "arm64", "1"),
		NewAsset(semver.MustParse("1.0.1"), "linux", "arm64", "2"),
		NewAsset(semver.MustParse("1.0.2"), "linux", "arm64", "3"),
		NewAsset(semver.MustParse("2.0.5"), "linux", "arm64", "4"),
		NewAsset(semver.MustParse("2.1.6"), "linux", "arm64", "5"),
		NewAsset(semver.MustParse("2.1.7"), "linux", "arm64", "6"),
	}
	finder := NewFinder(NewMemoryAssetRepository(repoAssets))

	asset, err := finder.LatestForMajor("linux", "arm64", 0)
	require.Nil(t, err)
	assert.Equal(t, "0.0.1", asset.Semver().String())
	assert.Equal(t, "linux", asset.Os())
	assert.Equal(t, "arm64", asset.Arch())
	assert.Equal(t, "1", asset.Path())

	asset, err = finder.LatestForMajor("linux", "arm64", 1)
	require.Nil(t, err)
	assert.Equal(t, "1.0.2", asset.Semver().String())
	assert.Equal(t, "linux", asset.Os())
	assert.Equal(t, "arm64", asset.Arch())
	assert.Equal(t, "3", asset.Path())

	asset, err = finder.LatestForMajor("linux", "arm64", 2)
	require.Nil(t, err)
	assert.Equal(t, "2.1.7", asset.Semver().String())
	assert.Equal(t, "linux", asset.Os())
	assert.Equal(t, "arm64", asset.Arch())
	assert.Equal(t, "6", asset.Path())
}
