package ternarysearch

import "golang.org/x/exp/constraints"

type ComparableOrdered interface {
	comparable
	constraints.Ordered
}

func Search[T ComparableOrdered](items []T, value T) int {
	left, right := 0, len(items)-1
	for left <= right {
		mid1, mid2 := left+(right-left)/3, right-(right-left)/3

		if items[mid1] == value {
			return mid1
		}

		if items[mid2] == value {
			return mid2
		}

		if value < items[mid1] {
			right = mid1 - 1
		} else if value > items[mid2] {
			left = mid2 + 1
		} else {
			left, right = mid1+1, mid2-1
		}
	}

	return -1
}
