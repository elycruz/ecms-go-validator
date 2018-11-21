package is

import "testing"

func TestEmpty(t *testing.T) {
	type TestCaseEmpty struct {
		Name string
		Subj interface{}
		Expected bool
	}
	for _, tc := range []TestCaseEmpty{
		// True cases
		{"nil->true", nil, true},
		{"false->true", false, true},
		{"empty-string->true", "", true},
		{"empty-slice->true", make([]string, 0), true},
		{"empty-map->true", make(map[string]interface{}), true},
		{"empty-struct->true", struct{}{}, true},
		{"0.0->true", 0.0, true},
		{"0->true", 0, true},
		{"complex64(0)->true", complex64(0), true},
		{"complex128(0)->true", complex128(0), true},

		// False cases
		{"struct{Name:\"abc\"}->false", struct{Name string}{"abc"}, false},
		{"true->false", true, false},
		{"complex64(1)->false", complex64(1), false},
		{"complex128(1)->false", complex128(1), false},
		{"non-empty-string->false", "hello", false},
		{"non-empty-slice->false", []string{"hello"}, false},
		{"non-empty-map->false", map[string]int{"a": 0}, false},
		{"99.0->false", 99.0, false},
		{"1->false", 1, false},
	}{
		t.Run(tc.Name, func(t2 *testing.T) {
			result := Empty(tc.Subj)
			if result != tc.Expected {
				t2.Errorf("Expected %v; Got %v; For value %v", result, tc.Expected, tc.Subj)
			}
		})
	}
}
