package keyvalue

import (
	"errors"
	"strconv"
)

//go:generate moq -out keyvalue_mock.go . KeyValue
type KeyValue interface {
	Delete(key string) error
	GetString(key string) (string, error)
	PutString(key string, value string) error
	GetInt64(key string) (int64, error)
	PutInt64(key string, value int64) error
}

var (
	ErrInvalidType = errors.New("Invalid value type")
)

func parseInt64Key(kv KeyValue, key string) (int64, error) {
	str, err := kv.GetString(key)
	if err != nil {
		return 0, err
	}
	if str == "" {
		return 0, nil
	}
	return strconv.ParseInt(str, 10, 64)
}
