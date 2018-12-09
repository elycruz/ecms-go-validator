package ecms_validator

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
)

var (
	digitRegex *regexp.Regexp
	DigitValidatorMessageFuncs MessageTemplateFuncs
	digitValidator Validator
)

func init () {
	digitRegex = regexp.MustCompile("^\\d+$")
	DigitValidatorMessageFuncs = MessageTemplateFuncs{
		DoesNotMatchPattern: func(options ValidatorOptions, x interface{}) string {
			ops := options.(RegexValidatorOptions)
			return fmt.Sprintf("%v contains non-digital characters;  Received: %v", x, ops.Pattern.String())
		},
	}
	digitValidator = RegexValidator(newDigitValidatorOptions())
}

func newDigitValidatorOptions () RegexValidatorOptions {
	return RegexValidatorOptions{
		Pattern: digitRegex,
		MessageTemplates: &DigitValidatorMessageFuncs,
	}
}

func DigitValidator () Validator {
	return func (x interface{}) (bool, []string) {
		rv := reflect.ValueOf(x)
		k := rv.Kind()
		var converted string
		switch k {
		case reflect.Uint:
			converted = strconv.FormatUint(x.(uint64), 10)
		case reflect.Int:
			converted = strconv.FormatInt(rv.Int(), 10)
		case reflect.String:
			converted = x.(string)
			// @todo add byte-string check
			// @todo generate correct error message for invalid types here
		default:
			converted = ""
		}
		return digitValidator(converted)
	}
}
