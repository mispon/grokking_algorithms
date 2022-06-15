package slice

import "golang.org/x/exp/constraints"

// Sort is func that sorts all elements in slice by desc
func Sort[T constraints.Ordered](values []T) []T {
	for i := 0; i < len(values)-1; i++ {
		for j := i + 1; j < len(values); j++ {
			if values[i] > values[j] {
				values[i], values[j] = values[j], values[i]
			}
		}
	}

	return values
}
