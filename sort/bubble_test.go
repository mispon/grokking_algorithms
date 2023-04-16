package sort

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Bubble(t *testing.T) {
	tests := []struct {
		values []int
		want   []int
	}{
		{
			values: nil,
			want:   nil,
		},
		{
			values: []int{},
			want:   []int{},
		},
		{
			values: []int{3, 5, 6, 1, 4, 2},
			want:   []int{1, 2, 3, 4, 5, 6},
		},
	}

	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			require.Equal(t, tc.want, Bubble(tc.values))
		})
	}
}

func Benchmark_Bubble(b *testing.B) {
	const size = 1000
	values := make([]int, 0, size)
	for i := 0; i < size; i++ {
		values = append(values, rand.Intn(size))
	}

	for i := 0; i < b.N; i++ {
		Bubble(values)
	}
}
