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

	counts := polymerise(polymer, rulesMap, 40, map[string]map[string]int{})
	for key, value := range mystrings.CountLetters(polymer) {
		counts[key] += value
	}

	min, max := myints.MinAndMax(tools.MapIntValues(counts))
	return max - min
}

func polymerise(polymer string, rules map[string]string, steps int, cache map[string]map[string]int) map[string]int {
	cacheKey := fmt.Sprint(polymer, steps)
	if result, found := cache[cacheKey]; found {
		return result
	}
	counts := map[string]int{}
	if steps == 0 {
		return counts
	}

	for i := 0; i < len(polymer)-1; i++ {
		pair := polymer[i : i+2]
		letter := rules[pair]
		counts[letter]++

		for key, value := range polymerise(pair[:1]+letter+pair[1:], rules, steps-1, cache) {
			counts[key] += value
		}
	}

	cache[cacheKey] = counts
	return counts
}
