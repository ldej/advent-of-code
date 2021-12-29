package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/ldej/advent-of-code/tools"
	"github.com/ldej/advent-of-code/tools/myints"
	"github.com/ldej/advent-of-code/tools/mystrings"
)

func main() {
	result := run()
	fmt.Println(result)
}

func run() int {
	input := tools.ReadStrings()

	sum := 0

	for _, line := range input {
		parts := strings.Split(line, "|")
		output := strings.Fields(parts[1])

		lengths := mapLengths(strings.Fields(line))

		outputNumber := ""

		for _, part := range output {
			switch len(part) {
			case 2:
				outputNumber += "1"
			case 3:
				outputNumber += "7"
			case 4:
				outputNumber += "4"
			case 5:
				// either 2, 3, 5
				// if 1 (= 2 segments) is part of, it's number 3
				if commonSegmentCount(lengths[2][0], part) == 2 {
					outputNumber += "3"
					// if 4 (= 4 segments) has 2 overlapping segments, it's 2
				} else if commonSegmentCount(lengths[4][0], part) == 2 {
					outputNumber += "2"
					// if 4 (= 4 segments has 3 overlapping segments, it's 5
				} else if commonSegmentCount(lengths[4][0], part) == 3 {
					outputNumber += "5"
				} else {
					log.Fatal("Whoops")
				}
			case 6:
				// either 0, 6, 9
				// if 4 (=4 segments) is part of it, it's 9
				if commonSegmentCount(lengths[4][0], part) == 4 {
					outputNumber += "9"
					// if 7 (=3 segments) has 2 overlapping parts, it's 6
				} else if commonSegmentCount(lengths[3][0], part) == 2 {
					outputNumber += "6"
					// if 7 (=3 segments) has 3 overlapping parts, it's 0
				} else if commonSegmentCount(lengths[3][0], part) == 3 {
					outputNumber += "0"
				} else {
					log.Fatal("Whoops")
				}
			case 7:
				outputNumber += "8"
			}
		}

		sum += myints.ToInt(outputNumber)
	}
	return sum
}

func mapLengths(input []string) map[int][]string {
	m := map[int][]string{}
	for _, line := range input {
		m[len(line)] = append(m[len(line)], line)
	}
	return m
}

func commonSegmentCount(a string, b string) int {
	return len(mystrings.Intersection([]string{a, b}))
}
