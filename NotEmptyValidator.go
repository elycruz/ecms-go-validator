package ecms_validator

import (
	. "github.com/extensible-cms/ecms-go-validator/is"
)

type NotEmptyValidatorOptions struct {
	MessageTemplates map[int]MessageTemplateFunc
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

func NotEmptyValidator (options ValidatorOptions) Validator {
	return func (x interface{}) ValidationResult {
		if Empty(x) {
			return ValidationResult{false, []string{GetErrorMessageByKey(options, EmptyNotAllowed, x)}}
		}
		return ValidationResult{true, make([]string, 0)}
	}
}

func NotEmptyValidator1 () Validator {
	return NotEmptyValidator(NewNotEmptyValidatorOptions())
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
