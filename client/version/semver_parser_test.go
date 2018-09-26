package version

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMakeSemver(t *testing.T) {
	semver, err := MakeSemver("2.0.0")
	assert.NoError(t, err)

	assert.Equal(t, uint64(2), semver.Major)
	assert.Equal(t, uint64(0), semver.Minor)
	assert.Equal(t, uint64(0), semver.Patch)
	assert.Len(t, semver.Pre, 0)
}

func TestMakeSemverIgnoresPre(t *testing.T) {
	semver, err := MakeSemver("2.0.1-stable")
	assert.NoError(t, err)

	assert.Equal(t, uint64(2), semver.Major)
	assert.Equal(t, uint64(0), semver.Minor)
	assert.Equal(t, uint64(1), semver.Patch)
	assert.Len(t, semver.Pre, 0)
}
