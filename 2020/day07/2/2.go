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

type Bag struct {
	Color string
	Count int
}

func run() int {
	input := tools.ReadStrings("./2020/day07/input.txt")

	bags := BagsMap(input)

	count := Count(bags, Bag{"shiny gold", 1})

	return count
}

func Count(all map[string][]Bag, bag Bag) int {
	contains := 0
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

		bags := strings.Split(strings.TrimSuffix(rest, "."), ", ")

		for _, b := range bags {
			if b == "no other bags" {
				bagsMap[color] = []Bag{}
			} else {
				res := strings.SplitN(strings.Trim(b, " "), " ", 2)
				r := strings.TrimSuffix(strings.TrimSuffix(res[1], " bags"), " bag")
				bagsMap[color] = append(bagsMap[color], Bag{r, tools.ToInt(res[0])})
			}
		}
	}
	return bagsMap
}
