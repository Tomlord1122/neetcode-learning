func shortestPath(grid [][]int) int {
	res := 0
	ROW, COL := len(grid), len(grid[0])
	dirs := [][]int{{0,1}, {0,-1}, {1,0}, {-1,0}}
	visit := make(map[int]map[int]bool)
	for i := 0; i < ROW; i++{
		visit[i] = make(map[int]bool)
	}
	visit[0][0] = true
	queue := [][]int{{0,0}}
	for len(queue) != 0{
		length := len(queue)
		for i := 0; i < length; i++{
			// pop 
			r, c := queue[0][0], queue[0][1]
			if r == ROW-1 && c == COL-1{
				return res
			}
			queue = queue[1:]
			for _, d := range dirs{
				nr, nc := r + d[0], c + d[1]
				// check the condition
				if min(nr, nc) < 0 || nr == ROW || nc == COL || visit[nr][nc] == true ||
				grid[nr][nc] == 1{
					continue
				}
				// nr, nc are valid -> add to visit and queue
				visit[nr][nc] = true
				queue = append(queue, []int{nr, nc})
			}
		}
		res++
	}
	return -1
}
