package main

import (
	"fmt"

	"github.com/imwally/aoc2019/machine"
)

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func Part1(program []int) int {
	perms := permutations([]int{0, 1, 2, 3, 4})

	output := 0
	max := 0
	for _, perm := range perms {
		for i := 0; i < len(perm); i++ {
			m := machine.New(program)
			m.SaveOutput()
			m.MockInput([]int{perm[i], output})
			m.RunFor(1)
			m.Run()
			output = m.Output
			if output > max {
				max = output
			}
		}
		output = 0
	}

	return max
}

func Part2(program []int) int {
	perms := permutations([]int{5, 6, 7, 8, 9})

	output := 0
	max := 0
	for _, perm := range perms {

		// Setup new machines
		machines := []*machine.Machine{
			machine.New(program),
			machine.New(program),
			machine.New(program),
			machine.New(program),
			machine.New(program),
		}

		// Initialize machines on first round
		for i := 0; i < len(machines); i++ {
			m := machines[i]
			m.SaveOutput()
			m.MockInput([]int{perm[i], output})
			m.RunFor(1)
			m.Run()
			output = m.Output
		}

		// Start feedback loop
		for i := 0; i < len(machines); i++ {
			m := machines[i]
			m.SaveOutput()
			m.MockInput([]int{output})
			m.RunFor(1)
			m.Run()
			output = m.Output
			if output > max {
				max = output
			}
			if i == 4 && m.Halted {
				break
			}
			if i == 4 {
				i = -1
			}
		}
		output = 0

	}

	return max
}

func main() {
	program := []int{3, 8, 1001, 8, 10, 8, 105, 1, 0, 0, 21, 46, 59, 80,
		105, 122, 203, 284, 365, 446, 99999, 3, 9, 102, 3, 9, 9, 1001, 9, 5, 9,
		102, 2, 9, 9, 1001, 9, 3, 9, 102, 4, 9, 9, 4, 9, 99, 3, 9, 1002, 9, 2,
		9, 101, 2, 9, 9, 4, 9, 99, 3, 9, 101, 5, 9, 9, 1002, 9, 3, 9, 1001, 9,
		3, 9, 1002, 9, 2, 9, 4, 9, 99, 3, 9, 1002, 9, 4, 9, 1001, 9, 2, 9, 102,
		4, 9, 9, 101, 3, 9, 9, 102, 2, 9, 9, 4, 9, 99, 3, 9, 102, 5, 9, 9, 101,
		4, 9, 9, 102, 3, 9, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101,
		1, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3,
		9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 2, 9,
		4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001,
		9, 1, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4,
		9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2,
		9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9,
		1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4,
		9, 99, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001,
		9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3,
		9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9,
		4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 99, 3, 9,
		102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4,
		9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2,
		9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9,
		1002, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 99, 3, 9, 102, 2, 9, 9,
		4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1002,
		9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3,
		9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001,
		9, 1, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 99}

	fmt.Println("Part 1:", Part1(program))
	fmt.Println("Part 2:", Part2(program))
}
