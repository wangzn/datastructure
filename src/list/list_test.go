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
	l.Add(1)
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(4)
	l.Add(5)
	if l.Length() != 7 {
		t.Error("List length error", l.Length())
	}
	str := l.String()
	if str != "1|1|1|2|3|4|5" {
		t.Error("List string error", str)
	}
	if !l.Exist(1) {
		t.Error("Can't find exist value 1")
	}
	if l.Exist(100) {
		t.Error("Can find non-exist value 100")
	}
	del := l.Delete(1)
	if !del {
		t.Error("Delete fail")
	}
	if l.Length() != 4 {
		t.Error("Length error after delete", l.Length())
	}
	str = l.String()
	if str != "2|3|4|5" {
		t.Error("List string error after delete", str)
	}
	del = l.Delete(5)
	if !del {
		t.Error("Delete fail")
	}
}
