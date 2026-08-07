func searchMatrix(matrix [][]int, target int) bool {
    m, n := len(matrix), len(matrix[0])
    top, bot := 0, m-1
    for top <= bot{
        row := (top + bot) / 2
        if target < matrix[row][0]{
            bot = row - 1
        } else if target > matrix[row][n-1]{
            top = row + 1
        } else {
            l, r := 0, n-1
            for l <= r{
                mid := (l + r) / 2
                val := matrix[row][mid]
                if val == target{
                    return true
                } else if val < target{
                    l = mid + 1
                } else {
                    r = mid - 1
                }
            }
            return false
        }
    }
    return false
}
