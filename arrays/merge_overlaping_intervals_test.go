package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeOverlappingIntervals(t *testing.T) {
	// fixture
	cases := []struct {
		input    [][]int
		expected [][]int
	}{
		{
			input:    [][]int{{1, 2}, {3, 5}, {4, 7}, {6, 8}, {9, 10}},
			expected: [][]int{{1, 2}, {3, 8}, {9, 10}},
		},
	}

	for _, tt := range cases {
		// test
		result := MergeOverlappingIntervals(tt.input)

		// assertions
		assert.Equal(t, tt.expected, result)
	}
}
