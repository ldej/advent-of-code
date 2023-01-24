package main

import (
	"fmt"

	"github.com/ldej/advent-of-code/tools"
)

func main() {
	example1 := run("./example1.txt")
	fmt.Printf("\nExample 1:\n%v\n", example1)

	example2 := run("./example2.txt")
	fmt.Printf("\nExample 2:\n%v\n", example2)

	result := run()
	fmt.Printf("\nFinal:\n%v\n", result)
}

func run(file ...string) int {
	input := tools.ReadStrings(file...)
	return len(input)
}
