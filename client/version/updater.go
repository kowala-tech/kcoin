package version

import (
	"archive/zip"
	"github.com/blang/semver"
	"github.com/kowala-tech/kcoin/client/log"
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
	repository string
	current    semver.Version
	finder     *finder
	logger     log.Logger
}

func NewUpdater(repository string, logger log.Logger) (*updater, error) {
	current, err := semver.Make(params.Version)
	if err != nil {
		return nil, err
	}

	return &updater{
		repository: repository,
		current:    current,
		finder:     NewFinder(repository),
		logger:     logger,
	}, nil
}

func (u *updater) isCurrentLatestForMajor() (bool, error) {
	latestAsset, err := u.latestAssetForMajor()
	if err != nil {
		return true, err
	}

	return u.current.GTE(latestAsset.Semver()), nil
}

func (u *updater) latestAssetForMajor() (Asset, error) {
	return u.finder.LatestForMajor(runtime.GOOS, runtime.GOARCH, u.current.Major)
}

func (u *updater) Update() error {
	latestAsset, err := u.latestAssetForMajor()
	if err != nil {
		return err
	}

	if !latestAsset.Semver().GT(u.current) {
		// up to date
		u.logger.Info("Nothing to do binary is at latest version")
		return nil
	}

	if err := u.download(latestAsset); err != nil {
		return err
	}

	if err := u.unzip(latestAsset); err != nil {
		return err
	}

	if err := u.backupCurrentBinary(); err != nil {
		return err
	}

	if err := u.replaceNewBinary(latestAsset); err != nil {
		return err
	}

	u.logger.Info("Client is up to date please start with your normal options")

	return nil
}

func (u *updater) download(asset Asset) error {
	assetUrl := u.repository + "/" + asset.Path()
	u.logger.Info("Downloading latest version")

	out, err := os.Create(asset.Path())
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

func (u *updater) unzip(asset Asset) error {
	r, err := zip.OpenReader(asset.Path())
	if err != nil {
		return err
	}
	defer r.Close()

	u.logger.Info("Unzipping file")

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

	u.logger.Info("Backing up binary")

	if err = os.Rename(file, backupFile); err != nil {
		return err
	}

	return nil
}

func (u *updater) replaceNewBinary(asset Asset) error {
	file := os.Args[0]

	dir, oldFilename := filepath.Split(file)

	absdir, err := filepath.Abs(dir)
	if err != nil {
		return err
	}

	filename := "kcoin-" + asset.Os() + "-" + asset.Arch()
	binary := absdir + "/" + oldFilename

	if err = os.Rename(filename, binary); err != nil {
		return err
	}

	return nil
}
