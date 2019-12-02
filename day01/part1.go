package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func CalculateFuel(mass int) int {
	return (mass / 3) - 2
}

func main() {
	total := 0
	// Read input
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		mass, _ := strconv.Atoi(scanner.Text())
		total += CalculateFuel(mass)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	fmt.Println(total)
}
