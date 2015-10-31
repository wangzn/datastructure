package sort

import (
	"testing"

	"github.com/wangzn/datastructure/src/sorting/util"
)

func benchmarkSort(f func([]int) []int, n int, b *testing.B) {
	a := util.GenArray(n)
	for i := 0; i < b.N; i++ {
		f(a)
	}
}

func BenchmarkBubbleSort1(b *testing.B) {
	benchmarkSort(BubbleSort, 100, b)
}
func BenchmarkBubbleSort2(b *testing.B) {
	benchmarkSort(BubbleSort, 1000, b)
}
func BenchmarkBubbleSort3(b *testing.B) {
	benchmarkSort(BubbleSort, 10000, b)
}

func BenchmarkInsertSort1(b *testing.B) {
	benchmarkSort(InsertSort, 100, b)
}
func BenchmarkInsertSort2(b *testing.B) {
	benchmarkSort(InsertSort, 1000, b)
}
func BenchmarkInsertSort3(b *testing.B) {
	benchmarkSort(InsertSort, 10000, b)
}

func BenchmarkHeapSort1(b *testing.B) {
	benchmarkSort(HeapSort, 100, b)
}
func BenchmarkHeapSort2(b *testing.B) {
	benchmarkSort(HeapSort, 1000, b)
}
func BenchmarkHeapSort3(b *testing.B) {
	benchmarkSort(HeapSort, 10000, b)
}

func BenchmarkSelectionSort1(b *testing.B) {
	benchmarkSort(SelectionSort, 100, b)
}
func BenchmarkSelectionSort2(b *testing.B) {
	benchmarkSort(SelectionSort, 1000, b)
}
func BenchmarkSelectionSort3(b *testing.B) {
	benchmarkSort(SelectionSort, 10000, b)
}

func BenchmarkQuickSort1(b *testing.B) {
	benchmarkSort(QuickSort, 100, b)
}
func BenchmarkQuickSort2(b *testing.B) {
	benchmarkSort(QuickSort, 1000, b)
}
func BenchmarkQuickSort3(b *testing.B) {
	benchmarkSort(QuickSort, 10000, b)
}

func BenchmarkMergeSort1(b *testing.B) {
	benchmarkSort(MergeSort, 100, b)
}
func BenchmarkMergeSort2(b *testing.B) {
	benchmarkSort(MergeSort, 1000, b)
}
func BenchmarkMergeSort3(b *testing.B) {
	benchmarkSort(MergeSort, 10000, b)
}

func BenchmarkShellSort1(b *testing.B) {
	benchmarkSort(ShellSort, 100, b)
}
func BenchmarkShellSort2(b *testing.B) {
	benchmarkSort(ShellSort, 1000, b)
}
func BenchmarkShellSort3(b *testing.B) {
	benchmarkSort(ShellSort, 10000, b)
}
