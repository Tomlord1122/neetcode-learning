func isValidSudoku(board [][]byte) bool {
    rowCond := make(map[int] map[byte]bool)
    colCond := make(map[int] map[byte]bool)
    gridCond := make(map[int] map[byte]bool)

    // Init the value map[byte]bool
    for i := 0; i < 9; i++{
        rowCond[i] = make(map[byte]bool)
        colCond[i] = make(map[byte]bool)
        gridCond[i] = make(map[byte]bool)
    }

    for i := 0; i < 9; i++{
        for j := 0; j < 9; j++{
            val := board[i][j]
            if val == '.'{
                continue
            }
            gridIdx := (i / 3) * 3 + (j / 3)
            // If the val is already exist, then we should return false
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
