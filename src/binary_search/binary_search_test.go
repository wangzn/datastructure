package bsearch

import (
	"testing"
)

func TestBSearch(t *testing.T) {
	a := []int{}
	b := 3
	if Search(a, b) != -1 {
		t.Error("Search empty list error")
	}
	a = []int{1, 2, 4, 5}
	if Search(a, b) != -1 {
		t.Error("Search non-exist value error")
	}
	a = []int{1, 2, 3}
	if c := Search(a, b); c != 2 {
		t.Error("Seach value error", c)
	}
}