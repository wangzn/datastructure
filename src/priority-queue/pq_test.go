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
	v, pri := pq.Extract()
	if v != "a" || pri != 1 {
		t.Error("Extract value fail", v, pri)
	}
	v, pri = pq.Extract()
	if v != "b" || pri != 2 {
		t.Error("Extract value fail", v, pri)
	}
	v, pri = pq.Extract()
	if v != "c" || pri != 3 {
		t.Error("Extract value fail", v, pri)
	}
	v, pri = pq.Extract()
	if v != "d" || pri != 4 {
		t.Error("Extract value fail", v, pri)
	}
	if pq.Length() != 0 {
		t.Error("Length error after clean up")
	}
}
