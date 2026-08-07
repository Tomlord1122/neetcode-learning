func numIslands(grid [][]byte) int {
    rows, cols := len(grid), len(grid[0])
    dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

    var dfs func(r, c int)
    dfs = func(r, c int) {
        // Check bounds and water
        if r < 0 || c < 0 || r >= rows || c >= cols || grid[r][c] == '0' {
            return
        }
        // Mark as visited
        grid[r][c] = '0'
        // Recursively call in four directions
        for _, d := range dirs {
            dfs(r+d[0], c+d[1])
        }
    }

    islandCount := 0
    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            if grid[r][c] == '1' {
                dfs(r, c)
                islandCount++
            }
        }
    }
    return islandCount
}

// A typical graph problem where we use DFS to count the number of islands.