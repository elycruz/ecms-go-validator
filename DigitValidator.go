package ecms_validator

import (
	"fmt"
	"regexp"
)

var (
	digitRegex *regexp.Regexp
	DigitValidatorMessageFuncs MessageTemplateFuncs
)

func init () {
	digitRegex = regexp.MustCompile("^\\d+$")
	DigitValidatorMessageFuncs = MessageTemplateFuncs{
		DoesNotMatchPattern: func(options ValidatorOptions, x interface{}) string {
			ops := options.(RegexValidatorOptions)
			return fmt.Sprintf("%v contains non-digital characters;  Received: %v", x, ops.Pattern.String())
		},
	}
}

func newDigitValidatorOptions () RegexValidatorOptions {
	return RegexValidatorOptions{
		Pattern: digitRegex,
		MessageTemplates: DigitValidatorMessageFuncs,
	}
}

func DigitValidator (options RegexValidatorOptions) Validator {
	// @todo change `MessageTemplates` to pointer
	ops := newDigitValidatorOptions()
	return RegexValidator(ops)
}

func DigitValidator1 () Validator {
	ops := newDigitValidatorOptions()
	return RegexValidator(ops)
}
