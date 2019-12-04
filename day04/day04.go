package main

import (
	"fmt"
	"strconv"
)

func Duplicates(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			return true
		}
	}

	return false
}

func Increasing(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i+1] < s[i] {
			return false
		}
	}

	return true
}

func LargerGroup(s string) bool {
	dups := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		dups[s[i]]++
	}

	for _, v := range dups {
		if v == 2 {
			return false
		}
	}

	return true
}

func main() {
	start := 197487
	end := 673251

	// PART 1
	var part1 []string
	for i := start; i < end; i++ {
		num := strconv.Itoa(i)
		if Duplicates(num) && Increasing(num) {
			part1 = append(part1, num)
		}
	}
	fmt.Println("Part 1:", len(part1))

	// PART 2
	var part2 []string
	for i := start; i < end; i++ {
		num := strconv.Itoa(i)
		if Duplicates(num) && Increasing(num) && !LargerGroup(num) {
			part2 = append(part2, num)
		}
	}
	fmt.Println("Part 2:", len(part2))
}
