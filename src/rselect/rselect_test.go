package rselect

import (
	"testing"
)

func TestMain(t *testing.T) {
	a := []int{3, 1, 2, 5, 6, 7, 8, 4, 9, 10}
	for i := 1; i <= len(a); i++ {
		r := RSelect(a, i)
		if r != i {
			t.Error("Select ", i, " th number error", r)
		}
	}

}
