func exist(board [][]byte, word string) bool {
	ROWS, COLS := len(board), len(board[0])
	visited := make(map[[2]int]bool)
	dirs := [][]int{{0,1}, {0,-1}, {1,0}, {-1,0}}

	var dfs func(r, c, i int) bool
	dfs = func(r, c, i int) bool{
		if i == len(word){
			return true
		}
		if min(r, c) < 0 || r == ROWS || c ==COLS ||
		visited[[2]int{r, c}] || board[r][c] != word[i]{
			return false
		}
		// add the current index into visited
		visited[[2]int{r,c}] = true
		// recursive call
		res := false
		for _, d := range dirs{
			res = res || dfs(r+d[0], c+d[1], i+1)
		}
		visited[[2]int{r,c}] = false
		return res
	}

	for r := 0; r < ROWS; r++{
		for c := 0; c < COLS; c++{
			if board[r][c] == word[0]{
				if dfs(r, c, 0){
					return true
				}
			}
		}
	}
	return false
}
