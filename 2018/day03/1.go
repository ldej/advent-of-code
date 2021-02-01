package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	content, _ := ioutil.ReadFile("./day3/input.txt")
	lines := strings.Split(string(content), "\n")

	fabricMap := make(map[int]map[int][]int)
	compRegEx := regexp.MustCompile(`^#(?P<Id>\d+) @ (?P<Left>\d+),(?P<Top>\d+): (?P<Width>\d+)x(?P<Height>\d+)$`)

	for _, line := range lines {
		// #1 @ 342,645: 25x20
		match := compRegEx.FindStringSubmatch(line)

		resultsMap := make(map[string]int)
		for i, name := range compRegEx.SubexpNames() {
			if i > 0 && i <= len(match) {
				resultsMap[name], _ = strconv.Atoi(match[i])
			}
		}

		// Fill the fabric map
		for i := resultsMap["Left"]; i < resultsMap["Left"]+resultsMap["Width"]; i++ {
			for j := resultsMap["Top"]; j < resultsMap["Top"]+resultsMap["Height"]; j++ {
				if column, foundCol := fabricMap[i]; foundCol {
					if _, foundRow := column[j]; foundRow {
						fabricMap[i][j] = append(fabricMap[i][j], resultsMap["Id"])
					} else {
						fabricMap[i][j] = []int{resultsMap["Id"]}
					}
				} else {
					fabricMap[i] = map[int][]int{j: {resultsMap["Id"]}}
				}
			}
		}
	}

	// Count overlapping
	overlapping := 0
	for key := range fabricMap {
		for _, ids := range fabricMap[key] {
			if len(ids) > 1 {
				overlapping += 1
			}
		}
	}
	fmt.Println(overlapping)
}
