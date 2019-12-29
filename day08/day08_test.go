package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := "123456789012"

	testCases := []string{
		"123",
		"456",
		"789",
		"012",
	}

	width := 3
	chunkStart := 0

	for i := 0; i < (len(input) / width); i++ {
		got := input[chunkStart : chunkStart+width]
		chunkStart += width

		expected := testCases[i]
		if got != expected {
			t.Errorf("error: got %v, expectd %v", got, expected)
		}
	}
}
