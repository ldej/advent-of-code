package main

import (
	"fmt"
	"math"

	"github.com/ldej/advent-of-code/tools"
	"github.com/ldej/advent-of-code/tools/dijkstra/v1"
	"github.com/ldej/advent-of-code/tools/runegrid"
)

func main() {
	example1 := run("./example1.txt")
	fmt.Printf("\nExample 1:\n%v\n", example1)

	result := run()
	fmt.Printf("\nFinal:\n%v\n", result)
}

func run(file ...string) int {
	input := tools.ReadStrings(file...)
	graph := runegrid.FromStrings(input)
	start := graph.Find('S')
	end := graph.Find('E')
	graph.Set(start.Y, start.X, 'a')
	graph.Set(end.Y, end.X, 'z')

	d := dijkstra.NewGraph(graph.ToIntGrid())
	d.DistanceF = func(current *dijkstra.Node, neighbor *dijkstra.Node) int {
		if neighbor.Value-current.Value <= 1 {
			return current.Distance + 1
		}
		return math.MaxInt64
	}
	return d.Calculate(dijkstra.Coordinates{
		X: start.X,
		Y: start.Y,
	}, dijkstra.Coordinates{
		X: end.X,
		Y: end.Y,
	})
}
