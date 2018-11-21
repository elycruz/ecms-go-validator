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

func TestFloatWithinRange(t *testing.T) {
	type FloatTestCase struct {
		Name string
		Subj float64
		Min float64
		Max float64
		Expected bool
	}
	for _, tc := range []FloatTestCase{
		{Name: "0.0_within_range_-1_to_1", Subj: 0.0, Min: -1.0, Max: 1.0, Expected: true},
		{Name: "-0.0_within_range_-1_to_1", Subj: -0.0, Min: -1.0, Max: 1.0, Expected: true},
		{Name: "-0.0_within_range_0_to_1", Subj: -0.0, Min: 0.0, Max: 1.0, Expected: true},
		{Name: "99.0_within_range_89_to_99", Subj: 99.0, Min: 89.0, Max: 99.0, Expected: true},
		{Name: "99.0_not_within_range_-1_to_1", Subj: 99.0, Min: -1.0, Max: 1.0, Expected: false},
	}{
		t.Run(tc.Name, func(t2 *testing.T) {
			result := FloatWithinRange(tc.Min, tc.Max, tc.Subj)
			if result != tc.Expected {
				t2.Errorf("Expected %v; Got %v; For value %v", result, tc.Expected, tc.Subj)
			}
		})
	}
}
