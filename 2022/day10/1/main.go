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
	input := tools.ReadStrings()

	cycle := 0
	register := 1

	var value int
	for _, line := range input {
		if strings.HasPrefix(line, "noop") {
			cycle += 1
			if (cycle+20)%40 == 0 {
				value += cycle * register
			}
			continue
		}
		cycle += 1
		if (cycle+20)%40 == 0 {
			value += cycle * register
		}
		cycle += 1
		if (cycle+20)%40 == 0 {
			value += cycle * register
		}
		number := tools.FindInt(line)
		register += number
	}

	return value
}
