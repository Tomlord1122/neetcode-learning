func shortestPathBinaryMatrix(grid [][]int) int {
	dirs := [][]int{
		{0,1}, {0,-1}, {1,0}, {-1,0},
		{1,1}, {1,-1}, {-1,-1}, {-1,1},
	}
	n := len(grid)

	if grid[0][0] != 0{
		return -1
	}
	queue := [][]int{{0,0}}
	res := 1
	for len(queue) > 0{
		length := len(queue)
		for i := 0; i < length; i++{
			r, c := queue[0][0], queue[0][1]
			queue = queue[1:]
			if r == n-1 && c == n-1{
				return res
			}
			for _, d := range dirs{
				nr, nc := r+d[0], c+d[1]
				if min(nr, nc) < 0 || nr == n || nc == n || grid[nr][nc] == 1{
					continue
				}
				grid[nr][nc] = 1
				queue = append(queue, []int{nr,nc})
			}
		}
		res++
	}
	return -1
}
