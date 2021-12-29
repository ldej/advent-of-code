package main

import (
	"fmt"
	"strings"

	"github.com/ldej/advent-of-code/tools"
	"github.com/ldej/advent-of-code/tools/myints"
	"github.com/ldej/advent-of-code/tools/mystrings"
)

func main() {
	fmt.Println("Part 2")

	result := run("./2020/day16/example2.txt")
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

	var yourTicket = myints.ParseCsv(strings.TrimLeft(lines[1], "your ticket:\n"))[0]
	var nearbyTickets = myints.ParseCsv(strings.TrimLeft(lines[2], "nearby tickets:\n"))

	validRulesForIndex, invalid := ValidRulesForTickets(nearbyTickets, rules)
	fmt.Println("Invalid count:", invalid)

	validRulesCounted := CountValidRules(validRulesForIndex)
	validRules := RulesThatApplyToAllTickets(validRulesCounted, len(validRulesForIndex))
	matchedRules := ReduceRules(validRules)

	var result = 1
	for i, v := range matchedRules {
		if strings.HasPrefix(v, "departure") {
			result = result * yourTicket[i]
		}
	}
	return result
}

func ValidRulesForTickets(nearbyTickets [][]int, rules []Rule) ([]map[int][]string, int) {
	var validRulesForTickets []map[int][]string
	var invalid int
	for _, ticket := range nearbyTickets {
		anyInvalid, invalidValue, validRulesMap := ValidRulesForIndex(ticket, rules)
		if !anyInvalid {
			validRulesForTickets = append(validRulesForTickets, validRulesMap)
		} else {
			invalid += invalidValue
		}
	}
	return validRulesForTickets, invalid
}

func ReduceRules(rules map[int][]string) map[int]string {
	var matches = make(map[int]string)
	for len(matches) < len(rules) {
		index, rule := SingleRule(rules)
		matches[index] = rule

		for i, name := range rules {
			rules[i] = mystrings.Filter(name, rule)
		}
	}
	return matches
}

func SingleRule(rules map[int][]string) (int, string) {
	for i, value := range rules {
		if len(value) == 1 {
			return i, value[0]
		}
	}
	return -1, ""
}

func RulesThatApplyToAllTickets(countedRules map[int]map[string]int, expected int) map[int][]string {
	var rulesPerIndex = make(map[int][]string)
	for i, v := range countedRules {
		for rule, count := range v {
			if count == expected {
				rulesPerIndex[i] = append(rulesPerIndex[i], rule)
			}
		}
	}
	return rulesPerIndex
}

func CountValidRules(validRulesForIndexes []map[int][]string) map[int]map[string]int {
	count := make(map[int]map[string]int)
	for _, validRulesForIndex := range validRulesForIndexes {
		for index, validRules := range validRulesForIndex {
			for _, rule := range validRules {
				if _, ok := count[index]; ok {
					count[index][rule] = count[index][rule] + 1
				} else {
					count[index] = map[string]int{rule: 1}
				}
			}
		}
	}
	return count
}

func ValidRulesForIndex(ticket []int, rules []Rule) (bool, int, map[int][]string) {
	var validRulesForIndex = make(map[int][]string)
	var invalidValue int
	var anyInvalid bool
	for i, value := range ticket {
		var valid bool
		for _, rule := range rules {
			for _, values := range rule.Values {
				if myints.InRange(value, values[0], values[1]) {
					validRulesForIndex[i] = append(validRulesForIndex[i], rule.Name)
					valid = true
				}
			}
		}
		if !valid {
			anyInvalid = true
			invalidValue += value
		}
	}
	return anyInvalid, invalidValue, validRulesForIndex
}
