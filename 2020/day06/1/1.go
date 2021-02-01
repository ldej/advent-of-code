package main

import (
	"fmt"
	"strings"

	"github.com/ldej/advent-of-code/tools"
	"github.com/ldej/advent-of-code/tools/mystrings"
)

func main() {
	result := run()
	fmt.Println(result)
}

func run() int {
	lines := tools.ReadStringsDoubleNewlines("./2020/day06/input.txt")

	total := 0

	for _, line := range lines {
		people := strings.Split(line, "\n")
		total += len(mystrings.Intersection(people))
	}

	return total
}
