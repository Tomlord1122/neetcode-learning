func isValidSudoku(board [][]byte) bool {
    rowCondition := make(map[int]map[byte]bool)
    colCondition := make(map[int]map[byte]bool)
    gridCondition := make(map[int]map[byte]bool)

    // Init the map in values
    for i := 0; i < 9; i++{
        rowCondition[i] = make(map[byte]bool)
        colCondition[i] = make(map[byte]bool)
        gridCondition[i] = make(map[byte]bool)
    }

    for i := 0; i < 9; i++{
        for j := 0; j < 9; j++{

            num := board[i][j]
            if num == '.'{
                continue
            }
            if rowCondition[i][num] || colCondition[j][num] || gridCondition[(i/3) * 3 + (j/3)][num] {
                return false
            }
            rowCondition[i][num] = true
            colCondition[j][num] = true
            gridCondition[(i/3)*3 +(j / 3)][num] = true
        }
    }

    return true
}
