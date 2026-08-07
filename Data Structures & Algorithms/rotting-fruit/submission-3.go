func orangesRotting(grid [][]int) int {
    queue := [][2]int{}
	fresh := 0
	ROW, COL := len(grid), len(grid[0])
	for r := 0; r < ROW; r++{
		for c := 0; c < COL; c++{
			if grid[r][c] == 1{
				fresh++
			} else if grid[r][c] == 2{
				// append to queue
				queue = append(queue, [2]int{r, c})
			}
		}
	}

	// bfs to get the time
	dirs := [][]int{{0,1}, {0,-1}, {1,0}, {-1,0}}
	time := 0
	for len(queue) > 0 && fresh > 0{
		length := len(queue)
		for i := 0; i < length; i++{
			r, c := queue[0][0], queue[0][1]
			queue = queue[1:]
			for _, d := range dirs{
				nr, nc := r+d[0], c+d[1]
				if min(nr, nc) < 0 || nr == ROW || nc == COL || grid[nr][nc] != 1{
					continue
				}
				grid[nr][nc] = 2
				fresh--
				queue = append(queue, [2]int{nr, nc})
			}
		}
		time++
	}
	if fresh > 0{
		return -1
	}
	return time
}
