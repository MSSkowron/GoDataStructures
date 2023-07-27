package interpolationsearch

func Search(items []int, value int) int {
	low := 0
	high := len(items) - 1

	for low <= high && value >= items[low] && value <= items[high] {
		if low == high {
			if items[low] == value {
				return low
			}
			return -1
		}

		pos := low + ((value-items[low])*(high-low))/(items[high]-items[low])

		if items[pos] == value {
			return pos
		}

		if items[pos] < value {
			low = pos + 1
		} else {
			high = pos - 1
		}
	}

	return -1
}
