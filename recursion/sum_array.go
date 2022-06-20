package recursion

func sum(values []int) int {
	if len(values) == 0 {
		return 0
	}

	if len(values) == 1 {
		return values[0]
	}

	return values[0] + sum(values[1:])
}
