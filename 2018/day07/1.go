package main

import (
	"fmt"
	"sort"
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

	//startMap := map[string]bool{}
	//for l, letters := range stepMap {
	//	if a, found := startMap[l]; !a || !found {
	//		startMap[l] = true
	//	}
	//	fmt.Println(l, letters)
	//	for _, letter := range letters {
	//		startMap[letter] = false
	//	}
	//}

	startLetters := []string{"Y", "O", "U", "X"}
	//for letter, found := range startMap {
	//	if found {
	//		startLetters = append(startLetters, letter)
	//	}
	//}

	fmt.Println(startLetters)

	result := []string{}
	lettersToDo := startLetters
	currentLetter := ""
	sort.Strings(lettersToDo)

	for len(lettersToDo) > 0 {
		currentLetter, lettersToDo = lettersToDo[0], lettersToDo[1:]
		result = append(result, currentLetter)

		for _, letter := range stepMap[currentLetter] {
			dropped, rest := stepMap[currentLetter][0], stepMap[currentLetter][1:]
			stepMap[currentLetter] = rest
			if hasNoIncomingEdges(dropped, stepMap) {
				lettersToDo = append(lettersToDo, letter)
				sort.Strings(lettersToDo)
			}
		}
	}
	fmt.Println(strings.Join(result, ""))
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
