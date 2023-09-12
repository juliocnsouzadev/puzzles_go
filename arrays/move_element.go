package arrays

func MoveElementToEnd(array []int, toMove int) []int {
	result := make([]int, len(array))
	left := 0
	right := len(array) - 1
	for _, value := range array {
		if value == toMove {
			result[right] = value
			right--
			continue
		}
		result[left] = value
		left++
	}
	return result
}
