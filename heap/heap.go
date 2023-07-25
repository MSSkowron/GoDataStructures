package heap

import (
	"errors"

	"golang.org/x/exp/constraints"
)

var (
	// ErrHeapIsEmpty is returned when the heap is empty
	ErrHeapIsEmpty = errors.New("heap is empty")
)

func swap[T constraints.Ordered](arr []T, i, j int) {
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
