package heap

type MinHeap []int

func NewMinHeap(array []int) *MinHeap {
	// Do not edit the lines below.
	heap := MinHeap(array)
	ptr := &heap
	ptr.BuildHeap(array)
	return ptr
}

func (h *MinHeap) getParentIndex(childIdx int) int {
	return (childIdx - 1) / 2
}

func (h *MinHeap) BuildHeap(array []int) {

}

func (h *MinHeap) siftDown(parentIndex, endIndex int) {
	if parentIndex > endIndex || !h.isParent(parentIndex) {
		return
	}
	leftChildIdx, rightChildIdx := h.getChildrenIndices(parentIndex)

	parent := (*h)[parentIndex]
	leftChild := (*h)[leftChildIdx]
	rightChild := (*h)[rightChildIdx]

	if parent > leftChild {
		h.swap(parentIndex, leftChildIdx)
		leftChildIdx, rightChildIdx = h.getChildrenIndices(parentIndex)
		parent = (*h)[parentIndex]
		leftChild = (*h)[leftChildIdx]
		rightChild = (*h)[rightChildIdx]
	}

	if parent > rightChild {
		h.swap(parentIndex, rightChildIdx)
		leftChildIdx, rightChildIdx = h.getChildrenIndices(parentIndex)
		parent = (*h)[parentIndex]
		leftChild = (*h)[leftChildIdx]
		rightChild = (*h)[rightChildIdx]
	}

	if leftChild > rightChild {
		h.swap(leftChildIdx, rightChildIdx)
	}

}

func (h *MinHeap) swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MinHeap) getChildrenIndices(parentIdx int) (int, int) {
	doubleParentIdx := 2 * parentIdx
	return doubleParentIdx + 1, doubleParentIdx + 2
}

func (h *MinHeap) isParent(parentIdx int) bool {
	if r, l := getChildrenIndices(parentIdx); r < len(*h) || l < len(*h) {
		return true
	}
	return false
}

func (h *MinHeap) siftUp() {
	// Write your code here.
}

func (h MinHeap) Peek() int {
	// Write your code here.
	return -1
}

func (h *MinHeap) Remove() int {
	// Write your code here.
	return -1
}

func (h *MinHeap) Insert(value int) {
	// Write your code here.
}

func getChildrenIndices(parentIdx int) (int, int) {
	leftChild := 2*parentIdx + 1
	rightChild := 2*parentIdx + 2
	return leftChild, rightChild
}
