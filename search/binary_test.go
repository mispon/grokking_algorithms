package search

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/exp/rand"
)

func Test_Binary(t *testing.T) {
	tests := []struct {
		values []int
		target int
		want   int
	}{
		{
			values: nil,
			target: 10,
			want:   -1,
		},
		{
			values: []int{1, 3, 4, 6, 8, 10, 11, 12, 15},
			target: 8,
			want:   4,
		},
	}

	t.Run("loop", func(t *testing.T) {
		for _, tc := range tests {
			t.Run("", func(t *testing.T) {
				index := Binary(tc.values, tc.target)
				require.Equal(t, tc.want, index)
			})
		}
	})

	t.Run("recursion", func(t *testing.T) {
		for _, tc := range tests {
			t.Run("", func(t *testing.T) {
				index := BinaryRec(tc.values, tc.target)
				require.Equal(t, tc.want, index)
			})
		}
	})
}

func Benchmark_Binary(b *testing.B) {
	const size = 1000
	var (
		target = rand.Intn(size)
		values = make([]int, 0, size)
	)
	for i := 0; i < size; i++ {
		values = append(values, i+1)
	}

	b.Run("loop", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Binary(values, target)
		}
	})

	b.Run("recursion", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			BinaryRec(values, target)
		}
	})
}
