package insert

import (
	"fmt"
	"testing"
)

func TestInsertSort(t *testing.T) {
	InsertSort(make([]int, 0))
	a := []int{1}
	b := InsertSort(a)
	a = []int{3, 4, 7, 1, 2, 5, 6}
	b = InsertSort(a)
	for i := 0; i < len(a); i++ {
		if b[i] != i+1 {
			t.Error("Sorting error")
			break
		}
	}
	fmt.Println(b)
}
