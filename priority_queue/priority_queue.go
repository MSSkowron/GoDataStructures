package priorityqueue

import (
	"errors"

	"golang.org/x/exp/constraints"
)

var ErrPriorityQueueIsEmpty = errors.New("priority queue is empty")

// PriorityQueue is a data structure that stores items in order of priority.
// The item with the highest priority is always at the front of the queue.
// The priority queue is implemented using a max heap.
// K is the type of the priority.
// V is the type of the value.
type PriorityQueue[K constraints.Ordered, V any] struct {
	arr []item[K, V]
}

type item[K constraints.Ordered, V any] struct {
	priority K
	value    V
}

// New creates a new, empty priority queue.
// K is the type of the priority.
// V is the type of the value.
func New[K constraints.Ordered, V any]() *PriorityQueue[K, V] {
	return &PriorityQueue[K, V]{
		arr: make([]item[K, V], 0),
	}
}

// Enqueue adds an item to the priority queue.
func (pq *PriorityQueue[K, V]) Enqueue(priority K, value V) {
	pq.arr = append(pq.arr, item[K, V]{
		priority: priority,
		value:    value,
	})

	pq.heapifyUp()
}

// Dequeue removes the item with the highest priority from the priority queue.
func (pq *PriorityQueue[K, V]) Dequeue() (V, error) {
	var value V
	if pq.IsEmpty() {
		return value, ErrPriorityQueueIsEmpty
	}

	item := pq.arr[0]

	pq.arr[0] = pq.arr[len(pq.arr)-1]
	pq.arr = pq.arr[:len(pq.arr)-1]
	pq.heapifyDown()

	return item.value, nil
}

// IsEmpty returns true if the priority queue is empty, false otherwise.
func (pq *PriorityQueue[K, V]) IsEmpty() bool {
	return len(pq.arr) == 0
}

// Size returns the size of the priority queue.
func (pq *PriorityQueue[K, V]) Size() int {
	return len(pq.arr)
}

func (pq *PriorityQueue[K, V]) heapifyUp() {
	i := len(pq.arr) - 1
	for parent(i) >= 0 && pq.arr[parent(i)].priority < pq.arr[i].priority {
		swap(pq.arr, parent(i), i)
		i = parent(i)
	}
}

func (pq *PriorityQueue[K, V]) heapifyDown() {
	i := 0
	for left(i) < len(pq.arr) {
		max := left(i)
		if right(i) < len(pq.arr) && pq.arr[right(i)].priority > pq.arr[left(i)].priority {
			max = right(i)
		}
		if pq.arr[i].priority > pq.arr[max].priority {
			break
		}
		swap(pq.arr, i, max)
		i = max
	}
}

func swap[K constraints.Ordered, V any](arr []item[K, V], i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}
