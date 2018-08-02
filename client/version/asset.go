package version

import "github.com/blang/semver"

type Asset interface {
	Semver() semver.Version
	Os() string
	Arch() string
	Path() string
}

func NewAsset(semver semver.Version, os, arch, path string) Asset {
	return asset{
		semver: semver,
		os:     os,
		arch:   arch,
		path:   path,
	}
}

type asset struct {
	semver semver.Version
	os     string
	arch   string
	path   string
}

func (c asset) Semver() semver.Version {
	return c.semver
}

func (c asset) Os() string {
	return c.os
}

func (c asset) Arch() string {
	return c.arch
}

func (c asset) Path() string {
	return c.path
}
