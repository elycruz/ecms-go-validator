package ecms_validator

import (
	"testing"
)

type IntRangeValidatorTestCase struct {
	Name string
	TestValue interface {}
	ExpectedResultBln bool
	ExpectedMessagesLen int
	Min int64
	Max int64
}

type FloatRangeValidatorTestCase struct {
	Name string
	TestValue interface {}
	ExpectedResultBln bool
	ExpectedMessagesLen int
	Min float64
	Max float64
}

func TestIntRangeValidator(t *testing.T) {
	for _, testCase := range []IntRangeValidatorTestCase{
		{"validate_(nil)", nil, false, 1, 0, 0},
		{"validate_(0, 0, 0)", 0, true, 0, 0, 0},
		{"validate_(-10, 10, 0)", 0, true, 0, -10, 10},
		{"validate_(10, -10, 0)", 0, true, 0, 10, -10},
	} {
		t.Run(testCase.Name, func(t *testing.T) {
			validatorOptions := NewIntRangeValidatorOptions()
			validatorOptions.Min = testCase.Min
			validatorOptions.Max = testCase.Max
			validator := IntRangeValidator(validatorOptions)
			resultBln, messages := validator(testCase.TestValue)
			messagesLen := len(messages)
			if resultBln != testCase.ExpectedResultBln {
				t.Errorf("Expected %v for `resultBln.Result` but got %v",
					testCase.ExpectedResultBln, resultBln)
			}
			if messagesLen != testCase.ExpectedMessagesLen  {
				t.Errorf("Expected %d messages.  Got %d;  Messages: %v",
					testCase.ExpectedMessagesLen, messagesLen, messages)
			}
			for _, message := range messages {
				if len(message) == 0 {
					t.Error("Expected non-empty message strings.  " +
						"Received an empty message string.")
				}
			}
		})
	}
}

func TestFloatRangeValidator(t *testing.T) {
	for _, testCase := range []FloatRangeValidatorTestCase{
		{"float_validate_(nil)", nil, false, 1, 0, 0},
		{"float_validate_(0, 0, 0)", 0, true, 0, 0, 0},
		{"float_validate_(-10, 10, 0)", 0, true, 0, -10, 10},
		{"float_validate_(10, -10, 0)", 0, true, 0, 10, -10},
	} {
		t.Run(testCase.Name, func(t *testing.T) {
			validatorOptions := NewFloatRangeValidatorOptions()
			validatorOptions.Min = testCase.Min
			validatorOptions.Max = testCase.Max
			validator := FloatRangeValidator(validatorOptions)
			resultBln, messages := validator(testCase.TestValue)
			messagesLen := len(messages)
			if resultBln != testCase.ExpectedResultBln {
				t.Errorf("Expected %v for `resultBln.Result` but got %v",
					testCase.ExpectedResultBln, resultBln)
			}
			if messagesLen != testCase.ExpectedMessagesLen  {
				t.Errorf("Expected %d messages.  Got %d;  Messages: %v",
					testCase.ExpectedMessagesLen, messagesLen, messages)
			}
			for _, message := range messages {
				if len(message) == 0 {
					t.Error("Expected non-empty message strings.  " +
						"Received an empty message string.")
				}
			}
		})
	}
}
