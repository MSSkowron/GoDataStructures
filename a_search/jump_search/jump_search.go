package jumpsearch

import (
	"math"

	"golang.org/x/exp/constraints"
)

type ComparableOrdered interface {
	comparable
	constraints.Ordered
}

func Search[T ComparableOrdered](items []T, value T) int {
	n := len(items)
	step := int(math.Sqrt(float64(n)))
	prev := 0

	if step == 0 {
		return -1
	}

	for items[min(step, n)-1] < value {
		prev = step
		step += int(math.Sqrt(float64(n)))
		if prev >= n {
			return -1
		}
	}

	for items[prev] < value {
		prev++
		if prev == min(step, n) {
			return -1
		}
	}

	if items[prev] == value {
		return prev
	}

	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
