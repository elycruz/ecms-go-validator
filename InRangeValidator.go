package ecms_validator

import (
	"fmt"
	. "github.com/extensible-cms/ecms-go-validator/is"
	"reflect"
)

const (
	NotWithinRange = iota
	NotARangeType
)

var DefaultInRangeMessageFuncs = MessageTemplateFuncs{
	NotARangeType: func(options ValidatorOptions, x interface{}) string {
		return fmt.Sprintf("%v is not a validatable numeric type", x)
	},
}

type IntValidatorOptions struct {
	MessageTemplates MessageTemplateFuncs
	Min              int64
	Max              int64
	Inclusive        bool
}

type FloatValidatorOptions struct {
	MessageTemplates MessageTemplateFuncs
	Min              float64
	Max              float64
	Inclusive        bool
}

func NewIntRangeValidatorOptions() IntValidatorOptions {
	return IntValidatorOptions{
		map[int]MessageTemplateFunc{
			NotWithinRange: func(options ValidatorOptions, x interface{}) string {
				ops := options.(IntValidatorOptions)
				return fmt.Sprintf("%v is not within range %d and %d", x, ops.Min, ops.Max)
			},
			NotARangeType: DefaultInRangeMessageFuncs[NotARangeType],
		},
		0,
		0,
		true,
	}
}

func NewFloatRangeValidatorOptions() FloatValidatorOptions {
	return FloatValidatorOptions{
		map[int]MessageTemplateFunc{
			NotWithinRange: func(options ValidatorOptions, x interface{}) string {
				ops := options.(FloatValidatorOptions)
				return fmt.Sprintf("%v is not within range %f and %f", x, ops.Min, ops.Max)
			},
			NotARangeType: DefaultInRangeMessageFuncs[NotARangeType],
		},
		0.0,
		0.0,
		true,
	}
}

func IntRangeValidator(options ValidatorOptions) Validator {
	ops := options.(IntValidatorOptions)
	return func(x interface{}) (bool, []string) {
		var intToCheck int64
		rv := reflect.ValueOf(x)
		switch rv.Kind() {
		case reflect.Invalid:
			return false, []string{ops.GetErrorMessageByKey(NotARangeType, x)}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			intToCheck = rv.Int()
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			intToCheck = int64(rv.Uint())
		default:
			return false, []string{ops.GetErrorMessageByKey(NotARangeType, x)}
		}
		if !IntWithinRange(ops.Min, ops.Max, intToCheck) {
			return false, []string{ops.GetErrorMessageByKey(NotWithinRange, x)}
		}
		return true, nil
	}
}

func FloatRangeValidator(options ValidatorOptions) Validator {
	ops := options.(FloatValidatorOptions)
	return func(x interface{}) (bool, []string) {
		rv := reflect.ValueOf(x)
		var floatToCheck float64
		switch rv.Kind() {
		case reflect.Invalid:
			return false, []string{ops.GetErrorMessageByKey(NotARangeType, x)}
		case reflect.Float32, reflect.Float64:
			floatToCheck = rv.Float()
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			floatToCheck = float64(rv.Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			floatToCheck = float64(rv.Uint())
		default:
			return false, []string{ops.GetErrorMessageByKey(NotARangeType, x)}
		}
		if !FloatWithinRange(ops.Min, ops.Max, floatToCheck) {
			return false, []string{ops.GetErrorMessageByKey(NotWithinRange, x)}
		}
		return true, nil
	}
}

func (n IntValidatorOptions) GetErrorMessageByKey(key int, x interface{}) string {
	return GetErrorMessageByKey(n, key, x)
}

func (n IntValidatorOptions) GetMessageTemplates() MessageTemplateFuncs {
	return n.MessageTemplates
}

func (n IntValidatorOptions) GetValueObscurator() ValueObscurator {
	return DefaultValueObscurator
}

func (n FloatValidatorOptions) GetErrorMessageByKey(key int, x interface{}) string {
	return GetErrorMessageByKey(n, key, x)
}

func (n FloatValidatorOptions) GetMessageTemplates() MessageTemplateFuncs {
	return n.MessageTemplates
}

func (n FloatValidatorOptions) GetValueObscurator() ValueObscurator {
	return DefaultValueObscurator
}
