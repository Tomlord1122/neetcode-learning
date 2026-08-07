func searchMatrix(matrix [][]int, target int) bool {
	top, bot := 0, len(matrix) - 1
	col := len(matrix[0])
	for top <= bot{
		row := (top + bot) / 2
		// check the first element and last element in this row
		if matrix[row][0] > target{
			bot = row - 1
		} else if matrix[row][col-1] < target{
			top = row + 1
		} else {
			l, r := 0, col - 1
			for l <= r{
				m := (l+r) / 2
				if matrix[row][m] < target{
					l = m + 1
				} else if matrix[row][m] > target{
					r = m - 1
				} else {
					return true
				}
			}
			break
		}
	}
	return false
}