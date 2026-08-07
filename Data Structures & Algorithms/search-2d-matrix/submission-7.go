func searchMatrix(matrix [][]int, target int) bool {
	ROW, COL := len(matrix), len(matrix[0])
	top, bot := 0, ROW-1
	for top <= bot{
		row := top + (bot - top) / 2
		if matrix[row][0] > target{
			bot = row - 1
		} else if matrix[row][COL-1] < target{
			top = row + 1
		} else {
			l, r := 0, COL - 1
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
