package main

import (
	"reflect"
	"testing"
)

func TestSumValuesByKey(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string][]int
		expected map[string]int
	}{
		{
			"empty map", map[string][]int{}, map[string]int{},
		},
		{
			"map with one key-value pair", map[string][]int{"foo": {1, 2, 3}}, map[string]int{"foo": 6},
		},
		{
			"map with multiple key-value pairs", map[string][]int{"foo": {1, 2, 3}, "bar": {4, 5}, "baz": {}},
			map[string]int{"foo": 6, "bar": 9, "baz": 0},
		},
	}

	for _, test := range tests {
		got := sumValuesByKey(test.input)

		if !reflect.DeepEqual(got, test.expected) {
			t.Errorf("unexpected result for %s:\nexpected: %v\n     got: %v", test.name, test.expected, got)
		}
	}
}
