package pq

import (
	"github.com/wangzn/datastructure/src/heap"
)

type Item struct {
	value    interface{}
	priority int
}

type PQ struct {
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
	i := Item{
		value:    v,
		priority: pri,
	}
	pq.data.Insert(i)
}

func (pq *PQ) Extract() (interface{}, int) {
	i := pq.data.Extract().(Item)
	return i.value, i.priority
}

func (pq *PQ) Length() int {
	return pq.data.Length()
}
