package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const gridSize = 17800
const gridStart = gridSize / 2

var distance float64 = 999999.00
var totalSteps int = 999999
var stepsA int = 0
var stepsB int = 0

type grid [gridSize][gridSize]*GridPoint

type GridPoint struct {
	Intersection bool
	StepsA       int
	StepsB       int
}

func PrintGrid(g *grid) {
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g); j++ {
			fmt.Print(g[i][j], " ")
		}
		fmt.Println()
	}
}

func Trace(g *grid, x int, y int, steps int, dir string, wire string) (int, int) {
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

		if wire == "a" {
			stepsA++
		} else {
			stepsB++
		}

		gp := g[x][y]
		if gp == nil {
			gp = &GridPoint{}
		}
		gp.StepsA += stepsA
		gp.StepsB += stepsB

		g[x][y] = gp

		if gp.StepsA > 0 && gp.StepsB > 0 {
			gp.Intersection = true

			bothSteps := gp.StepsA + gp.StepsB
			if bothSteps < totalSteps {
				totalSteps = bothSteps
			}

			d := math.Abs(float64(gridStart-x)) + math.Abs(float64(gridStart-y))
			if d < distance {
				distance = d
			}
		}
	}

	return x, y
}

func TraceWire(g *grid, path []string, wire string) {
	stepsA, stepsB = 0, 0
	x, y := gridStart, gridStart
	for _, v := range path {
		direction := string(v[0])
		steps, err := strconv.Atoi(string(v[1:]))
		if err != nil {
			fmt.Println(err)
		}

		x, y = Trace(g, x, y, steps, direction, wire)
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

	TraceWire(g, paths[0], "a")
	TraceWire(g, paths[1], "b")

	fmt.Println("Part 1:", distance)
	fmt.Println("Part 2:", totalSteps)
}
