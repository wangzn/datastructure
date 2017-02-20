package quick

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	QuickSort(make([]int, 0))
	a := []int{0}
	b := Quick(a)
	fmt.Println(b)
	a = []int{7, 9, 3, 6, 1, 5, 4, 8, 2}
	b = Quick(a)
	for i := 0; i < len(a); i++ {
		if a[i] != i+1 {
			t.Error("Sort error")
			break
		}
	}
	fmt.Println(b)
}
