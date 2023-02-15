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

	for _, tc := range tests {
		got := isEven(tc.in)
		if got != tc.want {
			t.Errorf("isEven(%d) == %t, want %t", tc.in, got, tc.want)
		}
	}
}
