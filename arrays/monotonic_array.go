package arrays

func IsMonotonic(array []int) bool {
	if len(array) < 2 {
		return true
	}

	prev := array[0]
	current := array[1]
	i := 2

	for current == prev && i < len(array) {
		prev = current
		current = array[i]
		i++
	}

	increasing := prev < current

	for i < len(array) {
		prev = current
		current = array[i]
		i++

		if current == prev {
			continue
		}

		stillSameDirection := increasing == (prev < current)
		if !stillSameDirection {
			return false
		}
	}

	return true
}
