package bst

import (
	"fmt"
)

type Node struct {
	value   int
	parrent *Node
	left    *Node
	right   *Node
}

type Tree struct {
	head *Node
	size int
}

func NewNode(i int) *Node {
	return &Node{
		value: i,
	}
}

func NewTree(n *Node) *Tree {
	return &Tree{
		head: n,
		size: 1,
	}
}

func (n *Node) Compare(m *Node) int {
	if n.value < m.value {
		return -1
	} else if n.value > m.value {
		return 1
	}
	return 0
}

func (n *Node) Child() (*Node, *Node) {
	return n.left, n.right
}

func (n *Node) Value() int {
	return n.value
}

func (n *Node) Print() {
	fmt.Print("|", n.value)
}

func PrintNode(n *Node) {
	fmt.Println("----")
	if n != nil {
		fmt.Print("V:", n.value)
		if n.left != nil {
			fmt.Print("L:", n.left.value)
		}
		if n.right != nil {
			fmt.Print("R:", n.right.value)
		}
	}
	fmt.Println("\n----")
}

func Walk(h *Node, f func(n *Node)) {
	if h == nil {
		return
	}
	if h.left != nil {
		Walk(h.left, f)
	}
	f(h)
	if h.right != nil {
		Walk(h.right, f)
	}
}

func (t *Tree) Size() int {
	return t.size
}

func (t *Tree) Insert(i int) {
	n := NewNode(i)
	if t.head == nil {
		t.head = n
		t.size++
		return
	}
	head := t.head
	for {
		if n.Compare(head) < 0 {
			if head.left == nil {
				head.left = n
				n.parrent = head
				t.size++
				return
			} else {
				head = head.left
				continue
			}
		} else {
			if head.right == nil {
				head.right = n
				n.parrent = head
				t.size++
				return
			} else {
				head = head.right
				continue
			}
		}
	}
}

func (t *Tree) Search(i int) bool {
	if t == nil {
		return false
	}
	head := t.head
	n := NewNode(i)
	for head != nil {
		if n.Compare(head) < 0 {
			head = head.left
		} else if n.Compare(head) > 0 {
			head = head.right
		} else {
			return true
		}

	}
	return false
}

func (t *Tree) Delete(i int) bool {
	if t == nil || t.head == nil {
		return false
	}
	var parrent *Node
	head := t.head
	n := NewNode(i)
	for head != nil {
		if n.Compare(head) < 0 {
			parrent = head
			head = head.left
		} else if n.Compare(head) > 0 {
			parrent = head
			head = head.right
		} else {
			//just this node
			if head.left != nil {
				right := head.right //record temp right subtree
				head.value = head.left.value
				head.left = head.left.left
				if head.right != nil {
					head.right = head.right.right
				}
				if right != nil {
					//todo insert right subtree to current tree
					subtree := NewTree(head)
					Walk(right, func(n *Node) {
						subtree.Insert(n.value)
					})
				}
				t.size--
				return true
			}
			if head.right != nil {
				head.value = head.right.value
				head.left = head.right.left
				head.right = head.right.right
				t.size--
				return true
			}
			if parrent == nil {
				t.size = 0
				t.head = nil
				return true
			}
			if parrent.left == head {
				parrent.left = nil
			} else {
				parrent.right = nil
			}
			t.size--
			return true
		}
	}
	return false
}
