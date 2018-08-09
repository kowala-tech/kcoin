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

	if err := u.download(); err != nil {
		return err
	}

	if err := u.unzip(); err != nil {
		return err
	}

	if err := u.backupCurrentBinary(); err != nil {
		return err
	}

	if err := u.replaceNewBinary(); err != nil {
		return err
	}

	fmt.Println("Client is up to date please start with your normal options")

	return nil
}

func (u *updater) download() error {
	assetUrl := u.repository + "/" + u.latestAsset.Path()
	fmt.Println("downloading latest version")

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

	if _, err = io.Copy(out, response.Body); err != nil {
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

	fmt.Println("unziping file")

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

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

func (u *updater) backupCurrentBinary() error {
	file := os.Args[0]

	dir, filename := filepath.Split(file)

	absdir, err := filepath.Abs(dir)
	if err != nil {
		return err
	}

	backupFile := absdir + "/backup_" + filename

	fmt.Println("backing up binary")

	if err = os.Rename(file, backupFile); err != nil {
		return err
	}

	return nil
}

func (u *updater) replaceNewBinary() error {
	file := os.Args[0]

	dir, oldFilename := filepath.Split(file)

	absdir, err := filepath.Abs(dir)
	if err != nil {
		return err
	}

	filename := "kcoin-" + u.latestAsset.Os() + "-" + u.latestAsset.Arch()
	binary := absdir + "/" + oldFilename

	if err = os.Rename(filename, binary); err != nil {
		return err
	}

	return nil
}
