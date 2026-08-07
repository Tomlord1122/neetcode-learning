func pacificAtlantic(heights [][]int) [][]int {
    ROWS, COLS := len(heights), len(heights[0])
	pac := make(map[[2]int]bool)
	atl := make(map[[2]int]bool)
	dirs := [][]int{{0,1},{0,-1},{1,0},{-1,0}}
	var dfs func(r, c int, visited map[[2]int]bool, prevHeight int)
	dfs = func(r, c int, visited map[[2]int]bool, prevHeight int){
		// base case
		coord := [2]int{r, c}
		if r < 0 || c < 0 || r == ROWS || c == COLS ||
		 visited[coord] || heights[r][c] < prevHeight{
			return
		}
		visited[coord] = true
		for _, d := range dirs{
			dfs(r+d[0], c+d[1], visited, heights[r][c])
		}
	}
	for c := 0; c < COLS; c++{
		dfs(0, c, pac, heights[0][c])
		dfs(ROWS-1, c, atl, heights[ROWS-1][c])
	}
	for r := 0; r < ROWS; r++{
		dfs(r, 0, pac, heights[r][0])
		dfs(r, COLS-1, atl, heights[r][COLS-1])
	}
	res := [][]int{}
	for r := 0; r < ROWS; r++{
		for c := 0; c < COLS; c++{
			coord := [2]int{r,c}
			if pac[coord] && atl[coord]{
				res = append(res, []int{r, c})
			}
		}
	}
	return res
}