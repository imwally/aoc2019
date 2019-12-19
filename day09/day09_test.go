package main

import (
	"testing"

	"github.com/imwally/aoc2019/machine"
)

type TestCase struct {
	Program []int
	Result  []int
}

func TestPart1(t *testing.T) {
	cases := []TestCase{
		TestCase{
			Program: []int{1102, 34915192, 34915192, 7, 4, 7, 99, 0},
			Result:  []int{1219070632396864},
		},
		TestCase{
			Program: []int{104, 1125899906842624, 99},
			Result:  []int{1125899906842624},
		},
	}

	for _, tc := range cases {
		m := machine.New(tc.Program)
		m.MockInput([]int{1})
		m.SaveOutput()
		m.Run()

		got := m.Output
		expected := tc.Result[0]
		if got != expected {
			t.Errorf("got %v, expected %v", got, expected)
		}

	}

	program := []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}

	m := machine.New(program)
	m.Run()
}
