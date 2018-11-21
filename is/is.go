package is

import (
	"reflect"
)

func Empty(x interface{}) bool {
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
		if numFields == 0 { return true }
		for ind := 0; ind < numFields; ind += 1 {
			if !Empty(rv.Field(ind).Interface()) {
				return false
			}
		}
		return true
	case reflect.Interface, reflect.Ptr:
		if rv.IsNil() {
			return true
		}
		return Empty(rv.Elem().Interface())
	}
	return false
}

func IntWithinRange(min, max, x int64) bool {
	xMin, xMax := normalizeMinMaxForInt(min, max)
	return x >= xMin && x <= xMax
}

func FloatWithinRange(min, max, x float64) bool {
	xMin, xMax := normalizeMinMaxForFloat(min, max)
	return x >= xMin && x <= xMax
}

func normalizeMinMaxForInt(min, max int64) (int64, int64) {
	if max < min {
		return max, min
	}
	return min, max
}

func normalizeMinMaxForFloat(min, max float64) (float64, float64) {
	if max < min {
		return max, min
	}
	return min, max
}
