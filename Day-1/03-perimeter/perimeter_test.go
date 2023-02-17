package main

import (
	"testing"
)

func TestCirclePerimeter(t *testing.T) {
	cases := []struct {
		radius float64
		want   float64
	}{
		{10, 62.83185307179586},
		{5, 31.41592653589793},
		{1, 6.283185307179586},
	}

	for _, tc := range cases {
		got := circlePerimeter(tc.radius)
		if got != tc.want {
			t.Errorf("circlePerimeter(%g) == %g, want %g", tc.radius, got, tc.want)
		}
	}
}

func TestSquarePerimeter(t *testing.T) {
	cases := []struct {
		side float64
		want float64
	}{
		{10, 40},
		{5, 20},
		{1, 4},
	}
	for _, tc := range cases {
		got := squarePerimeter(tc.side)
		if got != tc.want {
			t.Errorf("squarePerimeter(%g) == %g, want %g", tc.side, got, tc.want)
		}
	}
}
