func searchMatrix(matrix [][]int, target int) bool {
    m, n := len(matrix), len(matrix[0])
    top, bot := 0, m - 1
    for top <= bot{
        row := (top + bot) / 2
        if matrix[row][0] > target{
            bot = row - 1
        } else if matrix[row][n-1] < target{
            top = row + 1
        } else {
            l, r := 0, n - 1
            for l <= r{
                mid := (l + r) / 2
                if matrix[row][mid] == target{
                    return true
                } else if matrix[row][mid] < target{
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
