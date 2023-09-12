package arrays

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMoveElementToEnd(t *testing.T) {
	// fixture
	cases := []struct {
		input    []int
		toMove   int
		expected []int
	}{
		{
			input:    []int{2, 1, 2, 2, 2, 3, 4, 2},
			toMove:   2,
			expected: []int{1, 3, 4, 2, 2, 2, 2, 2},
		},
	}

	// test
	for _, testCase := range cases {
		result := MoveElementToEnd(testCase.input, testCase.toMove)
		require.Equal(t, testCase.expected, result)
	}
}
