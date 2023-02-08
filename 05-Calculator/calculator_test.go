package main

import (
	"testing"
)

func TestCalculator(t *testing.T) {
	tests := []struct {
		op   string
		a    float64
		b    float64
		want float64
	}{
		{"+", 2, 3, 5},
		{"-", 5, 3, 2},
		{"*", 2, 3, 6},
		{"/", 6, 3, 2},
	}
	for _, c := range tests {
		got := calculator(c.op, c.a, c.b)
		if got != c.want {
			t.Errorf("calculator(%s, %f, %f) == %f, want %f", c.op, c.a, c.b, got, c.want)
		}
	}
}
