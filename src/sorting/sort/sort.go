package sort

import (
	"github.com/wangzn/datastructure/src/heap"
)

type MyInt int

func (i MyInt) Less(j heap.Item) bool {
	return int(i) < int(j.(MyInt))
}

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

func MergeSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	t := make([]int, len(a)/2+1)
	mid := len(a) / 2
	MergeSort(a[:mid])
	MergeSort(a[mid:])
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
	}
	return a
}

func ShellSort(a []int) []int {
	var increment int
	increment = len(a) / 2
	for increment > 0 {
		for i := 1; i < len(a); i++ {
			j := i - increment
			tmp := a[i]
			for j >= 0 && tmp < a[j] {
				a[j+increment] = a[j]
				j = j - increment
			}
			a[j+increment] = tmp
		}
		if increment == 2 {
			increment = 1
		} else {
			increment = increment / 2
		}
	}
	return a
}

func HeapSort(a []int) []int {
	h := heap.NewHeap()
	for i := 0; i < len(a); i++ {
		h.Insert(MyInt(a[i]))
	}
	i := 0
	for !h.IsEmpty() {
		a[i] = int(h.Extract().(MyInt))
		i++
	}
	return a
}
