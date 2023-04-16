package search

import (
	"golang.org/x/exp/constraints"
)

func Binary[T constraints.Ordered](values []T, target T) int {
	if len(values) == 0 {
		return -1
	}

	var (
		low  = 0
		high = len(values) - 1
	)

	for low <= high {
		mid := (low + high) / 2

		if values[mid] == target {
			return mid
		}

		if values[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func BinaryRec[T constraints.Ordered](values []T, target T) int {
	if len(values) == 0 {
		return -1
	}

	var (
		low  = 0
		high = len(values) - 1
	)

	mid := (low + high) / 2

	if values[mid] == target {
		return mid
	}

	if values[mid] > target {
		return BinaryRec(values[:mid-1], target)
	} else {
		return BinaryRec(values[mid+1:], target)
	}
}
