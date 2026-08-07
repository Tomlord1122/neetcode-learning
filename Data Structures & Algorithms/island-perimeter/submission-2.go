func islandPerimeter(grid [][]int) int {
	ROW, COL := len(grid), len(grid[0])
	visited := make(map[[2]int]bool)
	dirs := [][]int{{0,1}, {0,-1}, {1,0}, {-1,0}}
	var dfs func(r, c int) int
	dfs = func(r, c int) int{
		if visited[[2]int{r, c}]{
			return 0
		}
		if min(r, c) < 0 || r == ROW || c == COL || grid[r][c] == 0{
			return 1
		}
		visited[[2]int{r, c}] = true
		res := 0
		for _, dir := range dirs{
			nr, nc := r+dir[0], c + dir[1]
			res += dfs(nr, nc)
		}
		return res
	}
	res := 0
	for r := 0; r < ROW; r++{
		for c := 0; c < COL; c++{
			if grid[r][c] == 1{
				res = dfs(r,c)
				return res
			}
		}
	}
	return res
}
