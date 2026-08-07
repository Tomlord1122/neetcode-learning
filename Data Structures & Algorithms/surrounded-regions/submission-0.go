func solve(board [][]byte) {
    ROWS, COLS := len(board), len(board[0])
	dirs := [][]int{{0,1}, {0,-1}, {1,0}, {-1,0}}
	visit := make(map[[2]int]bool)
	var dfs func(r, c int)
	dfs = func(r, c int){
		coord := [2]int{r,c}
		if min(r, c) < 0 || r == ROWS || c == COLS || visit[coord] ||
		board[r][c] == 'X'{
			return 
		}
		// mark the current be visited
		visit[coord] = true
		for _, d := range dirs{
			dfs(r+d[0], c+d[1])
		}
	}


	for c := 0; c < COLS; c++{
		if board[0][c] == 'O'{
			dfs(0, c)
		}
		if board[ROWS-1][c] == 'O'{
			dfs(ROWS-1,c)
		}
	}
	for r := 0; r <ROWS; r++{
		if board[r][0] == 'O'{
			dfs(r, 0)
		}
		if board[r][COLS-1] == 'O'{
			dfs(r, COLS-1)
		}
	}

	for r := 0; r < ROWS; r++{
		for c := 0; c < COLS; c++{
			coord := [2]int{r,c}
			if board[r][c] == 'O' && !visit[coord]{
				board[r][c] = 'X'
			}
		}
	}
}
