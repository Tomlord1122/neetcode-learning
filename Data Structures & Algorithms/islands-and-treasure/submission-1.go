func islandsAndTreasure(grid [][]int) {
    ROWS, COLS := len(grid), len(grid[0])
	dirs := [][]int{{0,1}, {0,-1}, {1,0}, {-1,0}}
	queue := [][2]int{}
	visited := make(map[[2]int]bool)

	for r := 0; r < ROWS; r++{
		for c := 0; c < COLS; c++{
			if grid[r][c] == 0{
				queue = append(queue, [2]int{r, c})
				visited[[2]int{r, c}] = true
			}
		}
	}
	distance := 1
	for len(queue) > 0{
		length := len(queue)
		for i := 0; i < length; i++{
			r, c := queue[0][0], queue[0][1]
			queue = queue[1:]
			for _, d := range dirs{
				nr, nc := r+d[0], c+d[1]
				if min(nr, nc) < 0 || nr == ROWS || nc == COLS ||
				grid[nr][nc] == 0 || grid[nr][nc] == -1 || visited[[2]int{nr,nc}]{
					continue
				}
				grid[nr][nc] = distance
				visited[[2]int{nr, nc}] = true
				queue = append(queue, [2]int{nr, nc})
			}
		}
		distance++
	}
}
