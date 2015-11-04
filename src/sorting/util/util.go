package util

import (
	"math/rand"
)

func GenArray(n int) []int {
	a := make([]int, n)
	base := rand.Intn(10)
	for i := 0; i < len(a); i++ {
		a[i] = rand.Intn(base * n)
	}
	return a
}

func CheckArray(a []int) int {
	for i := 1; i < len(a); i++ {
		if a[i] < a[i-1] {
			return i
		}
	}
	return 0
}

func CheckArrayDesc(a []int) int {
	for i := 1; i < len(a); i++ {
		if a[i] > a[i-1] {
			return i
		}
	}
	return 0
}
