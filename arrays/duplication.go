package arrays

func FirstDuplicateValue(array []int) int {
	itemsCounter := map[int]int{}

	for i := 0; i < len(array); i++ {
		currentItem := array[i]
		_, ok := itemsCounter[currentItem]
		if ok {
			return currentItem
		}

		itemsCounter[currentItem] = 1
	}
	return -1
}
