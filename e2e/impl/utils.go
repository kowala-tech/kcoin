package impl

import (
	"math/big"
	"os"
	"path"

	"github.com/kowala-tech/kcoin/client/params"
)

func toWei(kcoin int64) *big.Int {
	return new(big.Int).Mul(big.NewInt(kcoin), big.NewInt(params.Kcoin))
}

func createDir(dir string) error {
	err := os.Mkdir(dir, 0755)
	if os.IsExist(err) {
		return nil
	}
	return err
}

func clearDir(dir string) error {
	dirRead, err := os.Open(dir)
	if err != nil {
		return err
	}

	const innerLevel = 0
	dirFiles, err := dirRead.Readdir(innerLevel)
	if err != nil {
		return err
	}

	for _, file := range dirFiles {
		filePath := path.Join(dir, file.Name())
		err = os.Remove(filePath)
		if err != nil {
			return err
		}
	}

	return nil
}
