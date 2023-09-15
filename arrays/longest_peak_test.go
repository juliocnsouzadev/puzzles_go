package arrays

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLongestPeak(t *testing.T) {
	// fixture
	cases := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{1, 2, 3, 3, 4, 0, 10, 6, 5, -1, -3, 2, 3},
			expected: 6,
		},
	}

	for _, testCase := range cases {

		// test
		result := LongestPeak(testCase.input)

		// assertions
		require.Equal(t, testCase.expected, result)
	}
}
