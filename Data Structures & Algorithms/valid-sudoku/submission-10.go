func isValidSudoku(board [][]byte) bool {
	rowCond := make(map[int]map[byte]bool)
	colCond := make(map[int]map[byte]bool)
	gridCond := make(map[int]map[byte]bool)

	for i := 0; i < 9; i++{
		rowCond[i] = make(map[byte]bool)
		colCond[i] = make(map[byte]bool)
		gridCond[i] = make(map[byte]bool)
	}

	ROWS, COLS := len(board), len(board[0])
	for r := 0; r < ROWS; r++{
		for c := 0; c < COLS; c++{
			v := board[r][c]
			if v == '.'{
				continue
			}
			gridIdx := (r/3) * 3 + c / 3
			if rowCond[r][v] || colCond[c][v] || gridCond[gridIdx][v]{
				return false
			}
			rowCond[r][v] = true
			colCond[c][v] = true
			gridCond[gridIdx][v] = true
		}
	}
	return true
}
