func maxAreaOfIsland(grid [][]int) int {
	res := 0
	directions := [][]int{{0,1}, {0,-1}, {1,0}, {-1,0}}
	row, col := len(grid), len(grid[0])
	var dfs func(r, c int) int
	dfs = func(r, c int) int{
		count := 0
		if r < 0 || c < 0 || r >= row || c >= col || grid[r][c] == 0{
			return 0
		}
		count++
		grid[r][c] = 0
		for _, dir := range directions{
			count += dfs(r+dir[0], c+dir[1])
		}
		return count
	}

	for r := 0; r < row; r++{
		for c := 0; c < col; c++{
			if grid[r][c] == 1{
				res = max(res, dfs(r, c))
			}
		}
	}
	return res
}
