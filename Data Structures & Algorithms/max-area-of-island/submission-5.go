func maxAreaOfIsland(grid [][]int) int {
    ROW, COL := len(grid), len(grid[0])
	dir := [][]int{{0,1}, {0,-1}, {-1,0}, {1,0}}

	cur := 0

	var dfs func(r, c int)
	dfs = func(r, c int){
		if min(r, c) < 0 || r == ROW || c == COL || grid[r][c] == 0{
			return
		}
		grid[r][c] = 0
		cur++
		for _, d := range dir{
			dfs(r+d[0], c+d[1])
		}
	}

	res := 0
	for r := 0; r < ROW; r++{
		for c := 0; c < COL; c++{
			if grid[r][c] == 1{
				cur = 0
				dfs(r, c)
				res = max(res, cur)
			}
		}
	}
	return res
}
