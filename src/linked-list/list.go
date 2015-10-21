package list

import (
	"errors"
)

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
	n := NewNode(v)
	if l == nil || l.IsEmpty() {
		l.head = n
		l.tail = n
		l.length = 1
		return
	}
	oldTail := l.tail
	l.tail.next = n
	n.prev = oldTail
	l.tail = n
	l.length++
}

func (l *List) Prepend(v interface{}) {
	n := NewNode(v)
	if l == nil || l.IsEmpty() {
		l.head = n
		l.tail = n
		l.length = 1
		return
	}
	oldHead = l.head
	l.head.prev = n
	n.next = oldHead
	l.head = n
	l.length++
}

func (l *List) Add(v interface{}, index int) error {
	if index >= l.Length() {
		return errors.New("Index out of range")
	}
	n := NewNode(v)

	if l.IsEmpty() {
		l.head = n
		l.tail = n
		l.length = 1
		return nil
	}
	i, err := l.Get(index)
	if err != nil {
		return err
	}
	j := i.next
	i.next = n
	n.prev = i
	if j != nil {
		n.next = j
		j.prev = n
	} else {
		l.tail = n
	}
	l.length++

	return nil
}

func (l *List) Get(index int) (*Node, error) {
	if l.Length() <= index {
		return nil, errors.New("Index out of range")
	}
	if l == nil || l.IsEmpty() {
		return nil, errors.New("List is empty")
	}
	p := l.head
	for i := 0; i < index; i++ {
		p = p.next
	}
	return p, nil
}

func (l *List) Find(n *Node) (int, error) {
	return 0, nil
}

func (l *List) Remove(v interface{}) int {
	if l == nil || l.IsEmpty() {
		return -1
	}
	count := 0
	for n := l.head; n != nil; n = n.next {
		if n.value == v {
			if n.prev == nil {
				//is head
				l.head = n.next
				if n.next == nil {
					//is also tail
					l.tail = nil
					l.length = 0
					return 1
				}
				n.next.prev = nil
			} else if n.next == nil {
				//is tail but at least have two nodes
				l.tail = n.prev
				n.prev.next = nil
			} else {
				n.prev.next = n.next
				n.next.prev = n.prev
			}
			l.length--
			count++
		}
	}
	return count
}

func (l *List) Pop() (*Node, error) {
	if l == nil || l.IsEmpty() {
		return nil, errors.New("List is empty")
	}
	n := l.tail
	if l.Length() == 1 {
		l.head = nil
		l.tail = nil
	} else {
		l.tail = n.prev
		l.tail.next = nil
	}
	l.length--
	return n, nil
}

func (l *List) Shift() (*Node, error) {
	if l == nil || l.IsEmpty() {
		return nil, errors.New("List is empty")
	}
	n := l.head
	if l.Length() == 1 {
		l.head = nil
		l.tail = nil
	} else {
		l.head = n.next
		l.head.prev = nil
	}
	l.length--
	return n, nil
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
