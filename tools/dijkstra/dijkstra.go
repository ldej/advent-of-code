// Package dijkstra implements Dijkstra's shortest path algorithm
// See: https://www.geeksforgeeks.org/dijkstras-shortest-path-algorithm-greedy-algo-7/
// It uses a priority queue to determine which unvisited vertex has the minimum distance
// See: https://pkg.go.dev/container/heap
package dijkstra

import (
	"container/heap"
	"math"
)

type Graph struct {
	Nodes         [][]*Node
	Min           Coordinates
	Max           Coordinates
	PriorityQueue *nodePriorityQueue
	DistanceF     func(current *Node, neighbor *Node) int
}

func NewGraph(input [][]int) Graph {
	graph := Graph{
		Min:           Coordinates{X: 0, Y: 0},
		Max:           Coordinates{X: len(input[0]) - 1, Y: len(input) - 1},
		PriorityQueue: &nodePriorityQueue{},
		DistanceF: func(current *Node, neighbor *Node) int {
			return current.Distance + neighbor.Value
		},
	}
	heap.Init(graph.PriorityQueue)
	for j := 0; j < len(input); j++ {
		var row []*Node
		for i := 0; i < len(input[j]); i++ {
			row = append(row, &Node{
				Position: Coordinates{
					X: i,
					Y: j,
				},
				Value:    input[j][i],
				Distance: math.MaxInt64,
				Visited:  false,
			})
		}
		graph.Nodes = append(graph.Nodes, row)
	}
	return graph
}

type Coordinates struct {
	X, Y int
}

type Node struct {
	Position Coordinates
	Value    int
	Distance int
	Visited  bool
}

func (g *Graph) unvisitedNeighbors(n *Node) []*Node {
	var neighbors []*Node
	if n.Position.X > 0 && !g.Nodes[n.Position.Y][n.Position.X-1].Visited {
		neighbors = append(neighbors, g.Nodes[n.Position.Y][n.Position.X-1])
	}
	if n.Position.X < g.Max.X && !g.Nodes[n.Position.Y][n.Position.X+1].Visited {
		neighbors = append(neighbors, g.Nodes[n.Position.Y][n.Position.X+1])
	}
	if n.Position.Y > 0 && !g.Nodes[n.Position.Y-1][n.Position.X].Visited {
		neighbors = append(neighbors, g.Nodes[n.Position.Y-1][n.Position.X])
	}
	if n.Position.Y < g.Max.Y && !g.Nodes[n.Position.Y+1][n.Position.X].Visited {
		neighbors = append(neighbors, g.Nodes[n.Position.Y+1][n.Position.X])
	}
	return neighbors
}

func (g *Graph) visit(n *Node) {
	neighbors := g.unvisitedNeighbors(n)
	for _, neighbor := range neighbors {
		distance := g.DistanceF(n, neighbor)
		if distance < neighbor.Distance {
			neighbor.Distance = distance
			// Only unvisited neighbors for which the distance has changed are put in the priority queue
			heap.Push(g.PriorityQueue, neighbor)
		}
	}
	n.Visited = true
}

func (g *Graph) currentNode() *Node {
	// The heap stores which nodes are eligible to be selected as the next current Node
	// The nodes are stored in a priority queue, meaning the Node with the shortest tentative distance
	// will always be at the front of the queue.
	return heap.Pop(g.PriorityQueue).(*Node)
}

func (g *Graph) Calculate(start Coordinates, end Coordinates) int {
	g.Nodes[start.Y][start.X].Distance = 0
	heap.Push(g.PriorityQueue, g.Nodes[start.Y][start.X])

	for g.PriorityQueue.Len() > 0 {
		current := g.currentNode()
		if current.Position == end {
			return current.Distance
		}
		g.visit(current)
	}
	return -1
}

type nodePriorityQueue []*Node

func (h nodePriorityQueue) Len() int           { return len(h) }
func (h nodePriorityQueue) Less(i, j int) bool { return h[i].Distance < h[j].Distance }
func (h nodePriorityQueue) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *nodePriorityQueue) Push(x interface{}) {
	*h = append(*h, x.(*Node))
}

func (h *nodePriorityQueue) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	old[n-1] = nil
	*h = old[0 : n-1]
	return x
}
