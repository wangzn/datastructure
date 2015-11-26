package pq

import (
	"errors"
	"sync"

	"github.com/wangzn/datastructure/src/heap"
	"github.com/wangzn/datastructure/src/queue"
)

type Item struct {
	value    interface{}
	priority int
}

type PQ struct {
	l1   sync.Mutex
	l2   sync.Mutex
	data *heap.Heap
}

func (a Item) Less(b heap.Item) bool {
	return a.priority < b.(Item).priority
}

func NewPQ() *PQ {
	return &PQ{
		data: heap.NewHeap(),
	}
}

func NewMinPQ() *PQ {
	return &PQ{
		data: heap.NewMinHeap(),
	}
}

func NewMaxPQ() *PQ {
	return &PQ{
		data: heap.NewMaxHeap(),
	}
}

func (pq *PQ) Insert(v interface{}, pri int) {
	pq.l1.Lock()
	defer pq.l1.Unlock()
	i := Item{
		value:    v,
		priority: pri,
	}
	pq.data.Insert(i)
}

func (pq *PQ) Length() int {
	pq.l1.Lock()
	defer pq.l1.Unlock()
	return pq.data.Length()
}

func (pq *PQ) IsEmpty() bool {
	return pq.Length() == 0
}

func (pq *PQ) Extract() (Item, error) {
	if pq.IsEmpty() {
		return Item{}, errors.New("Queue is empty")
	}
	pq.l1.Lock()
	defer pq.l1.Unlock()
	i := pq.data.Extract().(Item)
	return i, nil
}

func (pq *PQ) ChangePriority(v interface{}, pri int) bool {
	pq.l2.Lock()
	defer pq.l2.Unlock()
	found := false
	q := queue.New()
	item, err := pq.Extract()

	for err == nil {
		q.Push(item)
		if item.value == v {
			found = true
			break
		}
		item, err = pq.Extract()
	}
	item.priority = pri
	pq.data.Insert(item)
	for q.Length() > 0 {
		j, _ := q.Shift()
		pq.data.Insert(j.(heap.Item))
	}
	return found
}

func (i Item) Value() interface{} {
	return i.value
}
