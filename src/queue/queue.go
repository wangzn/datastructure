package queue

import (
	"errors"
	"sync"
)

type Queue struct {
	data   []interface{}
	length int
	sync.Mutex
}

func New() *Queue {
	return &Queue{
		data:   make([]interface{}, 0),
		length: 0,
	}
}

func (q *Queue) Length() int {
	q.Lock()
	defer q.Unlock()
	return q.length
}

func (q *Queue) IsEmpty() bool {
	return q.Length() == 0
}

func (q *Queue) Push(i interface{}) {
	q.Lock()
	defer q.Unlock()
	q.data = append(q.data, i)
	q.length++
}

func (q *Queue) Shift() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("Queue is empty")
	}
	q.Lock()
	defer q.Unlock()
	i := q.data[0]
	q.data = q.data[1:]
	q.length--
	return i, nil
}

func (q *Queue) Peek() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("Queue is empty")
	}
	return q.data[0], nil
}
