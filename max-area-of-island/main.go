func blockSurrounding(grid [][]int, x int, y int) int {
	counter := 1
	// up
	if x > 0 {
		if grid[x-1][y] == 1 {
			grid[x-1][y] = 0
			counter += blockSurrounding(grid, x-1, y)
		}
	}
	if y < len(grid[0])-1 {
		if grid[x][y+1] == 1 {
			grid[x][y+1] = 0
			counter += blockSurrounding(grid, x, y+1)
		}
	}
	if x < len(grid)-1 {
		if grid[x+1][y] == 1 {
			grid[x+1][y] = 0
			counter += blockSurrounding(grid, x+1, y)
		}
	}
	if y > 0 {
		if grid[x][y-1] == 1 {
			grid[x][y-1] = 0
			counter += blockSurrounding(grid, x, y-1)
		}
	}
	return counter
}

func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	if len(grid[0]) == 0 {
		return 0
	}

	max := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				grid[i][j] = 0
				currentArea := blockSurrounding(grid, i, j)
				if max < currentArea {
					max = currentArea
				}
			}
		}
	}
	return max
}
