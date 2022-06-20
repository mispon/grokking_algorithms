package quick

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Sort(values []int) []int {
	if len(values) < 2 {
		return values
	}

	var (
		left  = 0
		right = len(values) - 1
		pivot = rand.Int() % len(values)
	)
	
	values[pivot], values[right] = values[right], values[pivot]

	for i := range values {
		if values[i] < values[right] {
			values[left], values[i] = values[i], values[left]
			left++
		}
	}

	values[left], values[right] = values[right], values[left]

	Sort(values[:left])
	Sort(values[left+1:])

	return values
}
