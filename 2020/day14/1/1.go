package main

import (
	"fmt"
	"strings"

	"github.com/ldej/advent-of-code/tools"
)

func main() {
	fmt.Println("Part 1")

	result := run("./2020/day14/example1.txt")
	fmt.Println("Example:", result)

	result = run("./2020/day14/input.txt")
	fmt.Println("Result:", result)
}

func run(input string) int {
	lines := tools.ReadStrings(input)

	var memory = make(map[int]int)
	var mask string

	for _, line := range lines {
		instruction := strings.Split(line, " = ")
		if instruction[0] == "mask" {
			mask = instruction[1]
		} else {
			mem, value := tools.FindInt(instruction[0]), tools.ToInt(instruction[1])
			memory[mem] = applyMask(value, mask)
		}
	}

	return tools.MapSumValues(memory)
}

func applyMask(value int, mask string) int {
	for i, bit := range mask {
		pos := len(mask) - 1 - i
		switch bit {
		case '0':
			value = tools.SetBit(value, 0, pos)
		case '1':
			value = tools.SetBit(value, 1, pos)
		}
	}
	return value
}
