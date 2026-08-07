func maxAreaOfIsland(grid [][]int) int {
    ROWS, COLS := len(grid), len(grid[0])
	dirs := [][]int{{0,1}, {0,-1}, {1,0}, {-1,0}}
	res := 0

	var dfs func(r, c int) int
	dfs = func(r, c int) int{
		if min(r, c) < 0 || r == ROWS || c == COLS || grid[r][c] == 0{
			return 0
		}
		grid[r][c] = 0
		cur := 1
		for _, d := range dirs{
			cur += dfs(r+d[0], c+d[1])
		}
		return cur
	}

	for r := 0; r < ROWS; r++{
		for c := 0; c < COLS; c++{
			if grid[r][c] == 1{
				res = max(res, dfs(r, c))
			}
		}
	}
	return res
}
