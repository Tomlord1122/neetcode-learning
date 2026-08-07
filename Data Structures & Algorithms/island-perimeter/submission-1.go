func islandPerimeter(grid [][]int) int {
	ROW, COL := len(grid), len(grid[0])
	dirs := [][]int{{0,1}, {0,-1}, {1,0}, {-1,0}}

	visited := make(map[[2]int]bool)
	var dfs func(r, c int) int
	dfs = func(r, c int) int{
		if visited[[2]int{r, c}]{
			return 0
		}
		if min(r, c) < 0 || r == ROW || c == COL || grid[r][c] == 0{
			return 1
		}
		visited[[2]int{r,c}] = true
		count := 0
		for _, d := range dirs{
			count += dfs(r+d[0], c+d[1])
		}
		return count
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
