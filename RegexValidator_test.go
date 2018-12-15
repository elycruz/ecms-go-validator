package ecms_validator

import (
	"regexp"
	"testing"
)

type RegexValidatorTestCase struct {
	Name 			string
	PatternStr      string
	Value           interface{}
	Expected        bool
	ExpectedMsgsLen int
}

func TestNewRegexValidatorOptions(t *testing.T) {
	result := NewRegexValidatorOptions()
	if result.Pattern != nil {
		t.Errorf("Expected `*.PatternStr` to equal `nil`;  Received: %v", result.Pattern)
	}
}

func TestRegexValidator(t *testing.T) {
	for _, tc := range []RegexValidatorTestCase{
		{Name: "//_$_true", PatternStr: "", Value: "$", Expected: false, ExpectedMsgsLen: 1},
		{Name: "//_''_true", PatternStr: "", Value: "", Expected: false, ExpectedMsgsLen: 1},
		{Name: "/./_$_true", PatternStr: ".", Value: "$", Expected: true, ExpectedMsgsLen: 0},
		{Name: "/./_''_false", PatternStr: ".", Value: "", Expected: false, ExpectedMsgsLen: 1},
		{Name: "/./_nil_false", PatternStr: ".", Value: nil, Expected: false, ExpectedMsgsLen: 1},
		{Name: "nil_nil_true", Value: nil, Expected: true, ExpectedMsgsLen: 0},
		{Name: "/^\\d+$/_99_true", PatternStr: "^\\d+$", Value: "99", Expected: true, ExpectedMsgsLen: 0},
		{Name: "/\\d/_99_true", PatternStr: "\\d", Value: "99", Expected: true, ExpectedMsgsLen: 0},
		{Name: "/\\d/_abc_false", PatternStr: "\\d", Value: "abc", Expected: false, ExpectedMsgsLen: 1},
		{Name: "/^[a-z]{5}$/_aeiou_true", PatternStr: "^[a-z]{5}$", Value: "aeiou", Expected: true, ExpectedMsgsLen: 0},
		{Name: "/^[a-z]{5}$/_aeiouy_false", PatternStr: "^[a-z]{5}$", Value: "aeiouy", Expected: false, ExpectedMsgsLen: 1},
	} {
		t.Run(tc.Name, func(t2 *testing.T) {
			vOptions := NewRegexValidatorOptions()
			if len(tc.PatternStr) > 0 {
				regex := regexp.MustCompile(tc.PatternStr)
				vOptions.Pattern = regex
			}
			validator := RegexValidator(vOptions)
			result, msgs := validator(tc.Value)
			msgsLen := len(msgs)
			if  result != tc.Expected {
				t2.Errorf("Expected %v for `result` boolean but got %v",
					tc.Expected, result)
			}
			if msgsLen != tc.ExpectedMsgsLen {
				t2.Errorf("Expected %d messages.  Got %d",
					tc.ExpectedMsgsLen, msgsLen)
			}
		})
	}
}