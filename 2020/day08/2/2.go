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

type Instruction struct {
	Op    string
	Value int
}

func run() int {
	input := tools.ReadStrings("./2020/day08/input.txt")

	var instructions []Instruction

	for _, i := range input {
		values := strings.Split(i, " ")
		instructions = append(instructions, Instruction{Op: values[0], Value: tools.ToInt(values[1])})
	}

	for i := range instructions {
		updated := make([]Instruction, len(instructions), len(instructions))
		copy(updated, instructions)

		// Switch one statement
		if updated[i].Op == "jmp" {
			updated[i].Op = "nop"
		} else if updated[i].Op == "nop" {
			updated[i].Op = "jmp"
		} else {
			// nothing changed
			continue
		}

		accumulator, terminated := RunInstructions(updated)
		if terminated {
			return accumulator
		}
	}

	return -1
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
	}
}
