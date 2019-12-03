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

var distance float64 = 999999.00

func PrintGrid(g *grid) {
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g); j++ {
			fmt.Print(g[i][j], " ")
		}
		fmt.Println()
	}
}

func Trace(g *grid, x int, y int, steps int, dir string) (int, int) {
	for i := 0; i < steps; i++ {
		switch dir {
		case "L":
			y--
		case "R":
			y++
		case "U":
			x--
		case "D":
			x++
		}
		g[x][y]++
		if g[x][y] == 2 {
			d := math.Abs(float64(50000-x)) + math.Abs(float64(50000-y))
			if d < distance {
				distance = d
			}
		}
	}

	return x, y
}

func TraceWire(g *grid, path []string) {
	x, y := 50000, 50000
	for _, v := range path {
		direction := string(v[0])
		steps, err := strconv.Atoi(string(v[1:]))
		if err != nil {
			fmt.Println(err)
		}

		x, y = Trace(g, x, y, steps, direction)
	}
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

	fmt.Println(distance)

}
