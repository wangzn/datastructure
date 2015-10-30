package heap

import (
	"github.com/wangzn/datastructure/src/heap"
)

type MyInt int

func (i MyInt) Less(j heap.Item) bool {
	return int(i) < int(j.(MyInt))
}

func HeapSort(a []int) []int {
	h := heap.NewHeap()
	for i := 0; i < len(a); i++ {
		h.Insert(MyInt(a[i]))
	}
	i := 0
	for !h.IsEmpty() {
		a[i] = int(h.Extract().(MyInt))
		i++
	}
	return a
}
