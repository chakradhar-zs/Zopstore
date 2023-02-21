package main

import "testing"

func TestFactorial(t *testing.T) {
	testCases := []struct {
		input    int
		expected int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
		{5, 120},
	}

	for _, tc := range testCases {
		f := factorial()
		result := f(tc.input)

		if result != tc.expected {
			t.Errorf("Factorial(%d) returned %d, expected %d", tc.input, result, tc.expected)
		}
	}
}
