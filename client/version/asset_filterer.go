package version

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

type assetFilterer struct {
	assets []Asset
}

func NewAssetFilterer(assets []Asset) assetFilterer {
	return assetFilterer{assets: assets}
}

func (a assetFilterer) by(allowFilter assetFilterFunc) []Asset {
	var filteredAssets []Asset
	for _, asset := range a.assets {
		if allowFilter(asset) {
			filteredAssets = append(filteredAssets, asset)
		}
	}
	return filteredAssets
}
