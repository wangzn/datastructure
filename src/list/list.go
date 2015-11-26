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

func NewList() *List {
	return &List{nil, 0}
}

func (l *List) Length() int {
	return l.length
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
	}
	return false
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
