package sort

import "golang.org/x/exp/constraints"

func Bubble[T constraints.Ordered](values []T) []T {
	for i := 0; i < len(values); i++ {
		si := findSmallest(values, i)
		values[i], values[si] = values[si], values[i]
	}
	return values
}

func findSmallest[T constraints.Ordered](values []T, start int) int {
	si := start
	for i := start; i < len(values); i++ {
		if values[i] < values[si] {
			si = i
		}
	}
	return si
}
