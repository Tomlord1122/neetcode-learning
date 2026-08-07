func islandPerimeter(grid [][]int) int {
	ROW, COL := len(grid), len(grid[0])
	dirs := [][]int{{0,1}, {0,-1}, {1, 0}, {-1,0}}
	visited := make(map[[2]int]bool)
	var dfs func(r, c int) int
	dfs = func(r, c int) int{
		// base condition
		if min(r, c) < 0 || r == ROW || c == COL || grid[r][c] == 0{
			return 1
		}
		if visited[[2]int{r, c}]{
			return 0
		}
		res := 0
		visited[[2]int{r, c}] = true
		for _, d := range dirs{
			res += dfs(r+d[0], c+d[1])
		}
		return res
	}

	for r := 0; r < ROW; r++{
		for c := 0; c < COL; c++{
			if grid[r][c] == 1{
				return dfs(r, c)
			}
		}
	}
	return -1
}
