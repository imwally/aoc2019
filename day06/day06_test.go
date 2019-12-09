package main

import "testing"

func TestPart1(t *testing.T) {
	objects := []string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L"}
	_, pm := BuildMaps(objects)

	expected := 42
	got := AllOrbits(pm) - len(objects) - 1
	if got != expected {
		t.Errorf("error: got %v, expected %v", got, expected)
	}

}

func TestPart2(t *testing.T) {
	orbitMap := []string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L", "K)YOU", "I)SAN"}
	_, pm := BuildMaps(orbitMap)

	expected := 4
	got := OrbitsFromTo(pm, "YOU", "SAN")
	if got != expected {
		t.Errorf("error: got %v, expected %v", got, expected)
	}

}
