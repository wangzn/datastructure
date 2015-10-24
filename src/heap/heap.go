package heap

import (
	"errors"
	"sync"
)

type Item interface {
	Less(b Item) bool
}

type Heap struct {
	sync.Mutex //anonymous
	data       []Item
	min        bool
}

func NewHeap() *Heap {
	return &Heap{
		data: make([]Item, 0),
	}
}

func NewMaxHeap() *Heap {
	return &Heap{
		data: make([]Item, 0),
		min:  false,
	}
}

func NewMinHeap() *Heap {
	return &Heap{
		data: make([]Item, 0),
		min:  true,
	}
}

func (h *Heap) Length() int {
	return len(h.data)
}

func (h *Heap) IsEmpty() bool {
	return h == nil || h.Length() == 0
}

func (h *Heap) Get(index int) (Item, error) {
	if h.Length() <= index {
		return nil, errors.New("Index out of range")
	}
	return h.data[index], nil
}

func (h *Heap) Parrent(index int) (Item, error) {
	if h.Length() <= index {
		return nil, errors.New("Index out of range")
	}
	parrent := int((index - 1) / 2)
	return h.Get(parrent)
}

func (h *Heap) Children(index int) (Item, Item, error) {
	if h.Length() <= index {
		return nil, nil, errors.New("Index out of range")
	}
	ic1, ic2 := 2*index+1, 2*index+2
	vc1, _ := h.Get(ic1)
	vc2, _ := h.Get(ic2)
	return vc1, vc2, nil
}

func (h *Heap) Insert(i Item) {
	h.Lock()
	defer h.Unlock()
	h.data = append(h.data, i)
	//h.siftUp()
	return
}

func (h *Heap) Less(a, b Item) bool {
	if h.min {
		return a.Less(b)
	} else {
		return b.Less(a)
	}
}
