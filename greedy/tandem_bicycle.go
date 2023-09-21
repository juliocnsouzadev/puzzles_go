package greedy

import "sort"

func TandemBicycle(redShirtSpeeds []int, blueShirtSpeeds []int, fastest bool) int {
	sort.Ints(redShirtSpeeds)
	sort.Ints(blueShirtSpeeds)
	sum := 0

	shirtsLength := len(redShirtSpeeds)
	for redIdx := 0; redIdx < shirtsLength; redIdx++ {
		blueIdx := redIdx
		if fastest {
			blueIdx = shirtsLength - 1 - redIdx
		}

		faster := max(redShirtSpeeds[redIdx], blueShirtSpeeds[blueIdx])
		sum += faster

	}
	return sum
}
