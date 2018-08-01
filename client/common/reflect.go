package common

import (
	"reflect"
	"errors"
)

// SafeValueOf returns valueOf is interface is not nil value or nil pointer
func SafeValueOf(a interface{}) (reflect.Value, error) {
	v := reflect.ValueOf(a)
	if a == nil || (v.Kind() == reflect.Ptr && v.IsNil()) {
		return reflect.Value{}, errors.New("got `nil` argument")
	}
	return v, nil
}
