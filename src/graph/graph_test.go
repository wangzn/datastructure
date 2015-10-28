package graph

import (
	"testing"
)

func TestSimpleGraph(t *testing.T) {
	g := NewGraph()
	if g.VertexCount() != 0 {
		t.Error("Init count fail")
	}
	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)
	if g.VertexCount() != 3 {
		t.Error("Vertex count error")
	}
	if g.CheckVertex(1) == false {
		t.Error("Check vertex error")
	}
	g.AddEdge(1, 2, 12)
	g.AddEdge(1, 3, 13)
	if g.IsEdge(2, 3) {
		t.Error("Invalid edge")
	}
	if !g.IsEdge(1, 3) {
		t.Error("Is edge check fail")
	}
	if g.EdgeCount() != 2 {
		t.Error("Edge count error")
	}
	v, err := g.GetEdgeValue(1, 3)
	if err != nil {
		t.Error("Get edge value error", err)
	}
	if v != 13 {
		t.Error("Get invalid edge value", v)
	}
	err = g.RemoveEdge(1, 3)
	if err != nil {
		t.Error("Remove edge error", err)
	}
	if g.IsEdge(1, 3) {
		t.Error("Is edge error after removing")
	}
	if g.EdgeCount() != 1 {
		t.Error("Edge count error after removing")
	}
}
