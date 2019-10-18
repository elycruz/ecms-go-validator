package ecms_validator

import (
	"fmt"
	"regexp"
)

type RegexValidatorOptions struct {
	Pattern *regexp.Regexp
	MessageFuncs *MessageFuncs
}

var RegexValidatorMessageFuncs *MessageFuncs

func init() {
	RegexValidatorMessageFuncs = &MessageFuncs{
		DoesNotMatchPattern: func(options ValidatorOptions, x interface{}) string {
			ops := options.(RegexValidatorOptions)
			var pattern string
			if ops.Pattern != nil {
				pattern = ops.Pattern.String()
			}
			return fmt.Sprintf("%v does not match required pattern `%v`", x, pattern)
		},
	}
}

func NewRegexValidatorOptions () RegexValidatorOptions {
	return RegexValidatorOptions{
		Pattern: nil,
		MessageFuncs: RegexValidatorMessageFuncs,
	}
}

func RegexValidator(options RegexValidatorOptions) Validator {
	return func(x interface{}) (bool, []string) {
		if options.Pattern == nil {
			isValid := x == nil
			if !isValid {
				return false, []string {
					options.GetErrorMessageByKey(DoesNotMatchPattern, x),
				}
			}
			return true, nil
		}
		var match bool
		if options.Pattern != nil && x == nil {
			match = false
		} else {
			match = options.Pattern.Match([]byte(x.(string)))
		}

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

func (n RegexValidatorOptions) GetMessageFuncs() *MessageFuncs {
	return n.MessageFuncs
}

func (n RegexValidatorOptions) GetValueObscurator() ValueObscurator {
	return DefaultValueObscurator
}
