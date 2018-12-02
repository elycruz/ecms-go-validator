package ecms_validator

import "regexp"

type RegexValidatorOptions struct {
	Pattern regexp.Regexp
	MessageTemplates map[int]MessageTemplateFunc
}

func RegexValidator(options RegexValidatorOptions) (bool, []string) {
	return false, nil
}
