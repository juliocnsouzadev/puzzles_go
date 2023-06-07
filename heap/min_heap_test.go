package heap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMinHeap_GetParentIndexFromLeftChild(t *testing.T) {
	// given
	heap := NewMinHeap([]int{1, 2, 3, 4, 5, 6, 7})
	childIdx := 5

	// when
	parentIdx := heap.getParentIndex(childIdx)

	// then
	require.Equal(t, 2, parentIdx)
}

func TestMinHeap_GetParentIndexFromRightChild(t *testing.T) {
	// given
	heap := NewMinHeap([]int{1, 2, 3, 4, 5, 6, 7})
	childIdx := 6

	// when
	parentIdx := heap.getParentIndex(childIdx)

	// then
	require.Equal(t, 2, parentIdx)
}
