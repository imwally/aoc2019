package main

import (
	"fmt"
	"testing"

	"github.com/imwally/aoc2019/machine"
)

func TestPart1(t *testing.T) {
	// Test 1
	program := []int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}
	seq := []int{4, 3, 2, 1, 0}
	for i := 0; i < len(seq); i++ {
		m1 := machine.New(program)
		m1.SaveOutput()
		m1.MockInput(seq[i])
		m1.Run()
		fmt.Println(m1.Output)
	}

	/*
		expected := 43210
		got := seq[len(seq)-1]

		if got != expected {
			t.Errorf("error: got %v, expected %v", got, expected)
		}
	*/
}

/*
func TestPart2(t *testing.T) {

	program := [][]int{
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

	seq := []int{5, 6, 7, 8, 9}
	output := make([]int, 10)

	for i := 0; i < len(seq); i++ {
		program[i], output = machine.Run(program[i], []int{seq[i], output[0], 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	}
	program[0], output = machine.Run(program[0], []int{seq[0], output[0], 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	fmt.Println(0, program[0])
	fmt.Println(output)
}
*/
