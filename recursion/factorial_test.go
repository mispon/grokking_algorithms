package recursion

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_factorial(t *testing.T) {
	testCases := []struct {
		n   int
		exp int
	}{
		{n: 1, exp: 1},
		{n: 2, exp: 2},
		{n: 3, exp: 6},
		{n: 4, exp: 24},
		{n: 5, exp: 120},
		{n: 6, exp: 720},
		{n: 7, exp: 5040},
		{n: 8, exp: 40320},
		{n: 9, exp: 362880},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%d! is %d", tc.n, tc.exp), func(t *testing.T) {
			require.Equal(t, tc.exp, factorial(tc.n))
		})
	}
}
