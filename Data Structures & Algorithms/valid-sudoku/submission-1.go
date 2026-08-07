func isValidSudoku(board [][]byte) bool {
    rowCondition := make(map[int] map[byte]bool)
    colCondition := make(map[int] map[byte]bool)
    gridCondition := make(map[int] map[byte]bool)

    for i := 0; i < 9; i++{
        rowCondition[i] = make(map[byte]bool)
        colCondition[i] = make(map[byte]bool)
        gridCondition[i] = make(map[byte]bool)
    }

    for i := 0; i < 9; i++{
        for j := 0; j < 9; j++{
            val := board[i][j]
            if val == '.'{
                continue
            }
            gridIdx := (i / 3) * 3 + (j / 3)
            if rowCondition[i][val] || colCondition[j][val] || gridCondition[gridIdx][val]{
                return false
            }
            rowCondition[i][val] = true
            colCondition[j][val] = true
            gridCondition[gridIdx][val] = true
        }
    }
    return true
}
