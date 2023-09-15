package greedy

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestClassPhotos(t *testing.T) {
	// fixture
	cases := []struct {
		redShirts  []int
		blueShirts []int
		expected   bool
	}{
		{
			redShirts:  []int{5, 8, 1, 3, 4},
			blueShirts: []int{6, 9, 2, 4, 5},
			expected:   true,
		},
		{
			redShirts:  []int{6, 9, 2, 4, 5},
			blueShirts: []int{5, 8, 1, 3, 4},
			expected:   true,
		},
		{
			redShirts:  []int{5, 6},
			blueShirts: []int{5, 4},
			expected:   true,
		},
		{
			redShirts:  []int{3, 5, 6, 8, 2},
			blueShirts: []int{2, 4, 7, 5, 1},
			expected:   true,
		},
	}

	for i, testCase := range cases {
		// test
		result := ClassPhotos(testCase.redShirts, testCase.blueShirts)

		// assertions
		require.Equal(t, testCase.expected, result, fmt.Sprintf("test case: %d", i))
	}
}
