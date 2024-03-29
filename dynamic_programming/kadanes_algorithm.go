package dynamic_programming

import "puzzles/util"

func KadanesAlgorithm(array []int) int {
	if len(array) == 0 {
		return 0
	}

	greatestSum := array[0]
	currentSum := greatestSum

	for i := 1; i < len(array); i++ {
		currentValue := array[i]
		currentSum = util.Max(currentValue, currentSum+array[i])
		greatestSum = util.Max(greatestSum, currentSum)
	}
	return greatestSum
}
