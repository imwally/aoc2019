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
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if s[i] > s[j] {
				return false
			}
		}
	}

	return true
}

func MoreThanDup(s string) bool {
	dups := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		dups[s[i]]++
	}

	for _, dup := range dups {
		if dup > max {
			return true
		}
	}

	return false
}

func main() {
	start := 197487
	end := 673251

	for i := start; i < end; i++ {
		num := strconv.Itoa(i)
		if Duplicates(num) && Increasing(num) && !MoreThanDup(num) {
			fmt.Println(num)
		}
	}

}
