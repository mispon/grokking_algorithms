package recursion

// sum all values in slice recursively
func sum(values []int) int {
	if len(values) == 0 {
		return 0
	}

	if len(values) == 1 {
		return values[0]
	}

	return values[0] + sum(values[1:])
}

// count all values in slice recursively
func count(values []int) int {
	if len(values) == 0 {
		return 0
	}

	return 1 + count(values[1:])
}

// max finds the maximum value in slice recursively
func max(values []int) int {
	if len(values) == 0 {
		return 0
	}

	if len(values) == 1 {
		return values[0]
	}

	var (
		a = values[0]
		b = max(values[1:])
	)

	if a > b {
		return a
	} else {
		return b
	}
}
