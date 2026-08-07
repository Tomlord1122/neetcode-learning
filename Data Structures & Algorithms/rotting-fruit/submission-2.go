func orangesRotting(grid [][]int) int {
    fresh := 0
	time := 0
	queue := [][2]int{}
	ROW, COL := len(grid), len(grid[0])
	for r := 0; r < ROW; r++{
		for c := 0; c < COL; c++{
			if grid[r][c] == 1{
				fresh++
			} else if grid[r][c] == 2{
				queue = append(queue, [2]int{r,c})
			}
		}
	}

	dirs := [][]int{{0,1}, {0,-1}, {1,0}, {-1,0}}

	for len(queue) > 0 && fresh > 0 {
		length := len(queue)
		for i := 0; i < length; i++{
			pop := queue[0]
			queue = queue[1:]	
			r, c := pop[0], pop[1]
			for _, d := range dirs{
				nr, nc := r+d[0], c+d[1]
				if min(nr, nc) < 0 || nr == ROW || nc == COL || grid[nr][nc] != 1{
					continue
				}
				grid[nr][nc] = 2
				fresh--
				// put nr, nc into queue
				queue = append(queue, [2]int{nr, nc})
			}
		}
		time++
	}
	if fresh != 0{
		return -1
	}
	return time
}