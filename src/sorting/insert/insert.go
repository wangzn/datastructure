package insert

import ()

func InsertSort(a []int) []int {
	for i := 1; i < len(a); i++ {
		j := i - 1
		tmp := a[i]
		for j > -1 && tmp < a[j] {
			a[j+1] = a[j]
			j--
		}
		a[j+1] = tmp
	}
	return a
}
