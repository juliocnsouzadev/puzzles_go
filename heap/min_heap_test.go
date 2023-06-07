package heap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMinHeap_GetParentIndexFromLeftChild(t *testing.T) {
	// given
	childIdx := 5

	// when
	parentIdx := getParentIndex(childIdx)

	// then
	require.Equal(t, 2, parentIdx)
}

func TestMinHeap_GetParentIndexFromRightChild(t *testing.T) {
	// given
	childIdx := 6

	// when
	parentIdx := getParentIndex(childIdx)

	// then
	require.Equal(t, 2, parentIdx)
}
