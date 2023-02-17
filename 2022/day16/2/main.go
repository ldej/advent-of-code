package main

import (
	"fmt"
	"strings"

	"github.com/ldej/advent-of-code/tools"
	"github.com/ldej/advent-of-code/tools/dijkstra/v2"
	"github.com/ldej/advent-of-code/tools/myints"
	"github.com/ldej/advent-of-code/tools/mystrings"
)

func main() {
	example1 := run("./example1.txt")
	fmt.Printf("\nExample 1:\n%v\n", example1)

	result := run()
	fmt.Printf("\nFinal:\n%v\n", result)
}

func run(file ...string) int {
	input := tools.ReadStrings(file...)

	flowRates := map[string]int{}
	valveTunnels := map[string][]string{}
	nameMap := map[string]int{}
	var destinations []string

	for i, line := range input {
		parts := strings.Split(line, ";")

		valveInput := strings.Split(
			strings.ReplaceAll(
				strings.ReplaceAll(parts[0], "Valve ", ""),
				" has flow rate", ""),
			"=",
		)

		tunnels := strings.Split(
			strings.ReplaceAll(
				strings.ReplaceAll(parts[1], " tunnels lead to valves ", ""),
				" tunnel leads to valve ", ""),
			", ",
		)

		f := myints.ToInt(valveInput[1])
		name := valveInput[0]
		flowRates[name] = f
		valveTunnels[name] = tunnels
		nameMap[name] = i
		if f > 0 {
			destinations = append(destinations, name)
		}
	}

	g := dijkstra.NewGraph()
	for valve, tunnels := range valveTunnels {
		for _, tunnel := range tunnels {
			g.AddUndirectedEdge(valve, tunnel, 1)
		}
	}

	distanceMatrix := g.AllPaths()
	var max int
	for i := 1; i < len(destinations)/2+1; i++ {
		for newDestinations := range tools.CombinationsStr(destinations, i) {
			maxA := calculatePressure(distanceMatrix, flowRates, 0, 0, 0, "AA", newDestinations, 26)
			maxB := calculatePressure(distanceMatrix, flowRates, 0, 0, 0, "AA", mystrings.Complement(destinations, newDestinations), 26)
			if maxA+maxB > max {
				max = maxA + maxB
			}
		}
	}
	return max
}

func calculatePressure(matrix map[string]map[string]*dijkstra.Path, flowRates map[string]int, time int, pressure int, flow int, currentValve string, remainingValves []string, maxTime int) int {
	currentMaxPressure := pressure + (maxTime-time)*flow
	max := currentMaxPressure

	for _, nextValve := range remainingValves {
		timeToReachAndOpen := matrix[currentValve][nextValve].Distance() + 1
		if time+timeToReachAndOpen < maxTime {
			maxPressure := calculatePressure(
				matrix,
				flowRates,
				time+timeToReachAndOpen,
				pressure+timeToReachAndOpen*flow,
				flow+flowRates[nextValve],
				nextValve,
				mystrings.Filter(remainingValves, nextValve),
				maxTime,
			)
			if maxPressure > max {
				max = maxPressure
			}
		}
	}
	return max
}
