package main

import (
	"fmt"
	"strings"

	"github.com/ldej/advent-of-code/tools"
)

func main() {
	result := run()
	fmt.Println(result)
}

type Cave struct {
	Name        string
	Connections []*Cave
}

func (c *Cave) IsSmall() bool {
	return strings.ToLower(c.Name) == c.Name
}

func (c *Cave) AddConnection(cave *Cave) {
	c.Connections = append(c.Connections, cave)
}

var caveMap = map[string]*Cave{}

func run() int {
	input := tools.ReadStrings()
	for _, line := range input {
		parts := strings.Split(line, "-")
		from, to := parts[0], parts[1]
		if _, found := caveMap[from]; !found {
			caveMap[from] = &Cave{Name: from}
		}
		if _, found := caveMap[to]; !found {
			caveMap[to] = &Cave{Name: to}
		}
		caveMap[from].AddConnection(caveMap[to])
		caveMap[to].AddConnection(caveMap[from])
	}
	return len(visit(caveMap["start"], []*Cave{}, false))
}

func visit(current *Cave, currentPath []*Cave, revisited bool) [][]*Cave {
	if current.Name == "end" {
		currentPath = append(currentPath, current)
		return [][]*Cave{currentPath}
	}
	if current.IsSmall() {
		for _, cave := range currentPath {
			if current.Name == cave.Name {
				if current.Name == "start" {
					return [][]*Cave{}
				}
				if revisited {
					return [][]*Cave{}
				}
				revisited = true
			}
		}
	}
	var paths [][]*Cave
	currentPath = append(currentPath, current)
	for _, option := range current.Connections {
		newPath := append([]*Cave{}, currentPath...)
		paths = append(paths, visit(option, newPath, revisited)...)
	}
	return paths
}
