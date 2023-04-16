package sort

import (
	"math/rand"

	"golang.org/x/exp/constraints"
)

func Quick[T constraints.Ordered](values []T) []T {
	if len(values) < 2 {
		return values
	}

	var (
		left, right = 0, len(values) - 1
		pivot       = rand.Intn(right)
	)

	swap(values, pivot, right)

	for i := 0; i < len(values)-1; i++ {
		if values[i] < values[right] {
			swap(values, left, i)
			left++
		}
	}

	swap(values, left, right)

	Quick(values[:left])
	Quick(values[left+1:])

	return values
}

func swap[T any](values []T, i, j int) {
	values[i], values[j] = values[j], values[i]
}
