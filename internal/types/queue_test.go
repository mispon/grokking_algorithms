package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Queue(t *testing.T) {
	q := NewQueue[int](10)
	require.Equal(t, 0, q.Len())

	t.Run("enqueue", func(t *testing.T) {
		for i := 0; i < 3; i++ {
			q.Enqueue(i + 1)
		}
		require.Equal(t, 3, q.Len())
	})

	t.Run("dequeue", func(t *testing.T) {
		for i := 0; i < 3; i++ {
			item, ok := q.Dequeue()
			require.Equal(t, i+1, item)
			require.True(t, ok)
		}

		item, ok := q.Dequeue()
		require.Equal(t, 0, item)
		require.False(t, ok)

		require.Equal(t, 0, q.Len())
	})
}

func Benchmark_Queue(b *testing.B) {
	q := NewQueue[int](b.N)

	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}

	for i := 0; i < b.N; i++ {
		q.Dequeue()
	}
}
