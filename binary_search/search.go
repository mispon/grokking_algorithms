package binary_search

func search(values []int, v int) (int, bool) {
	var (
		low  = 0
		high = len(values) - 1
	)

	for low <= high {
		mid := (low + high) / 2
		guess := values[mid]

		if guess == v {
			return mid, true
		}

		if guess < v {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return 0, false
}

func rsearch(values []int, v int, low, high int) (int, bool) {
	if low > high {
		return 0, false
	}

	mid := (low + high) / 2
	guess := values[mid]

	if guess == v {
		return mid, true
	}

	if guess < v {
		low = mid + 1
	} else {
		high = mid - 1
	}

	return rsearch(values, v, low, high)
}
