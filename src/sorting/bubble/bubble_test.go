package bubble

import (
	"fmt"
	"testing"
)

func TestBubble(t *testing.T) {
	b := BubbleSort([]int{1})
	a := []int{4, 1, 7, 0, 3, 2, 6, 5}
	b = BubbleSort(a)
	for i := 0; i < 7; i++ {
		if b[i] != i {
			t.Error("Sort error")
			fmt.Println(b)
			return
		}
	}

}
