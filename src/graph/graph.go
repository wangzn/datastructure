package graph

import (
	"errors"
	"fmt"
)

type VertexId uint

type Vertices []VertexId

type Edge struct {
	from  VertexId
	to    VertexId
	value int
}

type Graph struct {
	edges      map[VertexId]map[VertexId]int
	edgesCount int
	isDirected bool
}

func NewGraph() *Graph {
	return &Graph{
		edges: make(map[VertexId]map[VertexId]int),
	}
}

func (g *Graph) CheckVertex(vertex VertexId) bool {
	_, ok := g.edges[vertex]
	return ok
}

func (g *Graph) TouchVertex(vertex VertexId) {
	if _, ok := g.edges[vertex]; !ok {
		g.edges[vertex] = make(map[VertexId]int)
	}
}

func (g *Graph) AddVertex(vertex VertexId) error {
	if _, ok := g.edges[vertex]; ok {
		return errors.New("Vertex have already exist")
	} else {
		g.edges[vertex] = make(map[VertexId]int)
	}
	return nil
}
func (g *Graph) RemoveVertex(vertex VertexId) error {
	if _, ok := g.edges[vertex]; !ok {
		return errors.New("Vertex does not exist")
	}
	delete(g.edges, vertex)

	for _, v := range g.edges {
		if _, connected := v[vertex]; connected {
			delete(v, vertex)
		}
	}
	return nil
}

func (g *Graph) IsVertex(vertex VertexId) bool {
	_, ok := g.edges[vertex]
	return ok
}

func (g *Graph) VertexCount() int {
	return len(g.edges)
}

func (g *Graph) AddEdge(from, to VertexId, value int) error {
	if from == to {
		return errors.New("From and to vertex is the same")
	}
	if !g.CheckVertex(from) {
		return errors.New(fmt.Sprintf("%s %v %s", "Vertex", from, "does not exist"))
	}
	if !g.CheckVertex(to) {
		return errors.New(fmt.Sprintf("%s %v %s", "Vertex", to, "does not exist"))
	}
	if e, fromok := g.edges[from]; fromok {
		if _, took := e[to]; took {
			return errors.New("Edge exists")
		}
	}
	g.TouchVertex(from)
	g.TouchVertex(to)
	g.edges[from][to] = value
	if !g.isDirected {
		g.edges[to][from] = value
	}
	g.edgesCount++
	return nil
}

func (g *Graph) RemoveEdge(from, to VertexId) error {
	if !g.IsVertex(from) {
		return errors.New("From vertex does not exist")
	}
	if !g.IsVertex(to) {
		return errors.New("To vertex does not exist")
	}
	if _, ok := g.edges[from][to]; !ok {
		return errors.New("Edge does not exist")
	}
	delete(g.edges[from], to)
	if !g.isDirected {
		if _, ok := g.edges[to][from]; !ok {
			return errors.New("Edge to -> from does not exist")
		}
		delete(g.edges[to], from)
	}
	g.edgesCount--
	return nil
}

func (g *Graph) IsEdge(from, to VertexId) bool {
	if !g.IsVertex(from) || !g.IsVertex(to) {
		return false
	}
	if _, ok := g.edges[from][to]; !ok {
		return false
	}
	if !g.isDirected {
		if _, ok := g.edges[to][from]; !ok {
			return false
		}
	}
	return true
}

func (g *Graph) EdgesCount() int {
	return g.edgesCount
}

func (g *Graph) GetEdgeValue(from, to VertexId) (int, error) {
	if g.IsEdge(from, to) {
		return g.edges[from][to], nil
	}
	return 0, errors.New("Edge does not exist")
}
