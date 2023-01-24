package main

import (
	"fmt"

	"github.com/ldej/advent-of-code/tools"
	"github.com/ldej/advent-of-code/tools/mystrings"
)

func main() {
	result := run()
	fmt.Println(result)
}

func run() int {
	input := tools.ReadString()
	markerLength := 4
	for i := 0; i < len(input)-markerLength; i++ {
		part := input[i : i+markerLength]
		if part == mystrings.RemoveDuplicateLetters(part) {
			return i + markerLength
		}
	}
	return -1
}
