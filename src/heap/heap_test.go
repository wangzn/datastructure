package heap

import (
	"fmt"
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
	fmt.Println("Insert 2 3 4 1 6 5")
	h.Insert(Int(3))
	h.Insert(Int(4))
	h.Insert(Int(1))
	h.Insert(Int(6))
	h.Insert(Int(5))
	if h.Length() != 6 {
		t.Error("Length error after insert", h.Length())
	}
	v, err := h.Get(0)
	if err != nil {
		t.Error("Get value error", err.Error())
	}
	if int(v.(Int)) != 1 {
		t.Error("Get invalid value", v)
	}
	h.Print()
	v = h.Extract()
	if int(v.(Int)) != 1 {
		t.Error("Extract invalid first value", v)
	}
	if h.Length() != 5 {
		t.Error("Length error after first extract")
	}
	h.Print()
	v = h.Extract()
	if int(v.(Int)) != 2 {
		t.Error("Extract invalid second value", v)
	}
	if h.Length() != 4 {
		t.Error("Length error after second extract")
	}
	h.Print()
	h.Extract()
	h.Print()
	h.Extract()
	h.Print()
	h.Extract()
	h.Print()
	v = h.Extract()
	if int(v.(Int)) != 6 {
		t.Error("Extract invalid last value")
	}
	if h.Length() != 0 {
		t.Error("length error of empty heap")
	}
	if !h.IsEmpty() {
		t.Error("Empty heap judg error")
	}
}
