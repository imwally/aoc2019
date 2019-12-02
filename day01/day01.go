package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Part 1
func CalculateFuel(mass int) int {
	return (mass / 3) - 2
}

// Part 2
func CalculateFuel2(mass, total int) int {
	if (mass/3)-2 < 0 {
		return total
	}

	calc := (mass / 3) - 2

	return CalculateFuel2(calc, calc+total)
}

func ReadInput(input *os.File) ([]int, error) {
	var modules []int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		module, _ := strconv.Atoi(scanner.Text())
		modules = append(modules, module)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return modules, nil
}

func main() {
	modules, err := ReadInput(os.Stdin)
	if err != nil {
		fmt.Println(err)
		return
	}

	part1Total := 0
	part2Total := 0
	for _, mass := range modules {
		part1Total += CalculateFuel(mass)
		part2Total += CalculateFuel2(mass, 0)
	}

	fmt.Println("Part 1:", part1Total)
	fmt.Println("Part 2:", part2Total)

}
