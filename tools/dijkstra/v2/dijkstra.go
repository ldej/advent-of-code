// Package dijkstra
//
// Provides Dijkstra's Shortest Path Algorithm
// Assumes that all vertices have unique name
// Supports directed and undirected graphs
// Returns the taken path when calculating the shortest path
// Is able to calculate a distance

package dijkstra

import (
	"container/heap"
	"fmt"
	"strings"

	"github.com/ldej/advent-of-code/tools/myints"
)

type Graph struct {
	vertices    []*Vertex
	verticesMap map[string]*Vertex
}

func NewGraph() Graph {
	return Graph{
		verticesMap: map[string]*Vertex{},
	}
}

func (g *Graph) AddDirectedEdge(fromName string, toName string, distance int) {
	from, found := g.verticesMap[fromName]
	if !found {
		from = &Vertex{Name: fromName}
		g.verticesMap[fromName] = from
		g.vertices = append(g.vertices, from)
	}

	to, found := g.verticesMap[toName]
	if !found {
		to = &Vertex{Name: toName}
		g.verticesMap[toName] = to
		g.vertices = append(g.vertices, to)
	}

	if from.Edges == nil {
		from.Edges = map[*Vertex]int{}
	}
	from.Edges[to] = distance
}

func (g *Graph) AddUndirectedEdge(fromName string, toName string, distance int) {
	g.AddDirectedEdge(fromName, toName, distance)
	g.AddDirectedEdge(toName, fromName, distance)
}

func (g *Graph) AddIntGrid(input [][]int) {
	// consider the int above, below, left and right as connected
	for y, row := range input {
		for x := range row {
			if x+1 < len(row) {
				g.AddUndirectedEdge(
					fmt.Sprintf("x=%d,y=%d", x, y),
					fmt.Sprintf("x=%d,y=%d", x+1, y),
					input[y][x+1],
				)
			}
			if y+1 < len(input) {
				g.AddUndirectedEdge(
					fmt.Sprintf("x=%d,y=%d", x, y),
					fmt.Sprintf("x=%d,y=%d", x, y+1),
					input[y+1][x],
				)
			}
		}
	}
}

func (g *Graph) Shortest(fromName string, toName string) *Path {
	start, found := g.verticesMap[fromName]
	if !found {
		panic(fmt.Sprintf("from vertex %q not found", fromName))
	}
	end, found := g.verticesMap[toName]
	if !found {
		panic(fmt.Sprintf("to vertex %q not found", fromName))
	}

	queue := newVisitor(start)

	for queue.Len() > 0 {
		current, path := queue.Get()
		if current == end {
			return &path
		}
		queue.Visit(current)
	}
	return nil
}

func (g *Graph) ShortestIntGrid(fromX, fromY, toX, toY int) *Path {
	return g.Shortest(
		fmt.Sprintf("x=%d,y=%d", fromX, fromY),
		fmt.Sprintf("x=%d,y=%d", toX, toY),
	)
}

// AllPaths returns the shortest distances between each pair of vertices in the graph
func (g *Graph) AllPaths() map[string]map[string]*Path {
	results := map[string]map[string]*Path{}
	for _, v1 := range g.vertices {
		result := map[string]*Path{}
		for _, v2 := range g.vertices {
			p := g.Shortest(v1.Name, v2.Name)
			if p != nil {
				result[v2.Name] = p
			}
		}
		results[v1.Name] = result
	}
	return results
}

type Vertex struct {
	Name  string
	Edges map[*Vertex]int
}

func (v Vertex) String() string {
	return v.Name
}

type Path struct {
	path      []Vertex
	distances []int
}

func (p Path) Distance() int {
	return myints.Sum(p.distances)
}

func (p Path) Add(v Vertex, distance int) Path {
	return Path{
		path:      append(append([]Vertex{}, p.path...), v),
		distances: append(append([]int{}, p.distances...), distance),
	}
}

func (p Path) Print() string {
	var names []string
	for _, v := range p.path {
		names = append(names, v.Name)
	}
	return strings.Join(names, " -> ")
}

func newVisitor(start *Vertex) *visitor {
	q := &visitor{
		paths: map[string]Path{
			start.Name: {
				path:      []Vertex{*start},
				distances: []int{},
			},
		},
		visited: map[string]bool{},
	}
	heap.Init(q)
	heap.Push(q, start)
	return q
}

type visitor struct {
	toVisit []*Vertex
	paths   map[string]Path
	visited map[string]bool
}

func (h *visitor) Get() (*Vertex, Path) {
	v := heap.Pop(h).(*Vertex)
	return v, h.paths[v.Name]
}

func (h *visitor) Visit(current *Vertex) {
	for _, neighbor := range h.unvisitedNeighbors(current) {
		currentPath, found := h.paths[current.Name]
		if !found {
			panic("can't find current path")
		}
		proposedPath := currentPath.Add(*neighbor, current.Edges[neighbor])
		neighborPath, found := h.paths[neighbor.Name]
		if !found || proposedPath.Distance() < neighborPath.Distance() {
			h.paths[neighbor.Name] = proposedPath
			heap.Push(h, neighbor)
		}
	}
	h.visited[current.Name] = true
}

func (h *visitor) unvisitedNeighbors(v *Vertex) []*Vertex {
	var neighbors []*Vertex
	for n, _ := range v.Edges {
		if _, found := h.visited[n.Name]; !found {
			neighbors = append(neighbors, n)
		}
	}
	return neighbors
}

func (h *visitor) Len() int { return len(h.toVisit) }
func (h *visitor) Less(i, j int) bool {
	return h.paths[h.toVisit[i].Name].Distance() < h.paths[h.toVisit[j].Name].Distance()
}
func (h *visitor) Swap(i, j int) {
	h.toVisit[i], h.toVisit[j] = h.toVisit[j], h.toVisit[i]
}

func (h *visitor) Push(x interface{}) {
	h.toVisit = append(h.toVisit, x.(*Vertex))
}

func (h *visitor) Pop() interface{} {
	old := h.toVisit
	n := len(old)
	x := old[n-1]
	old[n-1] = nil
	h.toVisit = old[0 : n-1]
	return x
}
