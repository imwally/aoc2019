package main

import (
	"testing"

	"github.com/imwally/aoc2019/machine"
)

func TestPart1(t *testing.T) {
	// Test 1
	program := []int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15,
		99, 0, 0}

	seq := []int{4, 3, 2, 1, 0}
	output := 0
	for i := 0; i < len(seq); i++ {
		m := machine.New(program)
		m.SaveOutput()
		m.MockInput([]int{seq[i], output})
		m.Run()
		output = m.Output
	}

	expected := 43210
	got := output

	if got != expected {
		t.Errorf("error: got %v, expected %v", got, expected)
	}

	program2 := []int{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23,
		101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0}
	seq2 := []int{0, 1, 2, 3, 4}

	output2 := 0
	for i := 0; i < len(seq); i++ {
		m := machine.New(program2)
		m.SaveOutput()
		m.MockInput([]int{seq2[i], output2})
		m.Run()
		output2 = m.Output
	}

	expected = 54321
	got = output2

	if got != expected {
		t.Errorf("error: got %v, expected %v", got, expected)
	}

}

func TestPart2(t *testing.T) {

	programs := [][]int{
		[]int{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1,
			27, 26, 27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5},
		[]int{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1,
			27, 26, 27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5},
		[]int{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1,
			27, 26, 27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5},
		[]int{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1,
			27, 26, 27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5},
		[]int{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1,
			27, 26, 27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5},
	}

	machines := []*machine.Machine{
		machine.New(programs[0]),
		machine.New(programs[1]),
		machine.New(programs[2]),
		machine.New(programs[3]),
		machine.New(programs[4]),
	}

	seq := []int{5, 6, 7, 8, 9}
	output := 0

	for i := 0; ; i++ {
		m := machines[i]
		m.SaveOutput()
		m.MockInput([]int{seq[i], output, 0})
		m.RunFor(1)
		m.Run()
		output = m.Output
		if i == 4 && m.Halted {
			return
		}

		if i == len(seq)-1 {
			i = -1
		}
		m.IP = 0
	}

	got := output
	expected := 61696857

	if got != expected {
		t.Errorf("error: got %v, expected %v", got, expected)
	}
}
