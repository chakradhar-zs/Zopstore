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

	for _, tc := range tests {
		got := Sum(tc.n)
		if got != tc.want {
			t.Errorf("Sum(%d) == %d, want %d", tc.n, got, tc.want)
		}
	}
}
