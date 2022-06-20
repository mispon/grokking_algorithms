package recursion

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_sum(t *testing.T) {
	testCases := []struct {
		values   []int
		expected int
	}{
		{
			values:   []int{1, 2, 3, 4, 5},
			expected: 15,
		},
		{
			values:   []int{-5, 5},
			expected: 0,
		},
		{
			values:   []int{},
			expected: 0,
		},
		{
			values:   []int{7},
			expected: 7,
		},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			require.Equal(t, tc.expected, sum(tc.values))
		})
	}
}
