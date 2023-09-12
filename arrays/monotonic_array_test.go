package arrays

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIsMonotonic(t *testing.T) {
	// fixture
	cases := []struct {
		input    []int
		expected bool
	}{
		{
			input:    []int{-1, -5, -10, -1100, -1100, -1101, -1102, -9001},
			expected: true,
		},
	}
	// test
	for _, testCase := range cases {
		result := IsMonotonic(testCase.input)

		require.Equal(t, testCase.expected, result)
	}
}
