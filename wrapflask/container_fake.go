//+build svcw

package wrapflask

import (
	"reflect"
	"unsafe"
)

func Wrap(in interface{}) {
	wrap(reflect.ValueOf(in))
	return
}

// walk the struct objects tree and replace all dependencies with wrapped one
func wrap(v reflect.Value) {
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	nf := v.NumField()
	for i := 0; i < nf; i++ {
		itf := v.Field(i)
		if itf.Kind() != reflect.Interface || itf.IsNil() {
			continue
		}
		tf := itf.Elem()
		rtf := tf
		for tf.Kind() == reflect.Ptr {
			tf = tf.Elem()
		}
		if v.Kind() != reflect.Struct || !v.IsValid() {
			continue
		}
		wrapType, ok := container[itf.Type()]
		if !ok || wrapType == tf.Type() {
			continue
		}
		wrapV := reflect.New(wrapType).Elem()
		if itf.CanSet() {
			wrapV.Field(0).Set(itf)
			itf.Set(wrapV)
		} else {
			if v.CanAddr() {
				fv := getUnexportedField(itf, v.UnsafeAddr()+v.Type().Field(i).Offset)
				wrapV.Field(0).Set(fv)
				setUnexportedField(itf, v.UnsafeAddr()+v.Type().Field(i).Offset, wrapV)
			}
		}
		wrap(rtf)
		continue
	}
}

func getUnexportedField(field reflect.Value, ptr uintptr) reflect.Value {
	return reflect.NewAt(field.Type(), unsafe.Pointer(ptr)).Elem()
}

func setUnexportedField(field reflect.Value, ptr uintptr, value reflect.Value) {
	reflect.NewAt(field.Type(), unsafe.Pointer(ptr)).
		Elem().
		Set(value)
}
