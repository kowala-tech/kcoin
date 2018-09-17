package version

import (
	"github.com/blang/semver"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAsset(t *testing.T) {
	asset := NewAsset(semver.MustParse("4.1.1"), "linux", "arm64", "1")

	assert.Equal(t, "4.1.1", asset.Semver().String())
	assert.Equal(t, "linux", asset.Os())
	assert.Equal(t, "arm64", asset.Arch())
	assert.Equal(t, "1", asset.Path())
}
