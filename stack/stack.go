package stack

import "errors"

// ErrEmptyStack is returned when the stack is empty.
var ErrEmptyStack = errors.New("stack is empty")

// Stack represents a stack of items of type T.
type Stack[T any] struct {
	data []T
}

// New creates a new, empty stack of items of type T.
func New[T any]() *Stack[T] {
	return &Stack[T]{
		data: make([]T, 0),
	}
}

// Push adds an item to the stack.
func (s *Stack[T]) Push(item T) {
	s.data = append(s.data, item)
}

// Pop removes an item from the stack.
func (s *Stack[T]) Pop() (T, error) {
	var item T

	if s.IsEmpty() {
		return item, ErrEmptyStack
	}

	item = s.data[len(s.data)-1]

	s.data = s.data[:len(s.data)-1]

	return item, nil
}

// Peek returns the item at the top of the stack.
func (s *Stack[T]) Peek() (T, error) {
	var item T

	if s.IsEmpty() {
		return item, ErrEmptyStack
	}

	return s.data[len(s.data)-1], nil
}

// Size returns the size of the stack.
func (s *Stack[T]) Size() int {
	return len(s.data)
}

// IsEmpty returns true if the stack is empty, false otherwise.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}
