package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Add(x, y int) int {
	return x + y
}

func Multiply(x, y int) int {
	return x * y
}

func main() {
	var ops []int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		for _, op := range strings.Split(scanner.Text(), ",") {
			opInt, _ := strconv.Atoi(op)
			ops = append(ops, opInt)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	operations := map[int]func(int, int) int{
		1: Add,
		2: Multiply,
	}

	origOps := make([]int, len(ops))
	copy(origOps, ops)

	for j := 0; j < 100; j++ {
		for k := 0; k < 100; k++ {
			copy(ops, origOps)
			ops[1], ops[2] = j, k
			for i := 0; ; i = i + 4 {
				op := ops[i]

				if op == 99 || op > 2 || ops[i+1] > len(ops) || ops[i+2] > len(ops) || ops[i+3] > len(ops) {
					break
				}

				x := ops[ops[i+1]]
				y := ops[ops[i+2]]
				pos := ops[ops[i+3]]

				if pos > len(ops) {
					break
				}

				val := operations[op](x, y)

				if val == 19690720 {
					fmt.Println("found")
					fmt.Println(ops[1], ops[2])
					return
				}

				ops[pos] = val

			}
		}
	}

}
