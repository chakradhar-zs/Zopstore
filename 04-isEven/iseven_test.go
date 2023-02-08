package main

import (
	"testing"
)

func TestIsEven(t *testing.T) {
	tests := []struct {
		in   int
		want bool
	}{
		{2, true},
		{3, false},
		{0, true},
	}
	for _, c := range tests {
		got := isEven(c.in)
		if got != c.want {
			t.Errorf("isEven(%d) == %t, want %t", c.in, got, c.want)
		}
	}
}
