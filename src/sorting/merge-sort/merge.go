package merge

import (
	"fmt"
)

func MergeSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	t := make([]int, len(a)/2+1)
	mid := len(a) / 2
	MergeSort(a[:mid])
	MergeSort(a[mid:])
	fmt.Println("sepearte sorting", a[:mid], a[mid:])
	l, r := 0, mid
	copy(t, a[:mid])
	for i := 0; ; i++ {
		if t[l] < a[r] {
			a[i] = t[l]
			l++
			if l == mid {
				break
			}
		} else {
			a[i] = a[r]
			r++
			if r == len(a) {
				copy(a[i+1:], t[l:mid])
				break
			}
		}
	}
	return a
}
