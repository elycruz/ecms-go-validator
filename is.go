package ecms_validator

import "reflect"

func IsEmpty (x interface{}) bool {
	if x == nil {
		return true
	}
	rv := reflect.ValueOf(x)
	switch rv.Kind() {
	case reflect.Bool:
		return rv.Bool() == false
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if rv.Int() <= 0 { return true }
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if rv.Uint() <= 0 { return true }
	case reflect.Float32, reflect.Float64:
		if rv.Float() <= 0.0 { return true }
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		if rv.Len() == 0 { return true }
	case reflect.Complex64, reflect.Complex128:
		var c complex128 = 0
		if rv.Complex() == c { return true }
	}
	return false
}
