package scraper

import (
	"github.com/pkg/errors"
)

func Init() error {
	return errors.New("hardware does not support the oracle service")
}

func Free() error               { return nil }
func GetPrice() ([]byte, error) { return nil, nil }
