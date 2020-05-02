package wrapflask

import (
	"reflect"
	"sync"
)

var (
	container = map[reflect.Type]reflect.Type{}
	mutex     = sync.Mutex{}
)

func Register(in interface{}, as interface{}) interface{} {
	inV := reflect.ValueOf(in)
	asV := reflect.ValueOf(as)
	for asV.Kind() == reflect.Ptr {
		asV = asV.Elem()
	}
	if !inV.Type().Implements(asV.Type()) || inV.NumField() != 1 && inV.Field(0).Type() != asV.Type() {
		return nil
	}
	mutex.Lock()
	container[asV.Type()] = inV.Type()
	mutex.Unlock()
	return nil
}
