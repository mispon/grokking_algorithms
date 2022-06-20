package quick_sort

import (
	"math/rand"
	"testing"
	"time"
)

func Benchmark_sort(b *testing.B) {
	rand.Seed(time.Now().Unix())

	numbers := make([]int, 100)
	for i := 0; i < 100; i++ {
		numbers[i] = rand.Intn(256)
	}

	for i := 0; i < b.N; i++ {
		Sort(numbers)
	}
}
