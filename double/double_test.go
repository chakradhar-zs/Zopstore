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
	for _, c := range tests {
		got := double(c.in)
		if got != c.want {
			t.Errorf("double(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}
