package list

import (
	"testing"
)

func TestAdding(t *testing.T) {
	l1 := NewList()
	l2 := NewList()

	l1.Prepend(1)
	if l1.Length() != 1 {
		t.Error("length error after prepend")
	}
	l1.Prepend(2)
	if l1.Length() != 2 {
		t.Error("length error after prepend")
	}
	if l1.Head().Value() != 2 {
		t.Error("Head value is invalid", l1.Head().Value())
	}
	l1.Append(3)
	//2 1 3
	if l1.Tail().Value() != 3 {
		t.Error("Tail value is invalid", l1.Tail().Value())
	}
	nn, err := l1.Get(1)
	if err != nil || nn.Value() != 1 {
		t.Error("Get index fail")
	}
	r := l1.Remove(10)
	if r != 0 {
		t.Error("Remove non-exist value fail")
	}
	r = l1.Remove(1)
	if r != 1 {
		t.Error("Remove single exist value fail")
	}
	l2.Append(4)
	l2.Append(5)
	l1.Concat(l2)
	l1.Print()
	if l1.Length() != 4 {
		t.Error("Concat fail")
	}
	l1.Append(1)
	if l1.Length() != 5 {
		t.Error("Length error after append value")
	}
	if nn, err := l1.Pop(); err != nil || nn.Value() != 1 {
		t.Error("Pop value error")
	}
	l1.Print()
	if nn, err = l1.Shift(); err != nil || nn.Value() != 2 {
		t.Error("Shift value error")
	}
	l1.Print()
	if l1.Length() != 3 {
		t.Error("Length error after shift and pop")
	}
}
