package bubble

import ()

func BubbleSort(a []int) []int {
	for i := len(a); i > 0; i-- {
		for j := 0; j+1 < i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	return a
}
