package ecms_validator

import (
	"reflect"
)

type NotEmptyValidatorOptions struct {
	MessageTemplates map[int]MessageTemplateFunc
}

func (n NotEmptyValidatorOptions) GetErrorMessageByKey (key int, x interface{}) string {
	return GetErrorMessageByKey(n, key, x)
}

func (n NotEmptyValidatorOptions) GetMessageTemplates () MessageTemplateFuncs {
	return n.MessageTemplates
}

func (n NotEmptyValidatorOptions) GetValueObscurator () ValueObscurator {
	return DefaultValueObscurator
}

const (
	EmptyNotAllowed = iota
)

func isEmpty (x interface{}) bool {
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

func NewNotEmptyValidatorOptions () NotEmptyValidatorOptions {
	messageTemplates := map[int]MessageTemplateFunc{
		EmptyNotAllowed: func(ops ValidatorOptions, x interface{}) string {
			return "Empty values are not allowed."
		},
	}
	return NotEmptyValidatorOptions{
		MessageTemplates: messageTemplates,
	}
}

func NotEmptyValidatorGenerator (options ValidatorOptions) Validator {
	return func (x interface{}) ValidationResult {
		if isEmpty(x) {
			return ValidationResult{false, []string{GetErrorMessageByKey(options, EmptyNotAllowed, x)}}
		}
		return ValidationResult{true, make([]string, 0)}
	}
}

func NotEmptyValidator () Validator {
	return NotEmptyValidatorGenerator(NewNotEmptyValidatorOptions())
}
