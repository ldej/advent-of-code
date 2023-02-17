package dijkstra

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDijkstra(t *testing.T) {
	g := NewGraph()

	g.AddUndirectedEdge("0", "1", 4)
	g.AddUndirectedEdge("0", "7", 8)
	g.AddUndirectedEdge("1", "7", 11)
	g.AddUndirectedEdge("1", "2", 8)
	g.AddUndirectedEdge("2", "3", 7)
	g.AddUndirectedEdge("2", "5", 4)
	g.AddUndirectedEdge("2", "8", 2)
	g.AddUndirectedEdge("3", "4", 9)
	g.AddUndirectedEdge("3", "5", 14)
	g.AddUndirectedEdge("4", "5", 10)
	g.AddUndirectedEdge("5", "6", 2)
	g.AddUndirectedEdge("6", "7", 1)
	g.AddUndirectedEdge("6", "8", 6)
	g.AddUndirectedEdge("7", "8", 7)

	type test struct {
		from          string
		to            string
		distances     []int
		totalDistance int
		path          string
	}
	tests := []test{
		{from: "0", to: "1", distances: []int{4}, totalDistance: 4, path: "0 -> 1"},
		{from: "0", to: "2", distances: []int{4, 8}, totalDistance: 12, path: "0 -> 1 -> 2"},
		{from: "0", to: "3", distances: []int{4, 8, 7}, totalDistance: 19, path: "0 -> 1 -> 2 -> 3"},
		{from: "0", to: "4", distances: []int{8, 1, 2, 10}, totalDistance: 21, path: "0 -> 7 -> 6 -> 5 -> 4"},
		{from: "0", to: "5", distances: []int{8, 1, 2}, totalDistance: 11, path: "0 -> 7 -> 6 -> 5"},
		{from: "0", to: "6", distances: []int{8, 1}, totalDistance: 9, path: "0 -> 7 -> 6"},
		{from: "0", to: "7", distances: []int{8}, totalDistance: 8, path: "0 -> 7"},
		{from: "0", to: "8", distances: []int{4, 8, 2}, totalDistance: 14, path: "0 -> 1 -> 2 -> 8"},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("From %s to %s", tc.from, tc.to), func(t *testing.T) {
			p := g.Shortest(tc.from, tc.to)
			require.NotNil(t, p)
			assert.Equal(t, tc.distances, p.distances)
			assert.Equal(t, tc.totalDistance, p.Distance())
			assert.Equal(t, tc.path, p.Print())
		})
	}
}

func TestDijkstraIntGrid(t *testing.T) {
	g := NewGraph()
	g.AddIntGrid([][]int{
		{1, 1, 6, 3, 7, 5, 1, 7, 4, 2},
		{1, 3, 8, 1, 3, 7, 3, 6, 7, 2},
		{2, 1, 3, 6, 5, 1, 1, 3, 2, 8},
		{3, 6, 9, 4, 9, 3, 1, 5, 6, 9},
		{7, 4, 6, 3, 4, 1, 7, 1, 1, 1},
		{1, 3, 1, 9, 1, 2, 8, 2, 3, 7},
		{1, 3, 5, 9, 9, 1, 2, 4, 2, 1},
		{3, 1, 2, 5, 4, 2, 1, 6, 3, 9},
		{1, 2, 9, 3, 1, 3, 8, 5, 2, 1},
		{2, 3, 1, 1, 9, 4, 4, 5, 8, 1},
	})
	p := g.ShortestIntGrid(0, 0, 9, 9)
	require.NotNil(t, p)
	assert.Equal(t, "x=0,y=0 -> x=0,y=1 -> x=0,y=2 -> x=1,y=2 -> x=2,y=2 -> x=3,y=2 -> x=4,y=2 -> x=5,y=2 -> x=6,y=2 -> x=6,y=3 -> x=7,y=3 -> x=7,y=4 -> x=8,y=4 -> x=8,y=5 -> x=8,y=6 -> x=8,y=7 -> x=8,y=8 -> x=9,y=8 -> x=9,y=9", p.Print())
	assert.Equal(t, 40, p.Distance())
}
