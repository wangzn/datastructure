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
	if h.Size() != 6 {
		t.Error("Hashtable size error", h.Size())
	}
	v, err := h.Get("aa")
	if err != nil || v != "a" {
		t.Error("Get value fail")
	}
	err = h.Del("aa")
	if err != nil {
		t.Error("Del key fail", err)
	}
	if h.Size() != 5 {
		t.Error("Hashtable size error after delete", h.Size())
	}
	v, err = h.Get("aa")
	if err == nil || v == "a" {
		t.Error("Get non-exist key value", v)
	}
	h.Set("cc", "abc")
	v, err = h.Get("cc")
	if v != "abc" {
		t.Error("Get new value fail", v)
	}
	h.Print()
}
