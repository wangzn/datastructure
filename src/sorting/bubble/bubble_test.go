package bubble

import (
	"testing"

	"github.com/wangzn/datastructure/src/sorting/util"
)

func TestBubbleSort(t *testing.T) {
	BubbleSort(util.GenArray(0))
	a := BubbleSort(util.GenArray(100))
	c := util.CheckArray(a)
	if c != 0 {
		t.Error("Sort array fail, first fail index", c)
	}
}

func benchmarkBubbleSort(n int, b *testing.B) {
	a := util.GenArray(n)
	for i := 0; i < b.N; i++ {
		BubbleSort(a)
	}
}

func BenchmarkBubbleSort100(b *testing.B) {
	benchmarkBubbleSort(100, b)
}

func BenchmarkBubbleSort1000(b *testing.B) {
	benchmarkBubbleSort(1000, b)
}

func BenchmarkBubbleSort10000(b *testing.B) {
	benchmarkBubbleSort(10000, b)
}

func BenchmarkBubbleSort100000(b *testing.B) {
	benchmarkBubbleSort(100000, b)
}
