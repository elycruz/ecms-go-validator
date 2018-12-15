package ecms_validator

import (
	"fmt"
	"reflect"
	"regexp"
)

var (
	digitRegex                 *regexp.Regexp
	DigitValidatorMessageFuncs MessageFuncs
)

func init () {
	digitRegex = regexp.MustCompile("^\\d+$")
	DigitValidatorMessageFuncs = MessageFuncs{
		DoesNotMatchPattern: func(options ValidatorOptions, x interface{}) string {
			ops := options.(RegexValidatorOptions)
			return fmt.Sprintf("%v contains non-digital characters;  Received: %v", x, ops.Pattern.String())
		},
	}
}

func newDigitValidatorOptions () RegexValidatorOptions {
	return RegexValidatorOptions{
		Pattern: digitRegex,
		MessageFuncs: &DigitValidatorMessageFuncs,
	}
}

// DigitValidator - Returns `(true, nil)` for `uint`, `int`, and strings containing
// only digit characters (numbers).  For every other value the returned validator
// will always return `(false, []string{})` where the second return value is
// the error messages returned for current run.
// Note: The `Pattern` property of passed in `RegexValidatorOptions` gets ignored
// and the internally defined one gets used.
func DigitValidator (options RegexValidatorOptions) Validator {
	return func (x interface{}) (bool, []string) {
		rv := reflect.ValueOf(x)
		k := rv.Kind()
		switch k {
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return true, nil
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if rv.Int() < 0 {
				return false, []string{options.GetErrorMessageByKey(DoesNotMatchPattern, x)}
			}
			return true, nil
		case reflect.String:
			ops := newDigitValidatorOptions()
			ops.MessageFuncs = options.GetMessageFuncs()
			return RegexValidator(ops)(x.(string))
		default:
			return false, []string{options.GetErrorMessageByKey(DoesNotMatchPattern, x)}
		}
	}
}

// DigitValidator1 - Ignores options param and just returns a validator
// which contains default error messages in message templates
func DigitValidator1 () Validator {
	return DigitValidator(newDigitValidatorOptions())
}
