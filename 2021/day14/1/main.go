package main

import (
	"fmt"
	"strings"

	"github.com/ldej/advent-of-code/tools"
	"github.com/ldej/advent-of-code/tools/myints"
	"github.com/ldej/advent-of-code/tools/mystrings"
)

func main() {
	result := run()
	fmt.Println(result)
}
func run() int {
	input := tools.ReadStrings()
	polymer, rules := input[0], input[2:]

	rulesMap := map[string]string{}
	for _, rule := range rules {
		parts := strings.Split(rule, " -> ")
		rulesMap[parts[0]] = parts[1]
	}

	for step := 0; step < 10; step++ {
		polymerBuilder := strings.Builder{}
		polymerBuilder.WriteString(polymer[0:1])

		for i := 0; i < len(polymer)-1; i++ {
			polymerBuilder.WriteString(rulesMap[polymer[i:i+2]])
			polymerBuilder.WriteString(polymer[i+1 : i+2])
		}
		polymer = polymerBuilder.String()
	}
	min, max := myints.MinAndMax(tools.MapIntValues(mystrings.CountLetters(polymer)))
	return max - min
}
