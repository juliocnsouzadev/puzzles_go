package arrays

import (
	"fmt"
)

/*
Takes in a non-empty array of integers that are sorted
in ascending order and returns a new array of the same length with the squares
of the original integers also sorted in ascending order.
-> Time O(n), Space O(n)
*/
func SquareSortedArray(arrayIn []int) []int {
	squaredSortedArray := make([]int, len(arrayIn))
	left := 0
	right := len(arrayIn) - 1
	for i := len(arrayIn) - 1; i >= 0; i-- {
		smaller := arrayIn[left]
		larger := arrayIn[right]

		fmt.Printf("left: %d, right: %d, i:%d\n", smaller, larger, i)

		if larger > smaller {
			squaredSortedArray[i] = larger * larger
			right--

		} else {
			squaredSortedArray[i] = smaller * smaller
			left++
		}
	}
	return squaredSortedArray

}
