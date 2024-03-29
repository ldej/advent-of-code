package main

import (
	"fmt"
	"strings"

	"github.com/ldej/advent-of-code/tools"
	"github.com/ldej/advent-of-code/tools/myints"
)

func main() {
	fmt.Println("Part 1")

	result := run("./2020/day16/example1.txt")
	fmt.Println("Example:", result)

	result = run("./2020/day16/input.txt")
	fmt.Println("Result:", result)
}

type Rule struct {
	Name   string
	Values [][]int
}

func run(input string) int {
	lines := tools.ReadStringsDoubleNewlines()

	var rules []Rule
	runeLines := strings.Split(lines[0], "\n")
	for _, r := range runeLines {
		result := tools.RegexNamedGroups(r, `(?P<name>.*?): (?P<low1>\d+)-(?P<high1>\d+) or (?P<low2>\d+)-(?P<high2>\d+)`)
		rules = append(rules, Rule{
			Name: result["name"],
			Values: [][]int{
				{myints.ToInt(result["low1"]), myints.ToInt(result["high1"])},
				{myints.ToInt(result["low2"]), myints.ToInt(result["high2"])},
			},
		})
	}

	//yourTicket := tools.ParseCsv(lines[1])

	var nearbyTickets = myints.ParseCsv(strings.TrimLeft(lines[2], "nearby tickets:\n"))

	invalid := 0
	for _, ticket := range nearbyTickets {
		invalid += SumInvalidValues(ticket, rules)
	}

	return invalid
}

func SumInvalidValues(ticket []int, rules []Rule) int {
	sum := 0
	for _, value := range ticket {
		if !ValidValue(value, rules) {
			sum += value
		}
	}
	return sum
}

func ValidValue(value int, rules []Rule) bool {
	for _, rule := range rules {
		for _, values := range rule.Values {
			if myints.InRange(value, values[0], values[1]) {
				return true
			}
		}
	}
	return false
}
