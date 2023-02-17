package main

import "testing"

func TestGreet(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{"John", "Hello, John!"},
		{"Jane", "Hello, Jane!"},
		{"Bob", "Hello, Bob!"},
	}

	for _, tc := range tests {
		got := greet(tc.in)
		if got != tc.want {
			t.Errorf("02-greet(%q) == %q, want %q", tc.in, got, tc.want)
		}
	}
}
