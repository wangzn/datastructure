package avl

import "testing"

func Test(t *testing.T) {
	tree := NewTree()

	tree.Insert(4)
	tree.Insert(2)
	tree.Insert(1)
	tree.Insert(7)
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(6)
	tree.Detail()
	tree.Delete(4)
	tree.Detail()

	tree.Delete(7)
	tree.Detail()

	tree.Delete(6)
	tree.Print()
	tree.Detail()
}
