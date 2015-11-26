package list

import (
	"testing"
)

func TestMain(t *testing.T) {
	l := NewList()
	if l.Length() != 0 {
		t.Error("Empty list length error", l.Length())
	}
	l.Add(1)
	l.Add(2)
	l.Add(3)
	if l.Length() != 3 {
		t.Error("List length error", l.Length())
	}
	str := l.String()
	if str != "1|2|3" {
		t.Error("List string error", str)
	}
}
