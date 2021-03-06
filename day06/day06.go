package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Object struct {
	Name    string
	Parents []string
}

type ObjectMap struct {
	Objects map[string][]*Object
}

type ParentMap map[string][]string

func AllOrbits(pm ParentMap) int {
	i := 0
	for _, v := range pm {
		i += len(v)
	}

	return i
}

func OrbitsFromTo(p ParentMap, a, b string) int {
	for i := 0; i < len(p[a]) && i < len(p[b]); i++ {
		if p[a][i+1] != p[b][i+1] {
			x := len(p[a][i+1:]) - 1
			y := len(p[b][i+1:]) - 1
			return x + y
		}
	}

	return 0
}

func GetParents(om ObjectMap, current string, children []string, parents ParentMap) {
	children = append(children, current)
	for _, child := range om.Objects[current] {
		child.Parents = children
		GetParents(om, child.Name, children, parents)
	}
	parents[current] = children
}

func BuildMaps(objects []string) (ObjectMap, ParentMap) {
	om := ObjectMap{}
	newMap := make(map[string][]*Object)
	om.Objects = newMap
	for _, object := range objects {
		objs := strings.Split(object, ")")
		parent, child := objs[0], objs[1]
		childObject := &Object{Name: child}
		om.Objects[parent] = append(om.Objects[parent], childObject)
	}

	var child []string
	parents := make(map[string][]string)
	GetParents(om, "COM", child, parents)

	return om, parents
}

func main() {
	var objects []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		objects = append(objects, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	_, pm := BuildMaps(objects)
	// Part 1
	fmt.Println("Part 1:", AllOrbits(pm)-len(objects)-1)

	// Part 2
	fmt.Println("Part 2:", OrbitsFromTo(pm, "YOU", "SAN"))
}
