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

func run() string {
	input := tools.ReadStrings()

	cycle := 0
	register := 1

	var b strings.Builder
	for _, line := range input {
		if strings.HasPrefix(line, "noop") {
			cycle = step(cycle, register, &b)
			continue
		}

		cycle = step(cycle, register, &b)
		cycle = step(cycle, register, &b)
		number := tools.FindInt(line)
		register += number
	}

	return b.String()
}

func step(cycle int, register int, b *strings.Builder) int {
	if cycle >= register-1 && cycle <= register+1 {
		b.WriteString("#")
	} else {
		b.WriteString(".")
	}
	cycle = (cycle + 1) % 40
	if cycle == 0 {
		b.WriteString("\n")
	}
	return cycle
}
