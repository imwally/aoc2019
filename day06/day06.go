package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func PrintOrbits(orbitMap map[string][]string, current string, i int, j int) int {
	for _, v := range orbitMap[current] {
		j = PrintOrbits(orbitMap, v, i+1, j+i)
	}

	return j
}

func TotalOrbits(om []string) int {
	orbits := make(map[string][]string)

	for _, v := range om {
		objs := strings.Split(v, ")")
		orbits[objs[0]] = append(orbits[objs[0]], objs[1])
	}

	return PrintOrbits(orbits, "COM", 0, 0) + len(om)
}

func main() {
	var om []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		om = append(om, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Part 1:", TotalOrbits(om))
}
