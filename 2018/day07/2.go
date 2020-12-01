package main

import (
	"sort"

	"fmt"
	"strings"

	"github.com/ldej/advent-of-code/2018/common"
)

func main() {
	results := common.ReadAllLines("./day7/input.txt", `Step (?P<incoming>.*) must be finished before step (?P<from>.*) can begin.`)

	stepMap := map[string][]string{}
	for _, result := range results {
		if result["from"] != "" {
			if _, found := stepMap[result["incoming"]]; found {
				letters := append(stepMap[result["incoming"]], result["from"])
				sort.Strings(letters)
				stepMap[result["incoming"]] = letters
			} else {
				stepMap[result["incoming"]] = []string{result["from"]}
			}
		}
	}

	startLetters := []string{"Y", "O", "U", "X"}

	result := []string{}
	lettersToDo := startLetters
	currentLetter := ""
	sort.Strings(lettersToDo)

	workers := []int{0, 0, 0, 0, 0}
	workersletters := []string{"", "", "", "", ""}
	totalTime := 0

	lettermap := map[string]int{
		"A": 61,
		"B": 62,
		"C": 63,
		"D": 64,
		"E": 65,
		"F": 66,
		"G": 67,
		"H": 68,
		"I": 69,
		"J": 70,
		"K": 71,
		"L": 72,
		"M": 73,
		"N": 74,
		"O": 75,
		"P": 76,
		"Q": 77,
		"R": 78,
		"S": 79,
		"T": 80,
		"U": 81,
		"V": 82,
		"W": 83,
		"X": 84,
		"Y": 85,
		"Z": 86,
	}

	for len(lettersToDo) > 0 || workersletters[0] != "" || workersletters[1] != "" || workersletters[2] != "" || workersletters[3] != "" || workersletters[4] != "" {
		for idx := range workers {
			if workers[idx] <= 0 {
				if workersletters[idx] == "" {
					if len(lettersToDo) > 0 {
						currentLetter, lettersToDo = lettersToDo[0], lettersToDo[1:]
						workers[idx] = lettermap[currentLetter]
						workersletters[idx] = currentLetter
					}
				} else {
					currentLetter = workersletters[idx]
					result = append(result, currentLetter)

					for _, letter := range stepMap[currentLetter] {
						dropped, rest := stepMap[currentLetter][0], stepMap[currentLetter][1:]
						stepMap[currentLetter] = rest
						if hasNoIncomingEdges(dropped, stepMap) {
							lettersToDo = append(lettersToDo, letter)
							sort.Strings(lettersToDo)
						}
					}
					workersletters[idx] = ""
				}
			}
		}
		totalTime += 1
		for idx := range workers {
			workers[idx] -= 1
		}
	}

	// TODO wrong
	fmt.Println(strings.Join(result, ""))
	fmt.Println(totalTime)
}

func hasNoIncomingEdges(letter string, stepMap map[string][]string) bool {
	for _, letters := range stepMap {
		for _, l := range letters {
			if l == letter {
				return false
			}
		}
	}
	return true
}
