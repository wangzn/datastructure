package shell

import (
	"fmt"
	"testing"
)

func TestShellSort(t *testing.T) {
	ShellSort(make([]int, 0))
	a := []int{0}
	b := ShellSort(a)
	a = []int{1, 0}
	b = ShellSort(a)
	a = []int{3, 4, 9, 8, 7, 1, 2, 10, 5, 6}
	b = ShellSort(a)
	for i := 0; i < len(a); i++ {
		if b[i] != i+1 {
			t.Error("Sorting error")
			break
		}
	}

	fmt.Println(b)
}
