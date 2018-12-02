package ecms_validator

import (
	"fmt"
	"regexp"
)

const (
	DoesNotMatchPattern = iota
)

type RegexValidatorOptions struct {
	Pattern *regexp.Regexp
	MessageTemplates *MessageTemplateFuncs
}

var RegexValidatorMessageFuncs *MessageTemplateFuncs

func init() {
	RegexValidatorMessageFuncs = &MessageTemplateFuncs{
		DoesNotMatchPattern: func(options ValidatorOptions, x interface{}) string {
			ops := options.(RegexValidatorOptions)
			return fmt.Sprintf("%v is does not match required pattern %v", x, ops.Pattern.String())
		},
	}
}

func NewRegexValidatorOptions () RegexValidatorOptions {
	return RegexValidatorOptions{
		Pattern: nil,
		MessageTemplates: RegexValidatorMessageFuncs,
	}
}

func RegexValidator(options RegexValidatorOptions) Validator {
	return func(x interface{}) (bool, []string) {
		if x == nil {
			return false, []string {
				options.GetErrorMessageByKey(DoesNotMatchPattern, ""),
			}
		}
		match := options.Pattern.Match([]byte(x.(string)))
		if match != true {
			return false, []string{
				options.GetErrorMessageByKey(DoesNotMatchPattern, x),
			}
		}
		return true, nil
	}
}

func (n RegexValidatorOptions) GetErrorMessageByKey(key int, x interface{}) string {
	return GetErrorMessageByKey(n, key, x)
}

func (n RegexValidatorOptions) GetMessageTemplates() *MessageTemplateFuncs {
	return n.MessageTemplates
}

func (n RegexValidatorOptions) GetValueObscurator() ValueObscurator {
	return DefaultValueObscurator
}
