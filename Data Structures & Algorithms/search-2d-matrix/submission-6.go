func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	top, bot := 0, m-1
	for top <= bot{
		row := top + (bot - top) / 2
		// check if this row is our target
		if matrix[row][0] > target{
			bot = row - 1
		} else if matrix[row][n-1] < target{
			top = row + 1
		} else {
			// We've found the target row
			l, r := 0, n - 1
			for l <= r{
				m := l + (r - l) / 2
				val := matrix[row][m]
				if val < target{
					l = m + 1
				} else if val > target{
					r = m - 1
				} else {
					return true
				}
			}
			return false
		}
	}
	return false
}
