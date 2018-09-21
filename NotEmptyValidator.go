package ecms_validator

import (
	"reflect"
	"strings"
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
		failed := ValidationResult{false, []string{GetErrorMessageByKey(options, EmptyNotAllowed, x)}}
		if x == nil || x == 0 || x == false {
			return failed
		}
		typeName := reflect.TypeOf(x).Name()
		switch {
		case strings.HasPrefix(typeName, "map["):
			if len(x.(map[interface{}]interface{})) == 0 {
				return failed
			}
		case strings.HasPrefix(typeName, "[]"):
			if len(x.([]interface{})) == 0 {
				return failed
			}
		}
		return ValidationResult{true, make([]string, 0)}
	}
}

func NotEmptyValidator () Validator {
	return NotEmptyValidatorGenerator(NewNotEmptyValidatorOptions())
}
