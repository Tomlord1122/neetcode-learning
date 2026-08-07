func orangesRotting(grid [][]int) int {
	time := 0
	fresh := 0
	row, col := len(grid), len(grid[0])
	queue := [][2]int{}
	for r := 0; r < row; r++{
		for c := 0; c < col; c++{
			if grid[r][c] == 1{
				fresh++
			} else if grid[r][c] == 2{
				queue = append(queue, [2]int{r, c})
			}
		}
	}

	directions := [][]int{{0,1}, {0,-1}, {1,0}, {-1,0}}

	for fresh > 0 && len(queue) > 0{
		length := len(queue)
		for i := 0; i < length; i++{
			current := queue[0]
			queue = queue[1:]


			for _, dir := range directions{
				newRow := current[0] + dir[0]
				newCol := current[1] + dir[1]

				if newRow >= 0 && newRow < row && newCol >= 0 && newCol < col && grid[newRow][newCol] == 1{
					grid[newRow][newCol] = 2
					queue = append(queue, [2]int{newRow, newCol})
					fresh--
				}
			}
		}
		time++
	}
	if fresh == 0{
		return time
	}

	return -1
}
