package exponentialsearch

import (
	"golang.org/x/exp/constraints"
)

type ComparableOrdered interface {
	comparable
	constraints.Ordered
}

func Search[T ComparableOrdered](items []T, value T) int {
	n := len(items)

	if n == 0 {
		return -1
	}

	idx := 1
	for idx < n {
		if items[idx] >= value {
			break
		}
		idx *= 2
	}

	left, right := idx/2, min(n-1, idx)
	for left <= right {
		mid := (left + right) / 2

		if items[mid] == value {
			return mid
		}

		if items[mid] < value {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
