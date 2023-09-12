package arrays

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSpiralTraverse(t *testing.T) {
	// fixture
	cases := []struct {
		input    [][]int
		expected []int
	}{
		{
			input: [][]int{
				{1, 2, 3, 4},
				{12, 13, 14, 5},
				{11, 16, 15, 6},
				{10, 9, 8, 7},
			},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
		},
	}
	// test
	for _, testCase := range cases {
		result := SpiralTraverse(testCase.input)

		require.Equal(t, testCase.expected, result)
	}
}
