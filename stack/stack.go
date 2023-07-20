package stack

import "errors"

var (
	ErrEmptyStack = errors.New("stack is empty")
)

type Stack[T any] struct {
	data []T
}

func New[T any]() *Stack[T] {
	return &Stack[T]{
		data: make([]T, 0),
	}
}

func (s *Stack[T]) Push(item T) {
	s.data = append(s.data, item)
}

func (s *Stack[T]) Pop() (T, error) {
	var item T

	if s.IsEmpty() {
		return item, ErrEmptyStack
	}

	item = s.data[len(s.data)-1]

	s.data = s.data[:len(s.data)-1]

	return item, nil
}

func (s *Stack[T]) Peek() (T, error) {
	var item T

	if s.IsEmpty() {
		return item, ErrEmptyStack
	}

	return s.data[len(s.data)-1], nil
}

func (s *Stack[T]) Size() int {
	return len(s.data)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}
