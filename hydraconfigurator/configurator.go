package hydraconfigurator

import (
	"errors"
	"reflect"
)

const CUSTOM uint8 = iota

var wrongTypeError error = errors.New("type must be a pointer to a struct")

func GetCongiguration(confType uint8, obj interface{}, filename string) (err error) {
	rv := reflect.ValueOf(obj)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return wrongTypeError
	}

	rv = rv.Elem()
	if rv.Kind() != reflect.Struct {
		return wrongTypeError
	}

	switch confType {
	case CUSTOM:
		err = MarshalCustomConfig(rv, filename)
	}
	return err
}

