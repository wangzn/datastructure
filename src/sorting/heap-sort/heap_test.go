package heap

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	HeapSort(make([]int, 0))
	a := []int{0}
	b := HeapSort(a)
	a = []int{7, 9, 3, 6, 1, 5, 4, 8, 2}
	b = HeapSort(a)
	for i := 0; i < len(a); i++ {
		if a[i] != i+1 {
			t.Error("Sort error")
			break
		}
	}
	fmt.Println(b)
}
