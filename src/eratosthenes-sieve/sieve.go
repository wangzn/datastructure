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
	sqrt := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqrt; i++ {
		if !b[i] {
			for j := 2; j*i < n+1; j++ {
				b[j*i] = true
			}
		}
	}
	for i := 2; i < n+1; i++ {
		if !b[i] {
			ret = append(ret, i)
		}
	}
	return ret
}
