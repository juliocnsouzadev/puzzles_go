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

func TestMinHeap_childrenIndices(t *testing.T) {
	// given
	heap := NewMinHeap([]int{1, 2, 3, 4, 5, 6, 7})
	parentIdx := 2

	// when
	leftChildIdx, rightChildIdx := heap.getChildrenIndices(parentIdx)

	// expected
	expectedLeftChildIdx := 5
	expectedRightChildIdx := 6

	// then
	require.Equal(t, expectedLeftChildIdx, leftChildIdx)
	require.Equal(t, expectedRightChildIdx, rightChildIdx)
}

func TestMinHeap_isParent(t *testing.T) {
	// given
	heap := NewMinHeap([]int{1, 2, 3, 4, 7, 5, 6})
	parentIdx := 2

	// when
	isParent := heap.isParent(parentIdx)

	// then
	require.Equal(t, true, isParent)
}

func TestMinHeap_isNotParent(t *testing.T) {
	// given
	heap := NewMinHeap([]int{1, 2, 3, 4, 7, 5, 6})
	parentIdx := 5

	// when
	isParent := heap.isParent(parentIdx)

	// then
	require.Equal(t, false, isParent)
}
