package main

import (
	"testing"

	"github.com/imwally/aoc2019/machine"
)

type Codes struct {
	Input  []int
	Result []int
}

func TestIntcode(t *testing.T) {
	cases := []Codes{
		Codes{
			Input:  []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50},
			Result: []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50},
		},
		Codes{
			Input:  []int{1, 0, 0, 0, 99},
			Result: []int{2, 0, 0, 0, 99},
		},
		Codes{
			Input:  []int{2, 3, 0, 3, 99},
			Result: []int{2, 3, 0, 6, 99},
		},
		Codes{
			Input:  []int{2, 4, 4, 5, 99, 0},
			Result: []int{2, 4, 4, 5, 99, 9801},
		},
		Codes{
			Input:  []int{1, 1, 1, 4, 99, 5, 6, 0, 99},
			Result: []int{30, 1, 1, 4, 2, 5, 6, 0, 99},
		},
	}

	for _, c := range cases {
		got := machine.Run(c.Input)
		for i, v := range got {
			if v != c.Result[i] {
				t.Errorf("error: got %v, expected %v", v, c.Result[i])
			}
		}
	}
}
