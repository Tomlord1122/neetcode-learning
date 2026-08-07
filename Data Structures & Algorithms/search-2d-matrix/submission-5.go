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
			// find the target row
			l, r := 0, n - 1
			for l <= r{
				m := (l + r) / 2
				if matrix[row][m] == target{
					return true
				} else if matrix[row][m] < target{
					l = m + 1
				} else {
					r = m - 1
				}
			}
			return false
		}
	}
	return false
}


// First we need to find out the target row
// Second we find the target in the row
