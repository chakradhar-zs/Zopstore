package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		s    []int
		want []int
	}{
		{[]int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
	}
	for _, c := range tests {
		reverse(c.s)

		assert.Equal(t, c.want, c.s)
	}
}
