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

	// Map opcodes to functions
	opCodes := map[int]func(int, int) int{
		1: Add,
		2: Multiply,
	}

	// Replace values before running
	ops[1] = 12
	ops[2] = 2

	// Run the machine
	for i := 0; ; i = i + 4 {
		op := ops[i]

		if op == 99 {
			break
		}

		x := ops[ops[i+1]]
		y := ops[ops[i+2]]
		pos := ops[ops[i+3]]

		ops[pos] = opCodes[op](x, y)
	}

	// For some reason my result is in position 1
	fmt.Println(ops[1])
}
