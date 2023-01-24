package main

import (
	"fmt"
	"strings"

	"github.com/ldej/advent-of-code/tools"
	"github.com/ldej/advent-of-code/tools/myints"
)

func main() {
	result := run()
	fmt.Println(result)
}

func run() int {
	input := tools.ReadStrings()

	var count int
	for _, line := range input {
		parts := strings.Split(line, ",")
		x := myints.ToInts(strings.Split(parts[0], "-"))
		y := myints.ToInts(strings.Split(parts[1], "-"))

		if (x[0] >= y[0] && x[1] <= y[1]) || (y[0] >= x[0] && y[1] <= x[1]) {
			count++
		}
	}
	return count
}
