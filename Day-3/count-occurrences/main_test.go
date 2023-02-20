package main

import (
	"reflect"
	"testing"
)

func TestCountOccurrences(t *testing.T) {
	testCases := []struct {
		input    string
		expected map[string]int
	}{
		{
			input:    "Mississippi",
			expected: map[string]int{"M": 1, "i": 4, "s": 4, "p": 2},
		},
		{
			input:    "hello",
			expected: map[string]int{"h": 1, "e": 1, "l": 2, "o": 1},
		},
		{
			input:    "aaaaaa",
			expected: map[string]int{"a": 6},
		},
	}

	for _, tc := range testCases {
		result := countOccurrences(tc.input)
		
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("input: %s, expected: %v, but got: %v", tc.input, tc.expected, result)
		}
	}
}
