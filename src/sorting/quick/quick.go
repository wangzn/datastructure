package quick

import ()

func QuickSort(a []int) []int {
	if len(a) == 0 {
		return a
	}
	var partition func(left, right, pivot int) int
	var recurse func(left, right int)
	partition = func(left, right, pivot int) int {
		tmp := a[pivot]
		right--
		a[right], a[pivot] = a[pivot], a[right]
		for i := left; i < right; i++ {
			if a[i] <= tmp {
				a[i], a[left] = a[left], a[i]
				left++
			}
		}
		a[left], a[right] = a[right], a[left]
		return left
	}

	recurse = func(left, right int) {
		if left < right {
			var pivot int
			pivot = (left + right) / 2
			pivot = partition(left, right, pivot)
			recurse(left, pivot)
			recurse(pivot+1, right)
		}
	}

	recurse(0, len(a))
	return a
}
