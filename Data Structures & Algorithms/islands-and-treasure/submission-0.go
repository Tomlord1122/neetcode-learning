func islandsAndTreasure(grid [][]int) {
	queue := [][]int{}
	ROWS, COLS := len(grid), len(grid[0])
	for r := 0; r < ROWS; r++{
		for c := 0; c < COLS; c++{
			if grid[r][c] == 0{
				// append this to queue
				queue = append(queue, []int{r,c})
			}
		}
	}

	dirs := [][]int{{0,1},{0,-1},{1,0},{-1,0}}
	distance := 0
	visited := make(map[[2]int]bool)
	for len(queue) > 0{
		distance++
		length := len(queue)
		for i := 0; i < length; i++{
			r, c := queue[0][0], queue[0][1]
			queue = queue[1:]
			for _, d := range dirs{
				nr, nc := r+d[0], c+d[1]
				if min(nr, nc) < 0 || nr == ROWS || nc == COLS ||
				grid[nr][nc] == -1 || grid[nr][nc] == 0 || visited[[2]int{nr,nc}]{
					continue
				}
				visited[[2]int{nr,nc}] = true
				grid[nr][nc] = distance
				queue = append(queue, []int{nr, nc})
			}
		}
	}
	return 
}
