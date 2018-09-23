package ecms_validator

import (
	"fmt"
	"reflect"
)

const (
	LengthNotWithinRange = iota
	NotLengthable
)

type LengthValidatorOptions struct {
	MessageTemplates MessageTemplateFuncs
	Min int
	Max int
}

func (n LengthValidatorOptions) GetErrorMessageByKey (key int, x interface{}) string {
	return GetErrorMessageByKey(n, key, x)
}

func (n LengthValidatorOptions) GetMessageTemplates () MessageTemplateFuncs {
	return n.MessageTemplates
}

func (n LengthValidatorOptions) GetValueObscurator () ValueObscurator {
	return DefaultValueObscurator
}

func NewLengthValidatorOptions () LengthValidatorOptions {
	return LengthValidatorOptions{
		MessageTemplateFuncs{
			LengthNotWithinRange: func(options ValidatorOptions, x interface{}) string {
				ops := options.(LengthValidatorOptions)
				return fmt.Sprintf("%v is not within range %d - %d", x, ops.Min, ops.Max)
			},
		},
		0,
		0,
	}
}

func LengthValidator (options ValidatorOptions) Validator {
	ops := options.(LengthValidatorOptions)
	return func(x interface{}) ValidationResult {
		rv := reflect.ValueOf(x)
		switch rv.Kind() {
		case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
			l := rv.Len()
			if !isWithinRangeInt(ops.Min, ops.Max, l) {
				return ValidationResult{
					false,
					[]string{ops.GetErrorMessageByKey(LengthNotWithinRange, x)},
				}
			}
		default:
			return ValidationResult{
				false,
				[]string{ops.GetErrorMessageByKey(NotLengthable, x)},
			}
		}
		return ValidationResult{true, make([]string, 0)}
	}
}
