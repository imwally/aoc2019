package main

import (
	"testing"

	"github.com/imwally/aoc2019/machine"
)

func TestSave(t *testing.T) {
	program := []int{1002, 4, 3, 4, 33}
	expected := []int{1002, 4, 3, 4, 99}

	got := machine.Run(program)

	for i, _ := range expected {
		if got[i] != expected[i] {
			t.Errorf("error: got %v, expected %v", got, expected)
		}
	}
}
