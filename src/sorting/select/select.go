package selection

import (
	"fmt"
)

func SelectionSort(a []int) []int {
	var minIndex, min int
	for i := 0; i < len(a)-1; i++ {
		minIndex = -1
		min = a[i]
		for j := i + 1; j < len(a); j++ {
			if a[j] < min {
				minIndex = j
				min = a[j]
			}
		}
		if minIndex != -1 {
			a[i], a[minIndex] = a[minIndex], a[i]
		}
		fmt.Println(a)
	}
	return a
}
