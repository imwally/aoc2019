package main

import (
	"testing"
)

func TestParseOrbits(t *testing.T) {
	orbitMap := []string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L"}
	expected := 42
	got := TotalOrbits(orbitMap)
	if got != expected {
		t.Errorf("error: got %v, expected %v", got, expected)
	}
}
