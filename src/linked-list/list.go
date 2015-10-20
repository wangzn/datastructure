package list

import ()

type Node struct {
	value interface{}
	next  *Node
	prev  *Node
}

type List struct {
	head   *Node
	tail   *Node
	length int
}

func NewNode(v interface{}) *Node {
	return &Node{
		value: v,
	}
}

func NewList(n *Node) *List {
	return &List{
		head:   n,
		tail:   n,
		length: 1,
	}
}

func (n *Node) Value() interface{} {
	if n == nil {
		return nil
	}
	return n.value
}

func (n *Node) Prev() *Node {
	if n == nil {
		return nil
	}
	return n.prev
}

func (n *Node) Next() *Node {
	if n == nil {
		return nil
	}
	return n.next
}

func (l *List) Length() int {
	return l.length
}

func (l *List) IsEmpty() bool {
	return l.length == 0
}

func (l *List) Head() *Node {
	return l.head
}

func (l *List) Tail() *Node {
	return l.tail
}

func (l *List) Append(v interface{}) {
}

func (l *List) Prepend(v interface{}) {
}

func (l *List) Add(v interface{}, index int) error {
	return nil
}

func (l *List) Get(index int) (interface{}, error) {
	return nil, nil
}

func (l *List) Find(n *Node) (int, error) {
	return 0, nil
}

func (l *List) Remove(v interface{}) bool {
	return false
}

func (l *List) Pop() (interface{}, error) {
	return nil, nil
}

func (l *List) Shift() (interface{}, error) {
	return nil, nil
}

func (l *List) Clear() {
	l.length = 0
	l.head = nil
	l.tail = nil
}

func (l *List) Concat(ll *List) {
}

func (l *List) Walk(f func(n *Node)) {
	for node := l.head; node != nil; node = node.next {
		f(node)
	}
}
