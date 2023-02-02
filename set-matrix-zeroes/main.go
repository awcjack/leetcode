func setZeroes(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	columnToBeZero := make([]bool, len(matrix[0]), len(matrix[0]))
	rowToBeZero := make([]bool, len(matrix), len(matrix))
	for rowIndex := 0; rowIndex < len(matrix); rowIndex++ {
		for columnIndex := 0; columnIndex < len(matrix[rowIndex]); columnIndex++ {
			if matrix[rowIndex][columnIndex] == 0 {
				columnToBeZero[columnIndex] = true
				rowToBeZero[rowIndex] = true
			}
		}
	}

	for rowIndex := 0; rowIndex < len(matrix); rowIndex++ {
		for columnIndex := 0; columnIndex < len(matrix[rowIndex]); columnIndex++ {
			if rowToBeZero[rowIndex] || columnToBeZero[columnIndex] {
				matrix[rowIndex][columnIndex] = 0
			}
		}
	}
}
