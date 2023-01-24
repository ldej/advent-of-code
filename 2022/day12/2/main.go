package main

import (
	"fmt"
	"math"

	"github.com/ldej/advent-of-code/tools"
	"github.com/ldej/advent-of-code/tools/dijkstra"
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

	shortest := math.MaxInt64

	for cell := range graph.Cells() {
		if cell.Value != 'a' {
			continue
		}
		d := dijkstra.NewGraph(graph.ToIntGrid())
		d.DistanceF = func(current *dijkstra.Node, neighbor *dijkstra.Node) int {
			if neighbor.Value-current.Value <= 1 {
				return current.Distance + 1
			}
			return math.MaxInt64
		}
		steps := d.Calculate(dijkstra.Coordinates{
			X: cell.X,
			Y: cell.Y,
		}, dijkstra.Coordinates{
			X: end.X,
			Y: end.Y,
		})
		if steps > 0 && steps < shortest {
			shortest = steps
		}
	}
	return shortest
}
