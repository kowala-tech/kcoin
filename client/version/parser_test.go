package version

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFilenameParser(t *testing.T) {
	filenames := []struct {
		filename string
		version  string
		os       string
		arch     string
	}{
		{"kcoin-1.0.11-linux-amd64.zip", "1.0.11", "linux", "amd64"},
		{"kcoin-1.0.14-linux-amd64.zip", "1.0.14", "linux", "amd64"},
		{"kcoin-1.0.0-windows-4.0-amd64.zip", "1.0.0", "windows-4.0", "amd64"},
		{"kcoin-1.0.0-windows-4.0-amd64.exe.zip", "1.0.0", "windows-4.0", "amd64"},
		{"kcoin-1.0.0-osx-10.6-amd64.zip", "1.0.0", "osx-10.6", "amd64"},
	}
	for _, bm := range filenames {
		t.Run(bm.filename, func(t *testing.T) {
			asset, err := filenameParser(bm.filename)
			assert.NoError(t, err)
			assert.Equal(t, bm.filename, asset.Path())
			assert.Equal(t, bm.version, asset.Semver().String())
			assert.Equal(t, bm.os, asset.Os())
			assert.Equal(t, bm.arch, asset.Arch())
		})
	}
}
