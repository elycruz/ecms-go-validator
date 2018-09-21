package ecms_validator

import (
	"reflect"
)

func IsEmpty (x interface{}) bool {
	rv := reflect.ValueOf(x)
	switch rv.Kind() {
	case reflect.Invalid:
		return true
	case reflect.Bool:
		return !rv.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return rv.Int() <= 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return rv.Uint() <= 0
	case reflect.Float32, reflect.Float64:
		return rv.Float() <= 0.0
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		return rv.Len() == 0
	case reflect.Complex64, reflect.Complex128:
		var c complex128 = 0
		return rv.Complex() == c
	case reflect.Struct:
		numFields := rv.NumField()
		for ind := 0; ind < numFields; ind += 1 {
			if !IsEmpty(rv.Field(ind).Interface()) {
				return false
			}
		}
		return true
	case reflect.Interface, reflect.Ptr:
		if rv.IsNil() {
			return true
		}
		return IsEmpty(rv.Elem().Interface())
	}
	return false
}

func isWithinRange (min, max, x int) bool {
	return x >= min && x <= max
}
