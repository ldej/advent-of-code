package main

import (
	"github.com/ldej/advent-of-code/2018/common"
	"strings"
	"fmt"
	)

func main() {
	results := common.ReadAllLines("./day12/input.txt", `(?P<data>.*)`)

	initialState := strings.Replace(results[0]["data"], "initial state: ", "", 1)

	rules := map[string]string{}
	for idx := 2; idx < len(results); idx++ {
		rule := results[idx]["data"]
		if rule != "" {
			data := strings.Split(rule, " => ")
			if data[1] == "#" {
				rules[data[0]] = data[1]
			}
		}
	}

	newGeneration, offset := doGenerations(initialState, rules,500)

	total := count(newGeneration, offset)

	// Add 8 zeroes somewhere inbetween
	fmt.Println(total)
}

func count(input string, offset int) int {
	total := 0
	for i := 0; i < len(input); i++ {
		if input[i:i+1] == "#" {
			total += i - offset
		}
	}
	return total
}

func doGenerations(generation string, rules map[string]string, nrOfGenerations int) (string, int) {
	newGeneration := ""
	newOffset := 0
	currentGeneration, offset := checkBorders(generation)
	for i := 0; i < nrOfGenerations; i++ {
		newGeneration, newOffset = doGeneration(currentGeneration, rules)
		offset += newOffset
		currentGeneration = newGeneration
		//fmt.Printf("%d: %s\n", i+1, currentGeneration)
	}
	return newGeneration, offset
}

func checkBorders(generation string) (string, int) {
	offset := 0
	result := generation
	if generation[0:1] == "#" {
		result = "..." + generation
		offset += 3
	} else if generation[0:2] == ".#" {
		result = ".." + generation
		offset += 2
	} else if generation[0:3] == "..#" {
		result = "." + generation
		offset += 1
	} else if result[0:4] == "...." {
		result = result[1:]
		offset -= 1
	}

	if generation[len(generation)-1:] == "#" {
		result = result + "..."
	} else if generation[len(generation)-2:] == "#." {
		result = result + ".."
	} else if generation[len(generation)-3:] == "#.." {
		result = result + "."
	}
	return result, offset
}

func doGeneration(generation string, rules map[string]string) (string, int) {
	newGenaration := ""
	for i := 0; i < len(generation); i++ {
		rule := generation[common.Max(i-2, 0):common.Min(i+3, len(generation))]
		if plant, ok := rules[rule]; ok {
			newGenaration = newGenaration + plant
		} else {
			newGenaration = newGenaration + "."
		}
	}
	return checkBorders(newGenaration)
}
