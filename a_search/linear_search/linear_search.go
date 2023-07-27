package linearsearch

func Search[T comparable](items []T, value T) int {
	for idx, item := range items {
		if item == value {
			return idx
		}
	}

	return -1
}
