package ecms_validator

import "reflect"

func IsEmpty (x interface{}) bool {
	if x == nil || x == 0 || x == false {
		return true
	}
	rv := reflect.ValueOf(x)
	rk := rv.Kind()
	switch rk {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		if rv.Len() == 0 {
			return true
		}
	}
	return false
}
