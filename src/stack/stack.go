package stack

import (
	"errors"
	"sync"
)

type Stack struct {
	data   []interface{}
	length int
	sync.Mutex
}

func New() *Stack {
	return &Stack{
		data:   make([]interface{}, 0),
		length: 0,
	}
}

func (s *Stack) IsEmpty() bool {
	return s.length == 0
}

func (s *Stack) Length() int {
	s.Lock()
	defer s.Unlock()
	return s.length
}

func (s *Stack) Push(i interface{}) {
	s.Lock()
	defer s.Unlock()
	data := make([]interface{}, 1)
	data[0] = i
	s.data = append(data, s.data...)
	s.length++
}

func (s *Stack) Pop() (interface{}, error) {
	if s.Length() == 0 {
		return nil, errors.New("Stack is empty")
	}
	s.Lock()
	defer s.Unlock()
	i := s.data[0]
	s.data = s.data[1:]
	s.length--
	return i, nil
}

func (s *Stack) Peek() (interface{}, error) {
	if s.Length() == 0 {
		return nil, errors.New("Stack is empty")
	}
	s.Lock()
	defer s.Unlock()
	return s.data[0], nil
}
