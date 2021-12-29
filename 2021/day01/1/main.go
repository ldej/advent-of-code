package main

import (
	"fmt"

	"github.com/ldej/advent-of-code/tools"
)

func main() {
	result := run()
	fmt.Println(result)
}
func run() int {
	current, increased := 0, 0
	ints := tools.ReadInts()
	current, ints = ints[0], ints[1:]

	for _, i := range ints {
		if i > current {
			increased++
		}
		current = i
	}
	return increased
}
