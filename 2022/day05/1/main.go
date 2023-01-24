package main

import (
	"fmt"
	"strings"

	"github.com/ldej/advent-of-code/tools"
	"github.com/ldej/advent-of-code/tools/stack"
)

func main() {
	result := run()
	fmt.Println(result)
}

func run() string {
	input := tools.ReadStringsDoubleNewlines()
	stackLines, moves := input[0], input[1]

	lines := strings.Split(stackLines, "\n")

	var stacks []stack.RuneStack

	lastLine := lines[len(lines)-1]
	for i := range lastLine {
		if lastLine[i] != ' ' {
			s := stack.RuneStack{}
			for j := len(lines) - 2; j >= 0; j-- {
				if i < len(lines[j]) && rune(lines[j][i]) != ' ' {
					s.Push(rune(lines[j][i]))
				}
			}
			stacks = append(stacks, s)
		}
	}

	for _, line := range strings.Split(moves, "\n") {
		ints := tools.FindInts(line)
		count, from, to := ints[0], ints[1]-1, ints[2]-1

		for i := 0; i < count; i++ {
			box, _ := stacks[from].Pop()
			stacks[to].Push(box)
		}
	}

	var result string
	for _, s := range stacks {
		box, found := s.Peek()
		if found {
			result = result + string(box)
		}
	}

	return result
}
