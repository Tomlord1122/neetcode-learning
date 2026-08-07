func maxAreaOfIsland(grid [][]int) int {
    ROWS, COLS := len(grid), len(grid[0])
	maxArea := 0
	directions := [][]int{{0,1},{0,-1},{1,0},{-1,0}}	
	cur := 0

	var dfs func(r, c int)
	dfs = func(r, c int){
		if r < 0 || c < 0 || r >= ROWS || c >= COLS || grid[r][c] == 0{
			return
		}
		cur++
		grid[r][c] = 0
		for _, d := range directions{
			dfs(r+d[0], c+d[1])
		}
	}


	for r := 0; r < ROWS; r++{
		for c := 0; c <COLS; c++{
			if grid[r][c] == 1{
				cur = 0
				dfs(r,c)
				maxArea = max(maxArea, cur)
			}
		}
	}
	return maxArea
}
