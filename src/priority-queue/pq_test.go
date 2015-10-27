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
	v, err := pq.Extract()
	if err != nil {
		t.Error("Extract value with error", err)
	}
	if v.value != "a" || v.priority != 1 {
		t.Error("Extract value fail", v.value, v.priority)
	}
	v, err = pq.Extract()
	if v.value != "b" || v.priority != 2 {
		t.Error("Extract value fail", v.value, v.priority)
	}
	v, err = pq.Extract()
	if v.value != "c" || v.priority != 3 {
		t.Error("Extract value fail", v.value, v.priority)
	}
	v, err = pq.Extract()
	if v.value != "d" || v.priority != 4 {
		t.Error("Extract value fail", v.value, v.priority)
	}
	if pq.Length() != 0 {
		t.Error("Length error after clean up")
	}
	pq.Insert("b", 2)
	pq.Insert("d", 4)
	pq.Insert("a", 1)
	pq.Insert("c", 3)
	pq.ChangePriority("c", 0)
	v, err = pq.Extract()
	if v.value != "c" || v.priority != 0 {
		t.Error("change priority value error",
			v.value, v.priority)
	}
}
