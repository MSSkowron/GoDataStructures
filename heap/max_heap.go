package heap

import "golang.org/x/exp/constraints"

// MaxHeap represents a max heap with elements of type T.
type MaxHeap[T constraints.Ordered] struct {
	arr []T
}

// NewMaxHeap creates a new max heap with elements of type T.
func NewMaxHeap[T constraints.Ordered]() *MaxHeap[T] {
	return &MaxHeap[T]{
		arr: make([]T, 0),
	}
}

// Insert inserts a new element into the heap.
func (h *MaxHeap[T]) Insert(val T) {
	h.arr = append(h.arr, val)
	h.heapifyUp()
}

// Extract extracts the minimum element from the heap.
func (h *MaxHeap[T]) Extract() (T, error) {
	var max T

	if h.IsEmpty() {
		return max, ErrHeapIsEmpty
	}
	max = h.arr[0]
	h.arr[0] = h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]
	h.heapifyDown()
	return max, nil
}

// IsEmpty returns true if the heap is empty, false otherwise.
func (h *MaxHeap[T]) IsEmpty() bool {
	return len(h.arr) == 0
}

// Size returns the size of the heap.
func (h *MaxHeap[T]) Size() int {
	return len(h.arr)
}

func (h *MaxHeap[T]) heapifyUp() {
	i := len(h.arr) - 1
	for parent(i) >= 0 && h.arr[parent(i)] < h.arr[i] {
		swap(h.arr, parent(i), i)
		i = parent(i)
	}
}

func (h *MaxHeap[T]) heapifyDown() {
	i := 0
	for left(i) < len(h.arr) {
		max := left(i)
		if right(i) < len(h.arr) && h.arr[right(i)] > h.arr[left(i)] {
			max = right(i)
		}
		if h.arr[i] > h.arr[max] {
			break
		}
		swap(h.arr, i, max)
		i = max
	}
}
