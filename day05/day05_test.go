package main

import (
	"testing"

	"github.com/imwally/aoc2019/machine"
)

func TestInput(t *testing.T) {
	program := []int{1002, 4, 3, 4, 33}
	expected := []int{1002, 4, 3, 4, 99}

	m := machine.New(program)
	m.Run()

	got := m.DumpMemory()

	for i, _ := range expected {
		if got[i] != expected[i] {
			t.Errorf("error: got %v, expected %v", got, expected)
		}
	}
}

func TestOutput(t *testing.T) {
	program2 := []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20,
		1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1,
		46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98,
		99}

	m2 := machine.New(program2)
	m2.MockInput([]int{8})
	m2.Run()
}
