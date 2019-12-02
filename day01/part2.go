package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func CalculateFuel(mass, total int) int {
	if (mass/3)-2 < 0 {
		return total
	}

	calc := (mass / 3) - 2

	return CalculateFuel(calc, calc+total)
}

func main() {
	total := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		mass, _ := strconv.Atoi(scanner.Text())
		total += CalculateFuel(mass, 0)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	fmt.Println(total)
}
