func getOuterSum(n int, layer int) int {
	if layer == 0 {
		return 0
	}
	return ((n - (layer-1)*2) + (n - (layer-1)*2) - 1 + (n - (layer-1)*2) - 1 + (n - (layer-1)*2) - 2) + getOuterSum(n, layer-1)
}

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	for i := 0; i < int(math.Ceil(float64(n)/2)); i++ {
		for j := 0; j < n-i*2; j++ {
			matrix[i][j+i] = getOuterSum(n, i) + j + 1
		}
		for j := 0; j < n-i*2-1; j++ {
			matrix[i+1+j][n-i-1] = getOuterSum(n, i) + (n - i*2) + j + 1
		}
		for j := 0; j < n-i*2-1; j++ {
			matrix[n-1-i][n-1-i-1-j] = getOuterSum(n, i) + (n - i*2) + (n - i*2 - 1) + j + 1
		}
		for j := 0; j < n-i*2-2; j++ {
			matrix[n-1-i-1-j][i] = getOuterSum(n, i) + (n - i*2) + (n - i*2 - 1) + (n - i*2 - 1) + j + 1
		}
	}
	return matrix
}
