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

	for i := 0; ; i = i + 4 {
		op := ops[i]

		x := ops[ops[i+1]]
		y := ops[ops[i+2]]
		pos := ops[ops[i+3]]

		fmt.Println(op, x, y, pos)
		if op == 99 {
			break
		}
		val := operations[op](x, y)

		ops[pos] = val

		fmt.Println(ops)
	}

}
