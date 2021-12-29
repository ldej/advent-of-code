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

type Bag struct {
	Color string
	Count int
}

func run() int {
	input := tools.ReadStrings()

	bags := BagsMap(input)

	count := Count(bags, Bag{"shiny gold", 1})

	return count
}

func Count(all map[string][]Bag, bag Bag) int {
	var contains int
	for _, b := range all[bag.Color] {
		contains += b.Count + b.Count*Count(all, b)
	}
	return contains
}

func BagsMap(input []string) map[string][]Bag {
	bagsMap := map[string][]Bag{}
	for _, line := range input {
		statements := strings.Split(line, " bags contain ")

		color, rest := statements[0], statements[1]

		bags := tools.RegexNamedGroupsRepeat(rest, `(?P<count>\d+) (?P<color>.*?) bag`)

		for _, bag := range bags {
			bagsMap[color] = append(
				bagsMap[color],
				Bag{
					Color: bag["color"],
					Count: myints.ToInt(bag["count"]),
				},
			)
		}
	}
	return bagsMap
}
