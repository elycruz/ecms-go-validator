package ecms_validator

import (
	"fmt"
	"testing"
)

var (
	vowelsStr = "aeiou"
)

type ObscuratorTestCase struct {
	Name       string
	Control    string
	Expected   string
	ObscureLen int
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

func TestObscurateRight(t *testing.T) {

}