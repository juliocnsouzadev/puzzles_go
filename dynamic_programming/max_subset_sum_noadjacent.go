package dynamic_programming

func MaxSubsetSumNoAdjacent(array []int) int {
	if len(array) == 0 {
		return 0
	}
	if len(array) == 1 {
		return array[0]
	}
	if len(array) == 2 {
		return Max(array[0], array[1])
	}

	maxSums := make([]int, len(array))
	maxSums[0] = array[0]
	maxSums[1] = Max(array[0], array[1])

	maxSum := Max(maxSums[0], maxSums[1])
	for i := 2; i < len(array); i++ {
		maxSums[i] = Max(maxSums[i-1], maxSums[i-2]+array[i])
		maxSum = Max(maxSum, maxSums[i])
	}
	return maxSum
}
