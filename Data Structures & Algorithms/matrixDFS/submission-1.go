func countPaths(grid [][]int) int {
	visit := make(map[[2]int]bool)
	ROW, COL := len(grid), len(grid[0])
	dir := [][]int{{0,1}, {0,-1}, {-1,0}, {1,0}}
	var dfs func(r, c int) int
	dfs = func(r, c int) int{
		if min(r, c) < 0 || r == ROW || c == COL ||
		visit[[2]int{r, c}] == true || grid[r][c] == 1{
			return 0
		}
		if r == ROW-1 && c == COL-1{
			return 1
		}

		// add the current position into visit
		visit[[2]int{r, c}] = true
		path := 0
		for _, d := range dir{
			path += dfs(r+d[0], c+d[1])
		}
		// remove visit 
		delete(visit, [2]int{r, c})
		return path
	}
	return dfs(0, 0)
}
