package heap

import "golang.org/x/exp/constraints"

type MinHeap[T constraints.Ordered] struct {
	arr []T
}

// NewMinHeap returns a new min heap
func NewMinHeap[T constraints.Ordered]() *MinHeap[T] {
	return &MinHeap[T]{
		arr: make([]T, 0),
	}
}

// Insert inserts a new element into the heap
func (h *MinHeap[T]) Insert(val T) {
	h.arr = append(h.arr, val)
	h.heapifyUp()
}

// Extract extracts the minimum element from the heap
func (h *MinHeap[T]) Extract() (T, error) {
	var min T

	if h.IsEmpty() {
		return min, ErrHeapIsEmpty
	}
	min = h.arr[0]
	h.arr[0] = h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]
	h.heapifyDown()
	return min, nil
}

// IsEmpty returns true if the heap is empty, false otherwise
func (h *MinHeap[T]) IsEmpty() bool {
	return len(h.arr) == 0
}

// Size returns the size of the heap
func (h *MinHeap[T]) Size() int {
	return len(h.arr)
}

func (h *MinHeap[T]) heapifyUp() {
	i := len(h.arr) - 1
	for parent(i) >= 0 && h.arr[parent(i)] > h.arr[i] {
		swap(h.arr, parent(i), i)
		i = parent(i)
	}
}

func (h *MinHeap[T]) heapifyDown() {
	i := 0
	for left(i) < len(h.arr) {
		min := left(i)
		if right(i) < len(h.arr) && h.arr[right(i)] < h.arr[left(i)] {
			min = right(i)
		}
		if h.arr[i] < h.arr[min] {
			break
		}
		swap(h.arr, i, min)
		i = min
	}
}
