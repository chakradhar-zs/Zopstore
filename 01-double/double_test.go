package main

import "testing"

func TestDouble(t *testing.T) {
	tests := []struct {
		in, want int
	}{
		{2, 4},
		{-2, -4},
		{0, 0},
		{5, 10},
	}

	for _, tc := range tests {
		got := double(tc.in)
		if got != tc.want {
			t.Errorf("01-double(%d) == %d, want %d", tc.in, got, tc.want)
		}
	}
}
