package greedy

import (
	"sort"
)

func MinimumWaitingTime(tasks []int) int {
	sort.Ints(tasks)
	waitTime := 0
	totalWaitTime := 0
	for i := 0; i < len(tasks); i++ {
		if i > 0 {
			waitTime += tasks[i-1]
		}
		totalWaitTime += waitTime
	}
	return totalWaitTime
}
