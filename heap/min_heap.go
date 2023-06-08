package heap

type MinHeap []int

func (h *MinHeap) getParentIndex(childIdx int) int {
	return (childIdx - 1) / 2
}

func (h *MinHeap) swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MinHeap) getChildrenIndices(parentIdx int) (int, int) {
	doubleParentIdx := 2 * parentIdx
	return doubleParentIdx + 1, doubleParentIdx + 2
}

func (h *MinHeap) isParent(parentIdx int) bool {
	if r, l := getChildrenIndices(parentIdx); r < len(*h) && l < len(*h) {
		return true
	}
	return false
}

func getChildrenIndices(parentIdx int) (int, int) {
	leftChild := 2*parentIdx + 1
	rightChild := 2*parentIdx + 2
	return leftChild, rightChild
}

func NewMinHeap(array []int) *MinHeap {
	heap := MinHeap(array)
	ptr := &heap
	ptr.BuildHeap(array)
	return ptr
}

func (h *MinHeap) BuildHeap(array []int) {
	*h = array
	firstParent := h.getParentIndex(len(array) - 1)
	for i := firstParent; i >= 0; i-- {
		h.siftDown(i, len(array)-1)
	}
}

func (h *MinHeap) siftDown(currentIndex, endIndex int) {

	leftChild, rightChild := getChildrenIndices(currentIndex)

	// As long as we haven't reached the end of the heap
	for leftChild <= endIndex {

		// If the right child exists and is smaller than the left child, go to the right child
		if rightChild <= endIndex && (*h)[leftChild] > (*h)[rightChild] {
			leftChild++
			rightChild++
		}

		// If the child node is less than the current node, swap them and move to the child node
		if (*h)[leftChild] < (*h)[currentIndex] {
			h.swap(currentIndex, leftChild)

			currentIndex = leftChild
			leftChild, rightChild = getChildrenIndices(currentIndex)
		} else {

			// If the child node is not less than the current node, the heap property is maintained and we can stop
			return
		}
	}
}
func (h *MinHeap) siftUp() {
	currentIndex := len(*h) - 1
	parentIndex := h.getParentIndex(currentIndex)

	for currentIndex > 0 && (*h)[currentIndex] < (*h)[parentIndex] {
		h.swap(currentIndex, parentIndex)

		currentIndex = parentIndex
		parentIndex = h.getParentIndex(currentIndex)
	}
}

func (h MinHeap) Peek() int {
	return h[0]
}

func (h *MinHeap) Remove() int {
	removedValue := (*h)[0]

	// Move the last node to the root
	(*h)[0] = (*h)[len(*h)-1]

	// Resize the slice to remove the last node
	*h = (*h)[:len(*h)-1]

	// Sift down the root node to maintain the heap property
	h.siftDown(0, len(*h)-1)

	// Return the removed root node
	return removedValue
}

func (h *MinHeap) Insert(value int) {
	*h = append(*h, value)

	h.siftUp()
}
