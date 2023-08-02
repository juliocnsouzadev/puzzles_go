package arrays

func TransposeMatrix(matrix [][]int) [][]int {
	rows := len(matrix)

	if matrix == nil || rows == 0 {
		return nil
	}

	cols := len(matrix[0])

	result := make([][]int, cols)
	for i, row := range matrix {
		for j, value := range row {
			if result[j] == nil {
				result[j] = make([]int, rows)
			}
			result[j][i] = value
		}
	}
	return result
}
