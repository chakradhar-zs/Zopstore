package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateMapValue(t *testing.T) {
	tests := []struct {
	  name     string
	  inputMap map[string]int
	  key      string
	  value    int
	  expected map[string]int
	}{
	  {
		name: "update existing key",
		inputMap: map[string]int{"a": 1,"b": 2,"c": 3},key: "b",value: 4,
		expected: map[string]int{"a": 1,"b": 4,"c": 3},
	  },
	  {
		name: "add new key",inputMap: map[string]int{"a": 1,"b": 2,"c": 3},key: "d",value: 5,
		expected: map[string]int{"a": 1,"b": 2,"c": 3,"d": 5},
	  },
	}
  
	for _, tc := range tests {
		updateMapValue(tc.inputMap, tc.key, tc.value)

		assert.Equal(t,tc.expected,tc.inputMap)
	}
  }
  