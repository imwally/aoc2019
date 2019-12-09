package main

import (
	"testing"

	"github.com/imwally/aoc2019/machine"
)

func TestPart1(t *testing.T) {

	// Test 1
	program := []int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}
	seq := []int{4, 0, 3, 0, 2, 0, 1, 0, 0, 0, 0, 0}
	for i := 0; i < len(seq)-2; i++ {
		_, output := machine.Run(program, seq[i:])
		seq[i+3] = output[0]
		i++
	}

	expected := 43210
	got := seq[len(seq)-1]

	if got != expected {
		t.Errorf("error: got %v, expected %v", got, expected)
	}

	// Test 2
	program = []int{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23,
		101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0}
	seq = []int{0, 0, 1, 0, 2, 0, 3, 0, 4, 0, 0, 0}
	for i := 0; i < len(seq)-2; i++ {
		_, output := machine.Run(program, seq[i:])
		seq[i+3] = output[0]
		i++
	}

	expected = 54321
	got = seq[len(seq)-1]

	if got != expected {
		t.Errorf("error: got %v, expected %v", got, expected)
	}

	// Test 3
	program = []int{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33,
		1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0}
	seq = []int{1, 0, 0, 0, 4, 0, 3, 0, 2, 0, 0, 0}
	for i := 0; i < len(seq)-2; i++ {
		_, output := machine.Run(program, seq[i:])
		seq[i+3] = output[0]
		i++
	}

	expected = 65210
	got = seq[len(seq)-1]

	if got != expected {
		t.Errorf("error: got %v, expected %v", got, expected)
	}

}
