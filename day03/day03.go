package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type GridPoint struct {
	Intersection bool
	Wire         string
	Step         int
}

type grid [99999][99999]int

func PrintGrid(g *grid) {
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g); j++ {
			fmt.Print(g[i][j], " ")
		}
		fmt.Println()
	}
}

func TraceWire(g *grid, path []string) {
	distance := 999999.00
	x, y := 50000, 50000

	part2Counter := 0
	part2Max := 999999
	for _, v := range path {
		direction := string(v[0])
		steps, err := strconv.Atoi(string(v[1:]))
		if err != nil {
			fmt.Println(err)
		}

		if direction == "L" {
			for i := 0; i < steps; i++ {
				part2Counter++
				y--
				g[x][y]++
				if g[x][y] == 2 {
					d := math.Abs(float64(50000-x)) + math.Abs(float64(50000-y))
					if d < distance {
						distance = d
					}
					if part2Counter < part2Max {
						part2Max = part2Counter
					}
				}
			}
			x, y = x, y
		}

		if direction == "R" {
			for i := 0; i < steps; i++ {
				part2Counter++
				y++
				g[x][y]++
				if g[x][y] == 2 {
					d := math.Abs(float64(50000-x)) + math.Abs(float64(50000-y))
					if d < distance {
						distance = d
					}
					if part2Counter < part2Max {
						part2Max = part2Counter
					}
				}
			}
			x, y = x, y
		}

		if direction == "U" {
			for i := 0; i < steps; i++ {
				part2Counter++
				x--
				g[x][y]++
				if g[x][y] == 2 {
					d := math.Abs(float64(50000-x)) + math.Abs(float64(50000-y))
					if d < distance {
						distance = d
					}
					if part2Counter < part2Max {
						part2Max = part2Counter
					}
				}
			}
			x, y = x, y
		}

		if direction == "D" {
			for i := 0; i < steps; i++ {
				part2Counter++
				x++
				g[x][y]++
				if g[x][y] == 2 {
					d := math.Abs(float64(50000-x)) + math.Abs(float64(50000-y))
					if d < distance {
						distance = d
					}
					if part2Counter < part2Max {
						part2Max = part2Counter
					}
				}
			}
			x, y = x, y
		}
	}

	fmt.Println(distance)
	fmt.Println(part2Max)
}

func main() {
	var paths [][]string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		paths = append(paths, strings.Split(scanner.Text(), ","))
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	g := &grid{}

	/*
		TraceWire(g, []string{"R8", "U5", "L5", "D3"})
		TraceWire(g, []string{"U7", "R6", "D4", "L4"})
	*/
	TraceWire(g, paths[0])
	TraceWire(g, paths[1])

}
