package bst

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	n := NewNode(6)
	m := NewNode(11)
	if n.Compare(m) != -1 || m.Compare(n) != 1 || n.Compare(m) == 0 {
		t.Error("Test node compare fail")
	}
	tree := NewTree(n)
	if tree.Size() != 1 {
		t.Error("New tree size error")
	}

	tree.Insert(4)
	tree.Insert(3)
	tree.Insert(1)
	tree.Insert(7)
	tree.Insert(5)
	str := tree.Serialize()
	fmt.Println(str)
	tree2 := Deserialize(str)
	fmt.Println(tree2.Serialize())
	if tree.Size() != 6 {
		t.Error("Size error after insert")
	}
	if tree.Search(2) != false {
		t.Error("Search non-exist value fail")
	}
	if tree.Search(7) == false {
		t.Error("Search exist value fail")
	}
	l, r := n.Child()
	if l.Value() != 4 || r.Value() != 7 {
		t.Error("Tree node value error", l.Value(), r.Value())
	}
	if tree2.Size() != 6 {
		t.Error("Deserialized tree size error after insert")
	}
	if tree2.Search(2) != false {
		t.Error("Deserialized tree search non-exist value fail")
	}
	if tree2.Search(7) == false {
		t.Error("Deserialized tree search exist value fail")
	}
	l, r = tree2.Head().Child()
	if l.Value() != 4 || r.Value() != 7 {
		t.Error("Tree node value error", l.Value(), r.Value())
	}

	tree.Delete(2)
	if tree.Size() != 6 {
		t.Error("Size error after non-exist value delete")
	}
	tree.Delete(3)
	if tree.Size() != 5 {
		t.Error("Size error after exist value delete")
	}
	tree.Insert(8)
	if tree.Size() != 6 {
		t.Error("Size error after delete and insert")
	}
	if tree.Search(8) == false {
		t.Error("Search exist value fail")
	}
}
