package rselect

import (
	"math/rand"
)

func doSelect(a []int, n, i int) int {
	if n == 1 || len(a) == 1 {
		return a[0]
	}
	pivot := rand.Intn(n)
	pivot = n / 2
	temp := a[pivot]
	a[pivot], a[n-1] = a[n-1], a[pivot]
	left := 0
	for j := 0; j < n-1; j++ {
		if a[j] < temp {
			a[j], a[left] = a[left], a[j]
			left++
		}
	}
	a[left], a[n-1] = a[n-1], a[left]
	if left+1 == i {
		return a[left]
	} else if left+1 < i {
		return doSelect(a[left+1:], n-left-1, i-left-1)
	} else {
		return doSelect(a[:left], left, i)
	}
}

func RSelect(a []int, i int) int {
	return doSelect(a, len(a), i)
}
