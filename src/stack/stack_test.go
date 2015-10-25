package stack

import (
	"testing"
)

func TestMain(t *testing.T) {
	s := New()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	if s.Length() != 4 {
		t.Error("Stack length error after push")
	}
	v, err := s.Pop()
	if v != 4 {
		t.Error("Stack pop value error")
	}
	if s.Length() != 3 {
		t.Error("length error after pop")
	}
	v, err = s.Pop()
	v, err = s.Peek()
	if v != 2 {
		t.Error("Peek value error")
	}
	v, err = s.Pop()
	v, err = s.Pop()
	v, err = s.Pop()
	if v != nil || err == nil {
		t.Error("Pop should be invalid")
	}
}
