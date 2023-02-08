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
	for _, c := range tests {
		got := greet(c.in)
		if got != c.want {
			t.Errorf("02-greet(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
