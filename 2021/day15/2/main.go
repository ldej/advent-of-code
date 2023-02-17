package main

import (
	"fmt"

	"github.com/ldej/advent-of-code/tools"
)

func main() {
	result := run()
	fmt.Println(result)
}

func run() int {
	cave := tools.ReadIntGrid()
	factor := 5

	bigCave := tools.NewIntGrid(len(cave)*factor, len(cave[0])*5, 0)
	for i := 0; i < factor; i++ {
		for j := 0; j < factor; j++ {
			for r, row := range cave {
				for c, cell := range row {
					v := (cell + i + j) % 9
					if v == 0 {
						v = 9
					}
					bigCave[r+(i*len(cave))][c+(j*len(cave[0]))] = v
				}
			}
		}
	}

	graph := dijkstra.NewGraph(bigCave)
	return graph.Calculate(graph.Min, graph.Max)
}
