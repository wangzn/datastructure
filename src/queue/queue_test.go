package queue

import (
	"testing"
)

func TestMain(t *testing.T) {
	q := New()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)
	if q.Length() != 4 {
		t.Error("Length error after push")
	}
	v, err := q.Peek()
	if v != 1 {
		t.Error("Peek value error", v)
	}

	v, err = q.Shift()
	if v != 1 {
		t.Error("Shift value error", v)
	}
	if q.Length() != 3 {
		t.Error("Length error after shift", q.Length())
	}
	v, err = q.Shift()
	v, err = q.Shift()
	v, err = q.Shift()
	if !q.IsEmpty() {
		t.Error("Queue should be empty", q.Length())
	}
	v, err = q.Shift()
	if v != nil || err == nil {
		t.Error("Shift should get error")
	}
}
