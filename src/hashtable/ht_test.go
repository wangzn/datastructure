package ht

import (
	"testing"
)

func TestMain(t *testing.T) {
	h := New(5)
	h.Set("a", 1)
	h.Set("b", 2)
	h.Set("c", 3)
	h.Set("aa", "a")
	h.Set("bb", "b")
	h.Set("cc", "cc")
	h.Print()
}
