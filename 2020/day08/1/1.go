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

type Instruction struct {
	Op    string
	Value int
}

func run() int {
	input := tools.ReadStrings()

	var instructions []Instruction

	for _, i := range input {
		values := strings.Split(i, " ")
		instructions = append(instructions, Instruction{Op: values[0], Value: myints.ToInt(values[1])})
	}

	acc, _ := RunInstructions(instructions)

	return acc
}

// RunInstructions runs the instructions and returns the
// accumulator and if it was terminated. If it wasn't terminated
// then it has an infinite loop.
func RunInstructions(instructions []Instruction) (int, bool) {
	var accumulator = 0
	var visited = map[int]bool{}
	var index = 0

	for {
		// Detect out-of-bounds
		if index < 0 || index >= len(instructions) {
			return accumulator, true
		}

		// Detect loop
		if visited[index] {
			return accumulator, false
		} else {
			visited[index] = true
		}

		switch instructions[index].Op {
		case "nop":
			index += 1
		case "acc":
			accumulator += instructions[index].Value
			index += 1
		case "jmp":
			index += instructions[index].Value
		}

		if index >= len(instructions) {
			return accumulator, true
		}
	}
}
