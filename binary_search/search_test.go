package binary_search

import (
	"testing"

	"github.com/mispon/grokking_algorithms/sort/selection"
	"github.com/stretchr/testify/require"
)

func Test_search(t *testing.T) {
	testCases := []struct {
		values   []int
		item     int
		expIndex int
		expBool  bool
	}{
		{
			values:   []int{3, 5, 1, 9, 7},
			item:     3,
			expIndex: 1,
			expBool:  true,
		},
		{
			values:   []int{3, 5, 1, 9, 7},
			item:     6,
			expIndex: 0,
			expBool:  false,
		},
		{
			values:   []int{1, 2, 3, 4, 5},
			item:     3,
			expIndex: 2,
			expBool:  true,
		},
		{
			values:   []int{1, 1, 3, 4, 5},
			item:     1,
			expIndex: 0,
			expBool:  true,
		},
		{
			values:   []int{1, 1, 1, 4, 5},
			item:     1,
			expIndex: 2,
			expBool:  true,
		},
		{
			values:   []int{4, 1, 5, 1, 1, 8},
			item:     1,
			expIndex: 2,
			expBool:  true,
		},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			values := selection.Sort(tc.values)

			idx, ok := search(values, tc.item)

			require.Equal(t, tc.expIndex, idx)
			require.Equal(t, tc.expBool, ok)
		})
	}
}

func Test_rsearch(t *testing.T) {
	testCases := []struct {
		values   []int
		item     int
		expIndex int
		expBool  bool
	}{
		{
			values:   []int{3, 5, 1, 9, 7},
			item:     3,
			expIndex: 1,
			expBool:  true,
		},
		{
			values:   []int{3, 5, 1, 9, 7},
			item:     6,
			expIndex: 0,
			expBool:  false,
		},
		{
			values:   []int{1, 2, 3, 4, 5},
			item:     3,
			expIndex: 2,
			expBool:  true,
		},
		{
			values:   []int{1, 1, 3, 4, 5},
			item:     1,
			expIndex: 0,
			expBool:  true,
		},
		{
			values:   []int{1, 1, 1, 4, 5},
			item:     1,
			expIndex: 2,
			expBool:  true,
		},
		{
			values:   []int{4, 1, 5, 1, 1, 8},
			item:     1,
			expIndex: 1,
			expBool:  true,
		},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			values := selection.Sort(tc.values)

			idx, ok := rsearch(values, tc.item, 0, len(values))

			require.Equal(t, tc.expIndex, idx)
			require.Equal(t, tc.expBool, ok)
		})
	}
}
