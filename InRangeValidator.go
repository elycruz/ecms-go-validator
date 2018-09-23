package ecms_validator

import (
"fmt"
"reflect"
)

const (
	NotWithinRange = iota
	NotARangeType
)

var DefaultInRangeMessageFuncs =
	MessageTemplateFuncs{
		NotWithinRange: func(options ValidatorOptions, x interface{}) string {
			ops := options.(IntValidatorOptions)
			return fmt.Sprintf("%v is not within range %d - %d", x, ops.Min, ops.Max)
		},
		NotARangeType: func(options ValidatorOptions, x interface{}) string {
			return fmt.Sprintf("%v is not a numeric type", x)
		},
	}

type IntValidatorOptions struct {
	MessageTemplates MessageTemplateFuncs
	Min int64
	Max int64
	Inclusive bool
}

type FloatValidatorOptions struct {
	MessageTemplates MessageTemplateFuncs
	Min float64
	Max float64
	Inclusive bool
}

type ComplexValidatorOptions struct {
	MessageTemplates MessageTemplateFuncs
	Min complex128
	Max complex128
	Inclusive bool
}

func NewIntRangeValidatorOptions () IntValidatorOptions {
	return IntValidatorOptions{
		DefaultInRangeMessageFuncs,
		0,
		0,
		true,
	}
}

func IntRangeValidator (options ValidatorOptions) Validator {
	ops := options.(IntValidatorOptions)
	return func(x interface{}) ValidationResult {
		rv := reflect.ValueOf(x)
		switch rv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if !isWithinRangeInt(ops.Min, ops.Max, rv.Int()) {
				return ValidationResult{
					false,
					[]string{ops.GetErrorMessageByKey(NotWithinRange, x)},
				}
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			if !isWithinRangeInt(ops.Min, ops.Max, int64(rv.Uint())) {
				return ValidationResult{
					false,
					[]string{ops.GetErrorMessageByKey(NotWithinRange, x)},
				}
			}
		default:
			return ValidationResult{
				false,
				[]string{ops.GetErrorMessageByKey(NotARangeType, x)},
			}
		}
		return ValidationResult{true, make([]string, 0)}
	}
}

func FloatRangeValidator (options ValidatorOptions) Validator {
	ops := options.(FloatValidatorOptions)
	return func(x interface{}) ValidationResult {
		rv := reflect.ValueOf(x)
		switch rv.Kind() {
		case reflect.Float32, reflect.Float64:
			if !isWithinRangeFloat(ops.Min, ops.Max, rv.Float()) {
				return ValidationResult{
					false,
					[]string{ops.GetErrorMessageByKey(NotWithinRange, x)},
				}
			}
		case reflect.Complex64, reflect.Complex128:
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if !isWithinRangeFloat(ops.Min, ops.Max, float64(rv.Int())) {
				return ValidationResult{
					false,
					[]string{ops.GetErrorMessageByKey(NotWithinRange, x)},
				}
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			if !isWithinRangeFloat(ops.Min, ops.Max, float64(rv.Uint())) {
				return ValidationResult{
					false,
					[]string{ops.GetErrorMessageByKey(NotWithinRange, x)},
				}
			}

		default:
			return ValidationResult{
				false,
				[]string{ops.GetErrorMessageByKey(NotARangeType, x)},
			}
		}
		return ValidationResult{true, make([]string, 0)}
	}
}

func (n IntValidatorOptions) GetErrorMessageByKey (key int, x interface{}) string {
	return GetErrorMessageByKey(n, key, x)
}

func (n IntValidatorOptions) GetMessageTemplates () MessageTemplateFuncs {
	return n.MessageTemplates
}

func (n IntValidatorOptions) GetValueObscurator () ValueObscurator {
	return DefaultValueObscurator
}

func (n FloatValidatorOptions) GetErrorMessageByKey (key int, x interface{}) string {
	return GetErrorMessageByKey(n, key, x)
}

func (n FloatValidatorOptions) GetMessageTemplates () MessageTemplateFuncs {
	return n.MessageTemplates
}

func (n FloatValidatorOptions) GetValueObscurator () ValueObscurator {
	return DefaultValueObscurator
}
