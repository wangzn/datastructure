package selection

import (
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	SelectionSort(make([]int, 0))
	a := []int{0}
	b := SelectionSort(a)
	a = []int{4, 7, 1, 3, 2, 6, 5}
	b = SelectionSort(a)
	for i := 0; i < len(a); i++ {
		if a[i] != i+1 {
			t.Error("Sorting error")
			break
		}
	}
	fmt.Println(b)
}
