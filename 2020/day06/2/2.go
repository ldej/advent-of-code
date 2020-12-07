package main

import (
	"fmt"
	"strings"

	"github.com/ldej/advent-of-code/tools"
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
		total += len(tools.StringsUnion(people))
	}

	return total
}
