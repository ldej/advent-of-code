package runegrid

import (
	"io/ioutil"
	"log"
	"strings"
)

func Read(location string) RuneGrid {
	var result RuneGrid
	content, err := ioutil.ReadFile(location)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		result = append(result, []rune(line))
	}
	return result
}

func FromStrings(input []string) RuneGrid {
	var newGrid = make(RuneGrid, 0)

	for _, row := range input {
		newGrid = append(newGrid, []rune(row))
	}
	return newGrid
}
