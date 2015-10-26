package pq

import (
	"testing"
)

func TestMain(t *testing.T) {
	pq := NewMinPQ()
	if pq == nil {
		t.Error("New pq error")
	}
	pq.Insert("b", 2)
	pq.Insert("d", 4)
	pq.Insert("a", 1)
	pq.Insert("c", 3)
	if pq.Length() != 4 {
		t.Error("Length error after insert")
	}
}
