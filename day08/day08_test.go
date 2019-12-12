package main

import (
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {
	input := "123456789012"

	width := 3

	chunk := 0
	for i := 0; i < (len(input) / width); i++ {
		fmt.Println(input[chunk : chunk+width])
		chunk += width
	}
}
