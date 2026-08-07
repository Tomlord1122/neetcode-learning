func shortestPathBinaryMatrix(grid [][]int) int {
	ROW, COL := len(grid), len(grid[0])
	dirs := [][]int{{0,1},{0,-1}, {1,0}, {-1,0}, {1,1}, {1,-1},
	{-1,1},{-1,-1}}
	if grid[0][0] != 0{
		return -1
	}
	queue := [][]int{{0,0}}
	
	// init a visit hashMap
	visit := make(map[int]map[int]bool)
	for i := 0; i < ROW; i++{
		visit[i] = make(map[int]bool)
	}
	visit[0][0] = true
	shortest := 1
	for len(queue) > 0{
		length := len(queue)
		for i := 0; i < length; i++{
			// pop 
			r, c := queue[0][0], queue[0][1]
			if r == ROW -1 && c == COL - 1{
				return shortest
			}
			queue = queue[1:]
			for _, d := range dirs{
				nr, nc := r+d[0], c+d[1]
				if min(nr, nc) < 0 || nr == ROW || nc == COL || grid[nr][nc] == 1 ||
				visit[nr][nc] == true{
					continue
				}
				// add it to visit and queue
				visit[nr][nc] = true
				queue = append(queue, []int{nr, nc})
			}
		}
		shortest++
	}
	return -1
}
