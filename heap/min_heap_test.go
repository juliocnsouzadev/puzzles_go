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

func TestMinHeap_swap(t *testing.T) {
	// given
	heap := NewMinHeap([]int{1, 2, 3, 4, 5, 6, 7})
	lastIdx := len(*heap) - 1

	// when
	heap.swap(lastIdx, lastIdx-1)

	// expected
	expectedHeap := NewMinHeap([]int{1, 2, 3, 4, 5, 7, 6})

	// then
	require.Equal(t, *heap, *expectedHeap)
}

func TestMinHeap_SiftDown_notParent(t *testing.T) {
	// given
	heap := NewMinHeap([]int{1, 2, 3, 4, 7, 5, 6})
	parentIdx := 6

	// when
	heap.siftDown(parentIdx, len(*heap)-1)

	// expected
	expectedHeap := heap //no changes expected

	// then
	require.Equal(t, *heap, *expectedHeap)
}

func TestMinHeap_SiftDown_outOfBounds(t *testing.T) {
	// given
	heap := NewMinHeap([]int{1, 2, 3, 4, 7, 5, 6})
	parentIdx := len(*heap)

	// when
	heap.siftDown(parentIdx, len(*heap)-1)

	// expected
	expectedHeap := heap //no changes expected

	// then
	require.Equal(t, *heap, *expectedHeap)
}

func TestMinHeap_SiftDown(t *testing.T) {
	// given
	heap := NewMinHeap([]int{1, 2, 7, 4, 3, 6, 5})
	childIdx := 6

	// when
	parentIdx := heap.getParentIndex(childIdx)
	heap.siftDown(parentIdx, len(*heap)-1)

	// expected
	expectedHeap := NewMinHeap([]int{1, 2, 5, 4, 3, 6, 7})

	// then
	require.Equal(t, *expectedHeap, *heap)
}

func TestMinHeap_SiftDown_explicity(t *testing.T) {
	// given
	heap := NewMinHeap([]int{1, 2, 15, 4, 3, 13, 12})

	// when
	parentIdx := 2
	leftChildIdx := 5
	rightChildIdx := 6
	heap.siftDown(parentIdx, len(*heap)-1)

	// expected
	expectedParent := 12
	expectedChildLeft := 13
	expectedChildRight := 15

	// then
	require.Equal(t, expectedParent, (*heap)[parentIdx])
	require.Equal(t, expectedChildLeft, (*heap)[leftChildIdx])
	require.Equal(t, expectedChildRight, (*heap)[rightChildIdx])
}

func TestMinHeap_Build(t *testing.T) {
	// given
	array := []int{48, 12, 24, 7, 8, -5, 24, 391, 24, 56, 2, 6, 8, 41}

	expected := []int{-5, 2, 6, 7, 8, 8, 24, 391, 24, 56, 12, 24, 48, 41}

	// when
	minHeap := NewMinHeap(array)

	// then
	require.NotNil(t, minHeap)
	require.True(t, isMinHeap(expected, 0), "expected %v should be a min heap", expected)
	require.True(t, isMinHeap(*minHeap, 0), "built heap %v should be a min heap", *minHeap)
}

func isMinHeap(heap []int, index int) bool {
	// If the node is a leaf, it's a valid min heap.
	if index >= len(heap)/2 {
		return true
	}

	leftChild, rightChild := getChildrenIndices(index)

	// Check the heap property.
	isHeap := hasHeapProperty(heap, index, leftChild, rightChild)
	if !isHeap {
		return false
	}

	// Recursively check the children.
	return isMinHeap(heap, leftChild) && isMinHeap(heap, rightChild)
}

func hasHeapProperty(heap []int, parent int, leftChild int, rightChild int) bool {
	if heap[parent] > heap[leftChild] || (rightChild < len(heap) && heap[parent] > heap[rightChild]) {
		//fmt.Printf("\nnot min heap because:\nparent:%v left:%v right:%v", heap[parent], heap[leftChild], heap[rightChild])
		return false
	}
	return true
}
