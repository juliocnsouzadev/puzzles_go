package arrays

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestArrayOfProducts(t *testing.T) {
	// fixture
	cases := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{5, 1, 4, 2},
			expected: []int{8, 40, 10, 20},
		},
		{
			input:    []int{0, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			input:    []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: []int{362880, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}

	for _, testCase := range cases {
		// test
		result := ArrayOfProducts(testCase.input)

		// assertions
		require.Equal(t, testCase.expected, result)
	}
}
