package list

import (
	"fmt"
	"strings"
)

type Node struct {
	value interface{}
	next  *Node
}

type List struct {
	head   *Node
	length int
}

func NewNode(v interface{}) *Node {
	return &Node{
		value: v,
		next:  nil,
	}
}

func (n *Node) Value() interface{} {
	return n.value
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) SetNext(t *Node) {
	n.next = t
}

func (n *Node) String() string {
	ss := make([]string, 0)
	cur := n
	for cur != nil {
		ss = append(ss, fmt.Sprintf("%v", cur.value))
		cur = cur.next
	}
	return strings.Join(ss, "|")
}

func NewList() *List {
	return &List{nil, 0}
}

func (l *List) Length() int {
	return l.length
}

func (l *List) Head() *Node {
	return l.head
}

func (l *List) Add(v interface{}) {
	n := &Node{
		value: v,
		next:  nil,
	}
	if l.head == nil {
		l.head = n
		l.length = 1
		return
	}
	cur := l.head
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = n
	l.length++
}

func (l *List) Exist(v interface{}) bool {
	cur := l.head
	for cur != nil {
		if cur.value == v {
			return true
		}
		cur = cur.next
	}
	return false
}

func (l *List) Delete(v interface{}) bool {
	if l.head == nil {
		return false
	}
	del := 0
	for l.head.value == v {
		l.head = l.head.next
		l.length--
		del++
	}
	cur := l.head
	for cur.next != nil {
		if cur.next.value == v {
			cur.next = cur.next.next
			l.length--
			del++
		}
		cur = cur.next
		if cur == nil {
			break
		}
	}
	return del > 0
}

func (l *List) String() string {
	ss := make([]string, 0)
	cur := l.head
	for cur != nil {
		ss = append(ss, fmt.Sprintf("%v", cur.value))
		cur = cur.next
	}
	return strings.Join(ss, "|")
}
