package selection

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_Sort(t *testing.T) {
	testCases := []struct {
		values   []int
		expected []int
	}{
		{
			values:   []int{2, 5, 1, 8, 3, 9, 0},
			expected: []int{0, 1, 2, 3, 5, 8, 9},
		},
		{
			values:   []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			values:   []int{7, 1, 3, 4, 3, 1, 6, 3, 5, 0},
			expected: []int{0, 1, 1, 3, 3, 3, 4, 5, 6, 7},
		},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			sorted := Sort(tc.values)
			require.Equal(t, tc.expected, sorted)
		})
	}
}

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
