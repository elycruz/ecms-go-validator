package ecms_validator

import "testing"

type LengthValidatorTestCase struct {
	Name string
	TestValue interface {}
	ExpectedResultBln bool
	ExpectedMessagesLen int
	Min int64
	Max int64
}

func init () {

}

func TestLengthValidator(t *testing.T) {
	for _, testCase := range []LengthValidatorTestCase{
		{
			Name:                "nil, false",
			TestValue:           nil,
			ExpectedResultBln:   false,
			ExpectedMessagesLen: 1,
			Min:                 1,
			Max:                 3,
		},
		{"\"\", false", "", false, 1, 1, 3},
		{"[]byte{}, false", make([]byte, 0), false, 1, 1, 3},
		{"[]byte{'a', 'b'}, true", []byte{'a', 'b'}, true, 0, 1, 3},
		{"[2]byte{}, true", [2]byte{}, true, 0, 1, 3},
		{"[2]byte{'a', 'b'}, true", [2]byte{'a', 'b'}, true, 0, 1, 3},
		{"map[int]int{}, false", make(map[int]int, 0), false, 1, 1, 3},
		{"map[int]int{1,2}, true", map[int]int{1: 2}, true, 0, 1, 3},
	} {
		t.Run(testCase.Name, func(t *testing.T) {
			validatorOptions := NewLengthValidatorOptions()
			validatorOptions.Min = testCase.Min
			validatorOptions.Max = testCase.Max
			validator := LengthValidator(validatorOptions)
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
