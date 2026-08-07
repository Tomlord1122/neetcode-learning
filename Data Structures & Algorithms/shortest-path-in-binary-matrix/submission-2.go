func shortestPathBinaryMatrix(grid [][]int) int {
	ROW, COL := len(grid), len(grid[0])
	dirs := [][]int{{0,1},{0,-1}, {1,0}, {-1,0}, {1,1},{1,-1}, {-1,1}, {-1,-1}}

	if grid[0][0] == 1{
		return -1
	}

	shortest := 1
	queue := [][]int{{0,0}}
	for len(queue) > 0{
		length := len(queue)
		for i := 0; i < length; i++{
			cur := queue[0]
			r, c := cur[0], cur[1]
			if r == ROW-1 && c == COL-1{
				return shortest
			}
			queue = queue[1:]
			for _, d := range dirs{
				nr, nc := r+d[0], c+d[1]
				if min(nr, nc) < 0 || nr == ROW || nc == COL || grid[nr][nc] == 1{
					continue
				}
				grid[nr][nc] = 1
				queue = append(queue, []int{nr, nc})
			}
		}
		shortest++
	}
	return -1
}
