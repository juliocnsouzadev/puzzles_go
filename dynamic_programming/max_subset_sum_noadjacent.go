package dynamic_programming

import "puzzles/util"

func MaxSubsetSumNoAdjacent(array []int) int {
	if len(array) == 0 {
		return 0
	}
	if len(array) == 1 {
		return array[0]
	}
	if len(array) == 2 {
		return util.Max(array[0], array[1])
	}

	maxSums := make([]int, len(array))
	maxSums[0] = array[0]
	maxSums[1] = util.Max(array[0], array[1])

	maxSum := util.Max(maxSums[0], maxSums[1])
	for i := 2; i < len(array); i++ {
		maxSums[i] = util.Max(maxSums[i-1], maxSums[i-2]+array[i])
		maxSum = util.Max(maxSum, maxSums[i])
	}
	return maxSum
}
