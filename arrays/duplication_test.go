package arrays

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFirstDuplicateValue(t *testing.T) {
	// fixture
	cases := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{2, 1, 5, 2, 3, 3, 4},
			expected: 2,
		},
		{
			input:    []int{2, 1, 5, 3, 3, 2, 4},
			expected: 3,
		},
	}

	for _, testCase := range cases {
		// test
		result := FirstDuplicateValue(testCase.input)

		// assertions
		require.Equal(t, testCase.expected, result)
	}
}
