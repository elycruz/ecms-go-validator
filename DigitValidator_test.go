package ecms_validator

import (
	"testing"
)

type DigitValidatorTestCase struct {
	Name 			string
	Value           interface{}
	Expected        bool
	ExpectedMsgsLen int
}

func TestDigitValidator(t *testing.T) {
	for _, tc := range []DigitValidatorTestCase{
		{Name: "//_99_true", Value: 99, Expected: true, ExpectedMsgsLen: 0},
		{Name: "//_\"99\"_true", Value: "99", Expected: true, ExpectedMsgsLen: 0},
		{Name: "/./_\"hello\"_false", Value: "hello", Expected: false, ExpectedMsgsLen: 1},
		{Name: "/./_\"\"_false", Value: "", Expected: false, ExpectedMsgsLen: 1},
		{Name: "/./_nil_false", Value: nil, Expected: false, ExpectedMsgsLen: 1},
	} {
		t.Run(tc.Name, func(t2 *testing.T) {
			validator := DigitValidator()
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