func isValidSudoku(board [][]byte) bool {
	rowCond := make(map[int]map[byte]bool)
	colCond := make(map[int]map[byte]bool)
	gridCond := make(map[int]map[byte]bool)

	// init the hashMap
	for i := 0; i < 9; i++{
		rowCond[i] = make(map[byte]bool)
		colCond[i] = make(map[byte]bool)
		gridCond[i] = make(map[byte]bool)
	}

	for i := 0; i < 9; i++{
		for j := 0; j < 9; j++{
			if board[i][j] == '.'{
				continue
			}
			val := board[i][j]
			gridIdx := (i / 3) * 3 + j / 3
			if rowCond[i][val] || colCond[j][val] || gridCond[gridIdx][val]{
				return false
			}
			rowCond[i][val] = true
			colCond[j][val] = true
			gridCond[gridIdx][val] = true
		}
	}
	return true
}
