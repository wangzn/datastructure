package heap

import (
	"errors"
	"fmt"
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
		min:  true, //user min heap for default
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
	h.siftUp()
	return
}

func (h *Heap) Extract() Item {
	h.Lock()
	defer h.Unlock()
	if h.IsEmpty() {
		return nil
	}
	ret := h.data[0]
	last := h.data[h.Length()-1]
	if h.Length() == 1 {
		h.data = nil
	} else {
		h.data = append([]Item{last}, h.data[1:h.Length()-1]...)
		h.siftDown()
	}
	return ret

}

// siftup just move the last one to the right index
func (h *Heap) siftUp() {
	var p int
	for i := h.Length() - 1; i > 0; {
		p = int((i - 1) / 2)
		l, err := h.LessIndex(i, p)
		if err != nil {
			break
		}
		if l {
			//child is less than parrent
			h.data[i], h.data[p] = h.data[p], h.data[i]
		}
		i = p
	}
}

// siftdown just move the first one to the righe index
func (h *Heap) siftDown() {
	var c1, c2, minc int
	for i := 0; i < h.Length(); {
		c1 = 2*i + 1
		c2 = 2*i + 2
		minc = c1
		v1, err1 := h.Get(c1)
		v2, err2 := h.Get(c2)
		if err1 != nil {
			//no child
			break
		}
		if err2 != nil {
			//second child does not exist
		} else {
			if h.Less(v2, v1) {
				minc = c2
			}
		}
		if l, _ := h.LessIndex(minc, i); l {
			// if small child is less than parrent
			h.data[minc], h.data[i] = h.data[i], h.data[minc]
		}
		i = minc
	}
}

func (h *Heap) Less(a, b Item) bool {
	if h.min {
		return a.Less(b)
	} else {
		return b.Less(a)
	}
}

func (h *Heap) LessIndex(a, b int) (bool, error) {
	v1, err1 := h.Get(a)
	v2, err2 := h.Get(b)
	if err1 != nil || err2 != nil {
		return false, errors.New("Index out of range")
	}
	return h.Less(v1, v2), nil
}

func (h *Heap) Print() {
	for i := 0; i < h.Length(); i++ {
		fmt.Print(h.data[i], " ")
	}
	fmt.Println()
}
