func numIslands(grid [][]byte) int {
    count := 0
	directions := [][]int{{0,1}, {0,-1}, {1,0}, {-1,0}}
	rows, cols := len(grid), len(grid[0])
	var dfs func(r, c int)
	dfs = func(r, c int){
		if r < 0 || c < 0 || r >= rows || c>= cols || grid[r][c] == '0'{
			return
		}
		grid[r][c] = '0'
		for _, d := range directions{
			dfs(r+d[0], c+d[1])
		}
	}

	for r := 0; r < rows; r++{
		for c := 0; c < cols; c++{
			if grid[r][c] == '1'{
				dfs(r, c)
				count++
			}
		}
	}
	return count
}
