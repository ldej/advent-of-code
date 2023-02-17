package dijkstra

import (
	"bytes"
	"os"
	"text/template"
)

func (g *Graph) Mermaid(p *Path) string {
	type Edge struct {
		FromName    string
		ToName      string
		Distance    int
		Directional bool
	}
	edgesMap := map[string]*Edge{}
	var edges []*Edge

	// Collect all edges to be drawn
	for _, vertex := range g.vertices {
		for edge, distance := range vertex.Edges {
			e, found := edgesMap[edge.Name+","+vertex.Name]
			if found && e.Distance == distance {
				e.Directional = false
			} else {
				n := Edge{
					FromName:    vertex.Name,
					ToName:      edge.Name,
					Distance:    distance,
					Directional: true,
				}
				edgesMap[vertex.Name+","+edge.Name] = &n
				edges = append(edges, &n)
			}
		}
	}

	// If a path is provided, highlight the vertex and edges
	var pathEdgeIndexes []int
	var pathVertexNames []string
	var pathStartName, pathEndName string
	if p != nil {
		for i := 0; i < len(p.path)-1; i++ {
			for j, edge := range edges {
				if (p.path[i].Name == edge.FromName && p.path[i+1].Name == edge.ToName) ||
					(p.path[i].Name == edge.ToName && p.path[i+1].Name == edge.FromName) {
					pathEdgeIndexes = append(pathEdgeIndexes, j)
				}
			}
		}
		for _, v := range p.path {
			pathVertexNames = append(pathVertexNames, v.Name)
		}
		if len(p.path) > 0 {
			pathStartName = p.path[0].Name
			pathEndName = p.path[len(p.path)-1].Name
		}
	}

	file, err := os.ReadFile("mermaid.gotmpl")
	if err != nil {
		panic(err)
	}
	tmpl, err := template.New("mermaid").Parse(string(file))
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, struct {
		Edges           []*Edge
		PathEdgeIndexes []int
		PathVertexNames []string
		PathStartName   string
		PathEndName     string
	}{
		Edges:           edges,
		PathEdgeIndexes: pathEdgeIndexes,
		PathVertexNames: pathVertexNames,
		PathStartName:   pathStartName,
		PathEndName:     pathEndName,
	})
	if err != nil {
		panic(err)
	}
	return buf.String()
}
