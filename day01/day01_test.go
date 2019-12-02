package main

import (
	"testing"
)

func TestCalculateFuel(t *testing.T) {
	cases := map[int]int{
		12:     2,
		14:     2,
		1969:   654,
		100756: 33583,
	}

	for mass, expected := range cases {
		got := CalculateFuel(mass)
		if got != expected {
			t.Errorf("error: got %v, expected %v", got, expected)
		}
	}
}

func TestCalculateFuel2(t *testing.T) {
	cases := map[int]int{
		14:     2,
		1969:   966,
		100756: 50346,
	}

	for mass, expected := range cases {
		got := CalculateFuel2(mass, 0)
		if got != expected {
			t.Errorf("error: got %v, expected %v", got, expected)
		}
	}
}
