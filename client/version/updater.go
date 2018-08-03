package version

import (
	"archive/zip"
	"fmt"
	"github.com/blang/semver"
	"github.com/kowala-tech/kcoin/client/params"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

type Updater interface {
	Update() error
}

type updater struct {
	repository  string
	current     semver.Version
	latestAsset Asset
}

func NewUpdater(repository string) (*updater, error) {
	current, err := semver.Make(params.Version)
	if err != nil {
		return nil, err
	}

	finder := NewFinder(repository)
	latest, err := finder.Latest(runtime.GOOS, runtime.GOARCH)
	if err != nil {
		return nil, err
	}

	return &updater{
		repository:  repository,
		current:     current,
		latestAsset: latest,
	}, nil
}

func (u *updater) Update() error {
	if !u.latestAsset.Semver().GT(u.current) {
		// up to date
		fmt.Println("Nothing to do binary is at latest version")
		return nil
	}

	err := u.download()
	if err != nil {
		return err
	}

	err = u.unzip()

	return err
}

func (u *updater) download() error {
	assetUrl := u.repository + "/" + u.latestAsset.Path()
	fmt.Println("downloading latest version asset " + assetUrl)

	out, err := os.Create(u.latestAsset.Path())
	if err != nil {
		return err
	}
	defer out.Close()

	response, err := http.Get(assetUrl)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	_, err = io.Copy(out, response.Body)
	if err != nil {
		return err
	}

	return nil
}

func (u *updater) unzip() error {
	r, err := zip.OpenReader(u.latestAsset.Path())
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		fmt.Println("unziping file " + f.Name)

		path := filepath.Join("", f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			f, err := os.OpenFile(
				path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer f.Close()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
