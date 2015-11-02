package sieve

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	ret := Sieve(100)
	fmt.Println(ret)
}

func benchmarkGetAllPrimesTo(n int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sieve(n)
	}
}

func BenchmarkGetAllPrimesTo100(b *testing.B) {
	benchmarkGetAllPrimesTo(100, b)
}
func BenchmarkGetAllPrimesTo10000(b *testing.B) {
	benchmarkGetAllPrimesTo(10000, b)
}
func BenchmarkGetAllPrimesTo1000000(b *testing.B) {
	benchmarkGetAllPrimesTo(1000000, b)
}
