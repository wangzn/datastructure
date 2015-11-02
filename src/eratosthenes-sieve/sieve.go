package sieve

import (
	"math"
)

func Sieve(n int) []int {
	ret := make([]int, 0)
	if n < 2 {
		return ret
	}
	b := make([]bool, n+1)
	for i := 0; i < n+1; i++ {
		b[i] = true
	}
	for i := 2; float64(i) < math.Sqrt(float64(n)); i++ {
		if b[i] {
			for j := 2; j*i < n+1; j++ {
				b[j*i] = false
			}
		}
	}
	for i := 2; i < n+1; i++ {
		if b[i] {
			ret = append(ret, i)
		}
	}
	return ret
}
