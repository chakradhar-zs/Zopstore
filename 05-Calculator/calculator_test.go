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

	for _, tc := range tests {
		got := calculator(tc.op, tc.a, tc.b)
		if got != tc.want {
			t.Errorf("calculator(%s, %f, %f) == %f, want %f", tc.op, tc.a, tc.b, got, tc.want)
		}
	}
}
