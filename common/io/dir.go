package io

import (
	"os"
	"fmt"
)

func EnsureDir(dir string, mode os.FileMode) error {
	_, err := os.Stat(dir)
	if !os.IsNotExist(err) {
		return nil
	}

	err = os.MkdirAll(dir, mode)
	if err != nil {
		return fmt.Errorf("Could not create directory %v. %v", dir, err)
	}
	return nil
}