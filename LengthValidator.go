package ecms_validator

import (
	"fmt"
	"reflect"
)

type LengthValidatorOptions struct {
	MessageFuncs *MessageFuncs
	Min              int64
	Max              int64
	Inclusive        bool
}

var DefaultLengthValidatorMessageFuncs MessageFuncs

func init() {
	DefaultLengthValidatorMessageFuncs = MessageFuncs{
		NotAValidType: func(options ValidatorOptions, x interface{}) string {
			return fmt.Sprintf("%v is not a lengthable value type;  "+
				"Expected an array, a slice, a map, or a string value.", x)
		},
		NotWithinRange: func(options ValidatorOptions, x interface{}) string {
			ops := options.(IntValidatorOptions) // We pass the message getter to IntRangeValidator instance hence ...
			return fmt.Sprintf("%v is not within given length range of %d and %d", x, ops.Min, ops.Max)
		},
	}
}

func NewLengthValidatorOptions() LengthValidatorOptions {
	return LengthValidatorOptions{
		MessageFuncs: &DefaultLengthValidatorMessageFuncs,
		Min: 0,
		Max: 0,
		Inclusive: true,
	}
}

func LengthValidator (options LengthValidatorOptions) Validator {
	ops := NewIntRangeValidatorOptions()
	ops.Min = options.Min
	ops.Max = options.Max
	ops.Inclusive = options.Inclusive
	ops.MessageFuncs = options.MessageFuncs
	return func(x interface{}) (b bool, strings []string) {
		rv := reflect.ValueOf(x)
		var intToCheck int64
		switch rv.Kind() {
		case reflect.Invalid:
			return false, []string{ops.GetErrorMessageByKey(NotAValidType, x)}
		case reflect.Slice, reflect.Array, reflect.Map, reflect.String, reflect.Chan:
			intToCheck = int64(rv.Len())
		default:
			return false, []string{ops.GetErrorMessageByKey(NotAValidType, x)}
		}
		return IntRangeValidator(ops)(intToCheck)
	}
}

func (n LengthValidatorOptions) GetErrorMessageByKey(key int, x interface{}) string {
	return GetErrorMessageByKey(n, key, x)
}

func (n LengthValidatorOptions) GetMessageFuncs() *MessageFuncs {
	return n.MessageFuncs
}

func (n LengthValidatorOptions) GetValueObscurator() ValueObscurator {
	return DefaultValueObscurator
}
