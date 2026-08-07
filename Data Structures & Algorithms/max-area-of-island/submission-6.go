func maxAreaOfIsland(grid [][]int) int {
    maxArea := 0
	cur := 0
	dirs := [][]int{{0,1}, {0,-1}, {1,0}, {-1,0}}
	ROW, COL := len(grid), len(grid[0])
	
	var dfs func(r, c int)
	dfs = func(r, c int){
		if min(r, c) < 0 || r == ROW || c == COL || grid[r][c] == 0{
			return
		}
		grid[r][c] = 0
		cur++
		for _, d := range dirs{
			dfs(r+d[0], c+d[1])
		}
	}

	for r := 0; r < ROW; r++{
		for c := 0; c < COL; c++{
			if grid[r][c] == 1{
				dfs(r, c)
				maxArea = max(maxArea, cur)
				cur = 0
			}
		}
	}
	return maxArea
}
