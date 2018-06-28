package keyvalue

import (
	"github.com/go-redis/redis"
)

type redisNsKV struct {
	namespace string
	client    *redis.Client
}

// NewRedisNamespacedKeyValue Returns a redis based KeyValue that uses `namespace` key as a hash to store everything
func NewRedisNamespacedKeyValue(client *redis.Client, namespace string) KeyValue {
	return &redisNsKV{
		namespace: namespace,
		client:    client,
	}
}

func (kv *redisNsKV) GetString(key string) (string, error) {
	res := kv.client.HGet(kv.namespace, key)
	if res.Err() == redis.Nil {
		return "", nil
	}
	return res.Result()
}

func (kv *redisNsKV) PutString(key string, value string) error {
	res := kv.client.HSet(kv.namespace, key, value)
	return res.Err()
}

func (kv *redisNsKV) GetInt64(key string) (int64, error) {
	return parseInt64Key(kv, key)
}

func (kv *redisNsKV) PutInt64(key string, value int64) error {
	res := kv.client.HSet(kv.namespace, key, value)
	return res.Err()
}

func (kv *redisNsKV) Delete(key string) error {
	res := kv.client.HDel(kv.namespace, key)
	return res.Err()
}
