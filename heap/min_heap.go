package heap

type MinHeap []int

func getParentIndex(childIdx int) int {
	return (childIdx - 1) / 2
}
