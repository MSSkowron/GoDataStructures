package binarysearch

import "golang.org/x/exp/constraints"

type ComparableOrdered interface {
	comparable
	constraints.Ordered
}

func Search[T ComparableOrdered](items []T, value T) int {
	left, right := 0, len(items)-1

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
