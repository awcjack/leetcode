// so on9
// 0 ms	2 MB
func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	direction := "R"
	row := 0
	column := 0
	upBoundary := n
	rightBoundary := n
	leftBoundary := n
	downBoundary := n
	counter := 0
	for i := 1; i <= n*n; i++ {
		matrix[row][column] = i
		counter++

		if direction == "R" && counter < rightBoundary {
			column++
		} else if direction == "R" {
			direction = "D"
			counter = 0
			row++
			upBoundary--
			downBoundary--
		} else if direction == "D" && counter < downBoundary {
			row++
		} else if direction == "D" {
			column--
			direction = "L"
			counter = 0
			rightBoundary--
			leftBoundary--
		} else if direction == "L" && counter < leftBoundary {
			column--
		} else if direction == "L" {
			row--
			direction = "W"
			counter = 0
			upBoundary--
			downBoundary--
		} else if direction == "W" && counter < upBoundary {
			row--
		} else if direction == "W" {
			column++
			direction = "R"
			counter = 0
			leftBoundary--
			rightBoundary--
		}
	}
	return matrix
}
