package ecms_validator

import (
	"fmt"
	"testing"
)

var (
	vowelsStr = "aeiou"
	vowelsLen = len(vowelsStr)
)

type ObscuratorTestCase struct {
	Name       string
	Control    string
	Expected   string
	ObscureLen int
}

type GetErrorMessageByKeyTestCase struct {
	Name     string
	Options  ValidatorOptions
	Key      int
	Value    interface{}
	Expected string
}

func TestGetErrorMessageByKey(t *testing.T) {
	sharedOptions := NewNotEmptyValidatorOptions()

	// Length validator options
	intRangeValidatorOptions := NewIntRangeValidatorOptions()
	intRangeValidatorOptions.Min = 2
	intRangeValidatorOptions.Max = 5

	for _, tc := range []GetErrorMessageByKeyTestCase{
		{"(sharedOptions, EmptyNotAllowed, \"\")", sharedOptions, EmptyNotAllowed, "", DefaultEmptyNotAllowedMsg},
		{"(sharedOptions, EmptyNotAllowed, 0)", sharedOptions, EmptyNotAllowed, 0, DefaultEmptyNotAllowedMsg},
		{"(sharedOptions, EmptyNotAllowed, nil)", sharedOptions, EmptyNotAllowed, nil, DefaultEmptyNotAllowedMsg},
		{"(sharedOptions, NotAValidType, nil)", intRangeValidatorOptions, NotAValidType, nil,
			"<nil> is not a validatable numeric type."},
		{"(sharedOptions, NotAValidType, 99)", intRangeValidatorOptions, NotAValidType, 99,
			"99 is not a validatable numeric type."},
		{"GetErrorMessageByKey(sharedOptions, NotWithinRange, nil)", intRangeValidatorOptions, NotWithinRange, "aeiouy",
			"aeiouy is not within range 2 and 5."},
	} {
		t.Run(tc.Name, func(t2 *testing.T) {
			result := GetErrorMessageByKey(tc.Options, tc.Key, tc.Value)
			if result != tc.Expected {
				t2.Errorf("Expected \n\"%v\"\ngot\n\"%v\"\n", tc.Expected, result)
			}
		})
	}

	t.Run("Not found key", func(t2 *testing.T) {
		tc := GetErrorMessageByKeyTestCase{
			Name:     "GetErrorMessageByKey(sharedOptions, NotAValidType, nil) ",
			Options:  sharedOptions,
			Key:      NotAValidType,
			Value:    nil,
			Expected: "No error message found.",
		}
		rslt := GetErrorMessageByKey(tc.Options, tc.Key, tc.Value)
		if rslt != tc.Expected {
			t2.Errorf("Expected \n\"%v\"\ngot\n\"%v\"\n", tc.Expected, rslt)
		}
	})

}

func TestDefaultValueObscurator(t *testing.T) {
	for _, tc := range append(
		[]ObscuratorTestCase{
			{
				Name:       "DefaultObscurator(0,\"\") === \"\"",
				Control:    "",
				Expected:   "",
				ObscureLen: 0,
			},
		},
		func() []ObscuratorTestCase {
			xs := make([]ObscuratorTestCase, 0)
			for i, _ := range vowelsStr {
				obscured := ""
				for j := 0; j < i; j += 1 {
					obscured += "*"
				}
				expected := obscured + vowelsStr[i:]
				xs = append(xs, ObscuratorTestCase{
					Name:       fmt.Sprintf("DefaultObscurator(%v,\"%v\") === \"%v\"", i, vowelsStr, expected),
					Control:    vowelsStr,
					Expected:   expected,
					ObscureLen: i,
				})
			}
			return xs
		}()...
	) {
		t.Run(tc.Name, func(t2 *testing.T) {
			r := DefaultValueObscurator(tc.ObscureLen, tc.Control)
			if r != tc.Expected {
				t2.Errorf("expected `%v === %v`", r, tc.Expected)
			}
		})
	}
}

func TestObscurateLeft(t *testing.T) {
	for _, tc := range append(
		[]ObscuratorTestCase{
			{
				Name:       "ObscurateLeft(0,\"\") === \"\"",
				Control:    "",
				Expected:   "",
				ObscureLen: 0,
			},
		},
		func() []ObscuratorTestCase {
			xs := make([]ObscuratorTestCase, 0)
			for i, _ := range vowelsStr {
				obscured := ""
				for j := 0; j < i; j += 1 {
					obscured += "*"
				}
				expected := obscured + vowelsStr[i:]
				xs = append(xs, ObscuratorTestCase{
					Name:       fmt.Sprintf("ObscurateLeft(%v,\"%v\") === \"%v\"", i, vowelsStr, expected),
					Control:    vowelsStr,
					Expected:   expected,
					ObscureLen: i,
				})
			}
			return xs
		}()...
	) {
		t.Run(tc.Name, func(t2 *testing.T) {
			r := DefaultValueObscurator(tc.ObscureLen, tc.Control)
			if r != tc.Expected {
				t2.Errorf("expected `%v === %v`", r, tc.Expected)
			}
		})
	}
}

func TestObscurateRight(t *testing.T) {
	for _, tc := range append(
		[]ObscuratorTestCase{
			{
				Name:       "ObscurateRight(0,\"\") === \"\"",
				Control:    "",
				Expected:   "",
				ObscureLen: 0,
			},
		},
		func() []ObscuratorTestCase {
			xs := make([]ObscuratorTestCase, 0)
			for i, _ := range vowelsStr {
				obscured := ""
				for j := 0; j < i; j += 1 {
					obscured += "*"
				}
				expected := vowelsStr[0:vowelsLen-i] + obscured
				xs = append(xs, ObscuratorTestCase{
					Name:       fmt.Sprintf("ObscurateRight(%v,\"%v\") === \"%v\"", i, vowelsStr, expected),
					Control:    vowelsStr,
					Expected:   expected,
					ObscureLen: i,
				})
			}
			return xs
		}()...
	) {
		t.Run(tc.Name, func(t2 *testing.T) {
			r := ObscurateRight(tc.ObscureLen, tc.Control)
			if r != tc.Expected {
				t2.Errorf("expected `%v === %v`", r, tc.Expected)
			}
		})
	}
}
