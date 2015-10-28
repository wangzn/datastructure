package graph

import (
	"testing"
)

func TestMain(t *testing.T) {
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
}
