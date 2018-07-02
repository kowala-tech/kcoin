package keyvalue

import (
	"strconv"
	"sync"
)

type memory struct {
	mtx  sync.RWMutex
	data map[string]string
}

func NewMemoryKeyValue() KeyValue {
	return &memory{
		data: make(map[string]string),
	}
}

func (kv *memory) GetString(key string) (string, error) {
	kv.mtx.RLock()
	defer kv.mtx.RUnlock()
	return kv.data[key], nil
}

func (kv *memory) PutString(key string, value string) error {
	kv.mtx.Lock()
	defer kv.mtx.Unlock()
	kv.data[key] = value
	return nil
}

func (kv *memory) GetInt64(key string) (int64, error) {
	return parseInt64Key(kv, key)
}

func (kv *memory) PutInt64(key string, value int64) error {
	return kv.PutString(key, strconv.FormatInt(value, 10))
}

func (kv *memory) Delete(key string) error {
	kv.mtx.RLock()
	defer kv.mtx.RUnlock()
	delete(kv.data, key)
	return nil
}
