package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{1, 1},
		{2, 3},
		{3, 6},
		{4, 10},
	}
	for _, c := range tests {
		got := Sum(c.n)
		if got != c.want {
			t.Errorf("Sum(%d) == %d, want %d", c.n, got, c.want)
		}
	}
}
