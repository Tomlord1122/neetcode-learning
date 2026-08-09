func pacificAtlantic(heights [][]int) [][]int {
    pac, atl := make(map[[2]int]bool), make(map[[2]int]bool)
	ROWS, COLS := len(heights), len(heights[0])
	dirs := [][]int{{0,1}, {0,-1}, {1,0}, {-1,0}}

	var dfs func(r, c int, visited map[[2]int]bool, prevHeight int)
	dfs = func(r, c int, visited map[[2]int]bool, prevHeight int){
		if min(r, c) < 0 || r == ROWS || c == COLS ||
		visited[[2]int{r,c}] || heights[r][c] < prevHeight{
			return
		}
		visited[[2]int{r,c}] = true
		for _, d := range dirs{
			dfs(r+d[0], c+d[1], visited, heights[r][c])
		}
	}

	for r := 0; r < ROWS; r++{
		dfs(r, 0, pac, heights[r][0])
		dfs(r, COLS-1, atl, heights[r][COLS-1])
	}
	for c := 0; c < COLS; c++{
		dfs(0, c, pac, heights[0][c])
		dfs(ROWS-1, c, atl, heights[ROWS-1][c])
	}
	res := [][]int{}

	for r := 0; r < ROWS; r++{
		for c := 0; c < COLS; c++{
			if pac[[2]int{r,c}] && atl[[2]int{r,c}]{
				res = append(res, []int{r,c})
			}
		}
	}
	return res
}
