package heap

import (
	"testing"
)

type Int int

func (a Int) Less(b Item) bool {
	return int(a) < int(b.(Int))
}

func TestMain(t *testing.T) {
	h := NewHeap()
	h.Insert(Int(2))
	if h.Length() != 1 {
		t.Error("Length error after insert")
	}

}
