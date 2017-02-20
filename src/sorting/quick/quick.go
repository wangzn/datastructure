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

func Quick(a []int) []int {
	if len(a) < 2 {
		return a
	}
	var partition func(left, right, pivot int) int
	var recurse func(left, right int)

	partition = func(left, right, pivot int) int {
		if pivot > len(a)-1 {
			return -1
		}
		v := a[pivot]
		a[left], a[pivot] = a[pivot], a[left]
		j := left
		//generally, between i and j is those number who are larger than v
		for i := left + 1; i < right; i++ {
			if a[i] < v {
				//if a[i] is smaller than v
				j = j + 1
				//now a[j] is larger than v
				if i != j {
					a[j], a[i] = a[i], a[j]
				}
				//now a[j] is smaller than v
			}
		}
		//finally, a[j] is the last value smaller than v, swap to left
		a[j], a[left] = a[left], a[j]
		return j
	}

	recurse = func(left, right int) {
		if left < right {
			pivot := int(left/2 + right/2)
			pivot = partition(left, right, pivot)
			recurse(left, pivot)
			recurse(pivot+1, right)
		}
	}
	recurse(0, len(a))
	return a
}
