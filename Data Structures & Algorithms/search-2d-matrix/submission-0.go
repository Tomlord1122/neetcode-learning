func searchMatrix(matrix [][]int, target int) bool {
    rows, cols := len(matrix), len(matrix[0])

    top, bot := 0, rows - 1
    for top <= bot{
        midRow := top + (bot - top) / 2
        if matrix[midRow][0] > target{
            bot = midRow - 1
        } else if matrix[midRow][cols-1] < target{
            top = midRow + 1
        } else {
            break
        }
    }
    if top > bot{
        return false
    }

    // Use a calssical binary search to find the target
    row := top + (bot - top) / 2
    left, right := 0, cols - 1
    for left <= right{
        m := left + (right - left) / 2
        if matrix[row][m] < target{
            left = m + 1
        } else if matrix[row][m] > target{
            right = m - 1
        } else {
            return true
        }
    }
    return false
}
