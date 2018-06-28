package keyvalue

//go:generate moq -out value_mock.go . Value
type Value interface {
	GetString() (string, error)
	PutString(value string) error
	GetInt64() (int64, error)
	PutInt64(value int64) error
}

type valueWrapper struct {
	kv  KeyValue
	key string
}

func WrapKeyValue(kv KeyValue, key string) Value {
	return &valueWrapper{
		kv:  kv,
		key: key,
	}
}

func (wrapper *valueWrapper) GetString() (string, error) {
	return wrapper.kv.GetString(wrapper.key)
}

func (wrapper *valueWrapper) PutString(value string) error {
	return wrapper.kv.PutString(wrapper.key, value)
}

func (wrapper *valueWrapper) GetInt64() (int64, error) {
	return wrapper.kv.GetInt64(wrapper.key)
}

func (wrapper *valueWrapper) PutInt64(value int64) error {
	return wrapper.kv.PutInt64(wrapper.key, value)
}
