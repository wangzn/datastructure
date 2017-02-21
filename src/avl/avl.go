package avl

import (
	"fmt"
	"log"
	"strconv"
)

//Node describe a tree node
type Node struct {
	value  int
	parent *Node
	left   *Node
	right  *Node
}

//Tree define a avl tree
type Tree struct {
	depth int
	root  *Node
}

//HasChild return true if node has any child
func (n *Node) HasChild() bool {
	if n == nil {
		return false
	}
	return n.left != nil || n.right != nil
}

//Depth return the node's depth.
func (n *Node) Depth() int {
	i := 0
	if n == nil {
		return 0
	}
	for n.parent != nil {
		i++
		n = n.parent
	}
	return i
}

//Length return
func (n *Node) Length() int {
	if n == nil {
		return 0
	}
	l := n.left.Length()
	r := n.right.Length()
	return max(l, r) + 1
}

//Diff return a depth diff of the Node.
//0 for equal
//>0 for left is deeper than right
//<0 for right is deeper than left
func (n *Node) Diff() int {
	if n == nil {
		return 0
	}
	l := n.left.Length()
	r := n.right.Length()
	return l - r
}

//ValidDiff return true if n's |Diff()| is < 2
func (n *Node) ValidDiff() bool {
	d := n.Diff()
	if d > 1 || d < -1 {
		return false
	}
	return true
}

//IsLeft return if node is parent's left node
//return false if parent is nil
func (n *Node) IsLeft() bool {
	if n.parent == nil {
		return false
	}
	return n.parent.left == n
}

//IsRight return if node is parent's right node
//return false if parent is nil
func (n *Node) IsRight() bool {
	if n.parent == nil {
		return false
	}
	return n.parent.right == n
}

//LeftMost returns the most left value in n's children
func (n *Node) LeftMost() *Node {
	if n == nil {
		return nil
	}
	for n.left != nil {
		n = n.left
		log.Println(n.value)
	}
	return n
}

//RightMost returns the most right value in n's children
func (n *Node) RightMost() *Node {
	if n == nil {
		return nil
	}
	for n.right != nil {
		n = n.right
	}
	return n
}

//String return a string of node
func (n *Node) String() string {
	return strconv.Itoa(n.value)
}

//NewTree initialize a empty tree
func NewTree() *Tree {
	t := new(Tree)
	t.root = nil
	return t
}

//Insert inserts a value into the tree, return node pointer
func (t *Tree) Insert(v int) *Node {
	n := new(Node)
	n.value = v
	if t.root == nil {
		t.root = n
		t.depth = 1
		return n
	}
	i := t.root
	for {
		if i.value == v {
			return i
		} else if v < i.value {
			//to left
			if i.left == nil {
				//insert here
				n.parent = i
				i.left = n
				break
			} else {
				i = i.left
			}
		} else {
			//to right
			if i.right == nil {
				n.parent = i
				i.right = n
				break
			} else {
				i = i.right
			}
		}
	}
	t.Rebalance(n.parent)
	return n
}

//Delete deletes node contains value v; return false if value does not exists
//TODO
func (t *Tree) Delete(v int) bool {
	n := t.Find(v)
	log.Println("delete node found")
	var p, q *Node
	fullChildren := true
	if n == nil {
		return false
	}
	if n.left == nil {
		p = n.right
		fullChildren = false
	} else if n.right == nil {
		p = n.left
		fullChildren = false
	}
	if !fullChildren {
		if n.IsLeft() {
			n.parent.left = p
		} else if n.IsRight() {
			n.parent.right = p
		} else {
			//no parent, is root
			t.root = nil
		}
		q = n
	} else {
		//left and right exist, find least value in right, delete and replace
		log.Println("Finding leftmost")
		least := n.right.LeftMost()
		log.Println("Found Leftmost", least.value)
		//least could not be nil, and least'left is nil, and least IsLeft()
		n.value = least.value
		least.parent.left = least.right
		if least.right != nil {
			least.right.parent = least.parent
		}
		q = least
		//TODO rebalance with least's parent
	}
	for q != nil {
		if !q.ValidDiff() {
			log.Println("balance value ", q.value)
			t.Rebalance(q)
		}
		q = q.parent
	}
	return true
}

//Find finds the node contains value v, nil if nothing
func (t *Tree) Find(v int) *Node {
	for i := t.root; i != nil; {
		if i.value == v {
			return i
		} else if v < i.value {
			i = i.left
		} else {
			i = i.right
		}
	}
	return nil
}

//Rebalance finds the lowest n with |Diff()| > 1; and try to balance the tree
func (t *Tree) Rebalance(n *Node) bool {
	f := false
	i := n
	for i != nil {
		if !i.ValidDiff() {
			f = true
			break
		}
		i = i.parent
	}
	if f {
		newroot := t.balance(i)
		if newroot == nil {
			log.Fatalln("Invalid newroot")
		}
		if newroot.parent == nil {
			t.root = newroot
		}
		return true
	}
	return false
}

//balance do balancing with node n as current top node
func (t *Tree) balance(n *Node) *Node {
	//rebalance as n is the subtree root
	var newroot *Node
	diff := n.Diff()
	if diff > 1 {
		//left greater
		if n.left.Diff() >= 0 {
			//rotate right
			newroot = rotateRight(n)
		} else if n.left.Diff() < 0 {
			//rotate left and right
			newroot = rotateLeftRight(n)
		}
	} else {
		//right greater
		if n.right.Diff() <= 0 {
			//rotate left
			newroot = rotateLeft(n)
		} else if n.right.Diff() > 0 {
			//rotate right and left
			newroot = rotateRightLeft(n)
		}
	}
	return newroot
}

//Print a tree
func (t *Tree) Print() {
	Walk(t.root, func(n *Node) {
		fmt.Printf("%s,", n.String())
	})
}

//Detail prints detail info of the tree
func (t *Tree) Detail() {
	Walk(t.root, func(n *Node) {
		left, right, parent := 0, 0, 0
		if n.left != nil {
			left = n.left.value
		}
		if n.right != nil {
			right = n.right.value
		}
		if n.parent != nil {
			parent = n.parent.value
		}
		fmt.Printf("%s p:%d l:%d r:%d\n", n.String(), parent, left, right)
	})
}

//Walk will do func f for node n and all children of n
func Walk(n *Node, f func(n *Node)) {
	if n == nil {
		return
	}
	if n.left != nil {
		Walk(n.left, f)
	}
	f(n)
	if n.right != nil {
		Walk(n.right, f)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rotateRight(n *Node) *Node {
	log.Println("do rotate right")
	parent := n.parent //might be nil
	newroot := n.left
	lr := n.left.right
	if n.IsLeft() {
		parent.left = newroot
	} else if n.IsRight() {
		parent.right = newroot
	}
	n.left = lr //might be nil
	if lr != nil {
		lr.parent = n
	}
	newroot.parent = parent
	newroot.right = n
	n.parent = newroot
	return newroot
}

func rotateLeft(n *Node) *Node {
	log.Println("do rotate left")
	parent := n.parent
	newroot := n.right
	rl := n.right.left
	if n.IsLeft() {
		parent.left = newroot
	} else if n.IsRight() {
		parent.right = newroot
	}
	n.right = rl
	if rl != nil {
		rl.parent = n
	}
	newroot.parent = parent
	newroot.left = n
	n.parent = newroot
	return newroot
}

func rotateLeftRight(n *Node) *Node {
	log.Println("do rotate right left")
	parent := n.parent
	l1 := n.left
	//step1
	newroot := l1.right
	n.left = newroot
	newroot.parent = n
	if newroot.left != nil {
		newroot.left.parent = l1
	}
	l1.right = newroot.left
	newroot.left = l1
	l1.parent = newroot
	//step2:
	lr := newroot.right
	if n.IsLeft() {
		parent.left = newroot
	} else if n.IsRight() {
		parent.right = newroot
	}
	newroot.parent = parent
	newroot.right = n
	n.parent = newroot
	n.left = lr
	if lr != nil {
		lr.parent = n
	}
	return newroot
}

func rotateRightLeft(n *Node) *Node {
	log.Println("do rotate right left")
	parent := n.parent
	l1 := n.right
	//step1:
	newroot := l1.left
	n.right = newroot
	newroot.parent = n
	if newroot.right != nil {
		newroot.right.parent = l1
	}
	l1.left = newroot.right
	newroot.right = l1
	l1.parent = newroot
	//step2:
	rl := newroot.left
	if n.IsLeft() {
		parent.left = newroot
	} else if n.IsRight() {
		parent.right = newroot
	}
	newroot.parent = parent
	newroot.left = n
	n.parent = newroot
	n.right = rl
	if rl != nil {
		rl.parent = n
	}
	return newroot
}
