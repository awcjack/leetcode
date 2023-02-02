func minimumTotal(triangle [][]int) int {
	for i := len(triangle) - 1; i >= 0; i-- {
		for j := 0; j < len(triangle[i])-1; j++ {
			if triangle[i][j] > triangle[i][j+1] {
				if i > 0 {
					triangle[i-1][j] += triangle[i][j+1]
				} else {
					return triangle[i][j+1]
				}
			} else {
				if i > 0 {
					triangle[i-1][j] += triangle[i][j]
				} else {
					return triangle[i][j]
				}
			}
		}
	}
	return triangle[0][0]
}
