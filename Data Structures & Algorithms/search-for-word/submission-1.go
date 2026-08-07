func exist(board [][]byte, word string) bool {
	ROWS, COLS := len(board), len(board[0])
	path := make(map[[2]int]bool)
	dirs := [][]int{{0,1}, {0,-1}, {1,0}, {-1,0}}
	var dfs func(r, c, i int) bool
	dfs = func(r, c, i int) bool{
		if i == len(word){
			return true
		}
		if min(r, c) < 0 || r == ROWS || c ==COLS ||
		word[i] != board[r][c] || path[[2]int{r, c}]{
			return false
		}
		path[[2]int{r,c}] = true
		res := false
		for _, d := range dirs{
			if dfs(r+d[0], c+d[1], i+1){
				res = true
			}
		}
		path[[2]int{r,c}] = false
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
