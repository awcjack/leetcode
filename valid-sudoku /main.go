func isValidSudoku(board [][]byte) bool {
	hashmap := make(map[string]bool)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != 46 { // "."
				if hashmap[fmt.Sprintf("row%d%d", i, board[i][j])] || hashmap[fmt.Sprintf("col%d%d", j, board[i][j])] || hashmap[fmt.Sprintf("box%d%d%d", i/3, j/3, board[i][j])] {
					return false
				}
				hashmap[fmt.Sprintf("row%d%d", i, board[i][j])] = true
				hashmap[fmt.Sprintf("col%d%d", j, board[i][j])] = true
				hashmap[fmt.Sprintf("box%d%d%d", i/3, j/3, board[i][j])] = true
			}
		}
	}
	return true
}
