package arrays

func LongestPeak(array []int) int {
	if len(array) < 3 {
		return 0
	}
	longestPeak := 0

	var peakItems []int
	increasing := false

	prev := array[0]
	peakItems = append(peakItems, prev)

	for i := 1; i < len(array); i++ {
		var current = array[i]

		if prev < current {
			if !increasing && len(array) > 0 {
				peakItems = keepOnlyLastItem(peakItems)
			}
			increasing = true
			peakItems = append(peakItems, i)
		}

		if prev > current && len(peakItems) > 1 {
			increasing = false
			peakItems = append(peakItems, i)
			longestPeak = getLongest(peakItems, longestPeak)
		}

		if prev == current {
			peakItems = keepOnlyLastItem(peakItems)
		}

		prev = current

	}

	return longestPeak
}

func keepOnlyLastItem(items []int) []int {
	last := items[len(items)-1]
	result := []int{last}
	return result
}

func getLongest(items []int, longest int) int {
	if longest < len(items) {
		longest = len(items)
	}
	return longest
}
