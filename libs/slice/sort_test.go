package slice

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Sort(t *testing.T) {
	t.Run("int slice", func(t *testing.T) {
		var (
			values   = []int{4, 8, 2, 1, 9, 7, 6, 3, 5}
			expected = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		)

		sorted := Sort(values)
		require.Equal(t, expected, sorted)
	})

	t.Run("string slice", func(t *testing.T) {
		var (
			values   = []string{"Banana", "Kiwi", "Elderberry", "Apple", "Plum", "Cherry"}
			expected = []string{"Apple", "Banana", "Cherry", "Elderberry", "Kiwi", "Plum"}
		)

		sorted := Sort(values)
		require.Equal(t, expected, sorted)
	})
}
