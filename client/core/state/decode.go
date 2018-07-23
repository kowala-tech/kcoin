package state

import (
	"fmt"
	"math/big"
	"reflect"
	"runtime"

	"github.com/kowala-tech/kcoin/client/common"
)

var (
	firstKey = common.Big0
)

type Unmarshaler interface {
	UnmarshalState([]byte) error
}

type InvalidUnmarshalError struct {
	Type reflect.Type
}

func (e *InvalidUnmarshalError) Error() string {
	if e.Type == nil {
		return "state: Unmarshal(nil)"
	}

	if e.Type.Kind() != reflect.Ptr {
		return "state: Unmarshal(non-pointer " + e.Type.String() + ")"
	}

	return "state: Unmarshal(nil " + e.Type.String() + ")"
}

type decodeState struct {
	*stateObject
	key *big.Int
}

func (d *decodeState) init(data *stateObject) *decodeState {
	d.stateObject = data
	d.key = firstKey
	return d
}

func (d *decodeState) read() (ret []byte) {
	ret = d.GetState(d.db.db, common.BytesToHash(d.key.Bytes())).Bytes()
	d.key.Add(d.key, common.Big1)
}

func (d *decodeState) value(rv *reflect.Value) error {
	switch rv.Interface().(type) {
	case *big.Int:
		rv.Set(reflect.ValueOf(new(big.Int).SetBytes(d.read())))
		return nil
	}

	switch rv.Kind() {
	case reflect.Struct:
		rtype := rv.Type()
		for i := 0; i < rv.NumField(); i++ {
			fieldValue := rv.Field(i)
			structField := rtype.Field(i)
			if err := d.value(&fieldValue); err != nil {
				return err
			}
		}
		return nil
	}

	return nil
}

func (so *stateObject) UnmarshallState(v interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(runtime.Error); ok {
				panic(r)
			}
			err = r.(error)
		}
	}()

	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return &InvalidUnmarshalError{reflect.TypeOf(v)}
	}

	rv = rv.Elem()
	if rv.Kind() != reflect.Struct {
		fmt.Errorf("state: Unmarshal(non-struct %s)", rv.Type().String())
	}

	var d decodeState
	d.init(so)
	return d.value(&rv)
}
