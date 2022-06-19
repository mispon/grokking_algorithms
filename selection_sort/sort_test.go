package selection_sort

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_SortInt(t *testing.T) {
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
			sorted := SortInt(tc.values)
			require.Equal(t, tc.expected, sorted)
		})
	}
}
