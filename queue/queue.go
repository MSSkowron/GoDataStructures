package queue

import "errors"

var ErrEmptyQueue = errors.New("queue is empty")

// Queue is a generic queue implementation
type Queue[T any] struct {
	data []T
}

// New creates a new queue
func New[T any]() *Queue[T] {
	return &Queue[T]{
		data: make([]T, 0),
	}
}

// Enqueue adds an item to the queue
func (q *Queue[T]) Enqueue(item T) {
	q.data = append(q.data, item)
}

// Dequeue removes an item from the queue
func (q *Queue[T]) Dequeue() (T, error) {
	var item T

	if q.IsEmpty() {
		return item, ErrEmptyQueue
	}

	item = q.data[0]

	q.data = q.data[1:]

	return item, nil
}

// Peek returns the first item in the queue
func (q *Queue[T]) Peek() (T, error) {
	var item T

	if q.IsEmpty() {
		return item, ErrEmptyQueue
	}

	return q.data[0], nil
}

// IsEmpty returns true if the queue is empty, false otherwise
func (q *Queue[T]) IsEmpty() bool {
	return len(q.data) == 0
}

// Size returns the size of the queue
func (q *Queue[T]) Size() int {
	return len(q.data)
}
