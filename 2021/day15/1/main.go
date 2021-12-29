package main

import (
	"fmt"

	"github.com/ldej/advent-of-code/tools"
	"github.com/ldej/advent-of-code/tools/dijkstra"
)

func main() {
	result := run()
	fmt.Println(result)
}
func run() int {
	cave := tools.ReadIntGrid()
	graph := dijkstra.NewGraph(cave)
	return graph.Calculate(graph.Min, graph.Max)
}
