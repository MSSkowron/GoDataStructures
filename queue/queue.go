package queue

import "errors"

var ErrEmptyQueue = errors.New("queue is empty")

type Queue[T any] struct {
	data []T
}

func New[T any]() *Queue[T] {
	return &Queue[T]{
		data: make([]T, 0),
	}
}

func (q *Queue[T]) Enqueue(item T) {
	q.data = append(q.data, item)
}

func (q *Queue[T]) Dequeue() (T, error) {
	var item T

	if q.IsEmpty() {
		return item, ErrEmptyQueue
	}

	item = q.data[0]

	q.data = q.data[1:]

	return item, nil
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.data) == 0
}
