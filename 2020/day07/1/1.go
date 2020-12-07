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
	input := tools.ReadStrings("./2020/day07/input.txt")

	bags := BagsMap(input)

	canHoldShinyGold := 0
	for bag := range bags {
		if CanHoldShinyGold(bags, bag) {
			canHoldShinyGold += 1
		}
	}

	return canHoldShinyGold
}

func CanHoldShinyGold(bagsMap map[string][]string, bag string) bool {
	if tools.StringSliceContains(bagsMap[bag], "shiny gold") {
		return true
	}

	for _, b := range bagsMap[bag] {
		if CanHoldShinyGold(bagsMap, b) {
			return true
		}
	}
	return false
}

func BagsMap(input []string) map[string][]string {
	bagsMap := map[string][]string{}
	for _, line := range input {
		statements := strings.Split(line, " bags contain ")

		color, rest := statements[0], statements[1]

		bags := tools.RegexNamedGroupsRepeat(rest, `(?P<count>\d+) (?P<color>.*?) bag`)

		for _, bag := range bags {
			bagsMap[color] = append(bagsMap[color], bag["color"])
		}
	}
	return bagsMap
}
