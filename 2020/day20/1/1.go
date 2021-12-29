package main

import (
	"fmt"
	"strings"

	"github.com/ldej/advent-of-code/tools"
	"github.com/ldej/advent-of-code/tools/myrunes"
	"github.com/ldej/advent-of-code/tools/runegrid"
	"github.com/ldej/advent-of-code/tools/sets"
)

func main() {
	fmt.Println("Part 1")

	result := run("./2020/day20/example1.txt")
	fmt.Println("Example:", result)

	result = run("./2020/day20/input.txt")
	fmt.Println("Result:", result)
}

func run(input string) int {
	parts := tools.ReadStringsDoubleNewlines()

	var tiles = make(map[int]runegrid.RuneGrid)
	for _, part := range parts {
		lines := strings.Split(part, "\n")
		tiles[tools.FindInt(lines[0])] = runegrid.FromStrings(lines[1:])
	}

	orientations := []int{0, 90, 180, 270}
	flipped := []bool{false, true}

	var matches = make(map[int]sets.IntSet)

	for id1, tile1 := range tiles {
		for id2, tile2 := range tiles {
			if id1 != id2 {
				for _, flipped1 := range flipped {
					for _, orientation1 := range orientations {
						for _, orientation2 := range orientations {
							var check1 = tile1
							var check2 = tile2
							if flipped1 {
								check1 = check1.FlipHorizontal()
							}
							check1 = check1.Rotate(orientation1)
							check2 = check2.Rotate(orientation2)

							if myrunes.Equal(check1[0], check2[0]) {
								matchSet, found := matches[id1]
								if !found {
									matchSet = sets.NewIntSet()
								}
								matchSet.Add(id2)
								matches[id1] = matchSet
							}
						}
					}
				}
			}
		}
	}

	// corner pieces have only two tiles they match with
	var result = 1

	for id, matchSet := range matches {
		if matchSet.Len() == 2 {
			result *= id
		}
	}

	return result
}
