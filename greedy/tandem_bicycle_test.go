package greedy

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTandemBicycle(t *testing.T) {
	// fixture
	cases := []struct {
		redShirtSpeeds  []int
		blueShirtSpeeds []int
		fastest         bool
		expected        int
	}{
		{
			redShirtSpeeds:  []int{5, 5, 3, 9, 2},
			blueShirtSpeeds: []int{3, 6, 7, 2, 1},
			fastest:         true,
			expected:        32,
		},
		{
			redShirtSpeeds:  []int{5, 5, 3, 9, 2},
			blueShirtSpeeds: []int{3, 6, 7, 2, 1},
			fastest:         false,
			expected:        25,
		},
	}

	for _, testCase := range cases {
		// test
		result := TandemBicycle(testCase.redShirtSpeeds, testCase.blueShirtSpeeds, testCase.fastest)

		// assertions
		require.Equal(t, testCase.expected, result)
	}
}
