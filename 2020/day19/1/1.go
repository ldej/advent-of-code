package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ldej/advent-of-code/tools"
	"github.com/ldej/advent-of-code/tools/myints"
)

func main() {
	fmt.Println("Part 1")

	result := run("./2020/day19/example1.txt")
	fmt.Println("Example:", result)

	result = run("./2020/day19/input.txt")
	fmt.Println("Result:", result)
}

func run(input string) int {
	lines := tools.ReadStringsDoubleNewlines(input)
	ruleLines := strings.Split(lines[0], "\n")
	messages := strings.Split(lines[1], "\n")

	rulesMap := map[int]string{}
	for _, ruleLine := range ruleLines {
		parts := strings.Split(ruleLine, ": ")
		rulesMap[myints.ToInt(parts[0])] = parts[1]
	}

	regex := createRegex(rulesMap[0], rulesMap, make(map[string]string, 0))

	var correct int
	for _, message := range messages {
		if match, _ := regexp.MatchString(fmt.Sprintf("^%s$", regex), message); match {
			correct++
		}
	}
	return correct
}

func createRegex(value string, rules map[int]string, cache map[string]string) string {
	value = strings.Trim(value, "\"")

	if val, found := cache[value]; found {
		return val
	}

	if value == "a" || value == "b" {
		cache[value] = value
		return value
	} else if strings.Contains(value, " | ") {
		options := strings.Split(value, " | ")
		var regexes []string
		for _, option := range options {
			regexes = append(regexes, createRegex(option, rules, cache))
		}
		result := fmt.Sprintf("(%s)", strings.Join(regexes, "|"))
		cache[value] = result
		return result
	} else {
		numbers := strings.Split(value, " ")
		var results []string
		for _, number := range numbers {
			results = append(results, createRegex(rules[myints.ToInt(number)], rules, cache))
		}
		result := strings.Join(results, "")
		cache[value] = result
		return result
	}
}
