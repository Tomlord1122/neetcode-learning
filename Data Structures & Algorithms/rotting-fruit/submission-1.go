func orangesRotting(grid [][]int) int {
    fresh := 0
	ROWS, COLS := len(grid), len(grid[0])
	queue := [][2]int{}
	for r := 0; r < ROWS; r++{
		for c := 0; c < COLS; c++{
			if grid[r][c] == 1{
				fresh++
			} else if grid[r][c] == 2{
				// append current coordinate to queue
				queue = append(queue, [2]int{r, c})
			}
		}
	}

	dirs := [][]int{{0,1}, {0,-1}, {1,0}, {-1, 0}}
	time := 0
	for len(queue) != 0 && fresh != 0{
		length := len(queue)
		for i := 0; i < length; i++{
			pop := queue[0]
			queue = queue[1:]
			nr, nc := 0, 0
			for _, d := range dirs{
				nr = pop[0] + d[0]
				nc = pop[1] + d[1]
				if nr < 0 || nc < 0 || nr >= ROWS || nc >= COLS || grid[nr][nc] != 1{
					continue
				}
				if grid[nr][nc] == 1{
					fresh--
					grid[nr][nc] = 2
					queue = append(queue, [2]int{nr, nc})
				}
			}
		}
		time++
	}
	if fresh != 0{
		return -1
	}
	return time
}
