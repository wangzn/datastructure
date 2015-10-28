package graph

import (
	"testing"
)

func TestMain(t *testing.T) {
	g := NewGraph()
	if g.VertexCount() != 0 {
		t.Error("Init count fail")
	}
}
