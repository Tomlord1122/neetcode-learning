func maxAreaOfIsland(grid [][]int) int {
    rows, cols := len(grid), len(grid[0])
    maxArea := 0
    dirs := [][]int{{0,1}, {0,-1}, {1,0}, {-1,0}}
    var dfs func(r, c int) int
    dfs = func(r, c int) int{
        // Check the bounds and water
        if r < 0 || c < 0 || r >= rows || c >= cols || grid[r][c] == 0{
            return 0
        }
        area := 1
        grid[r][c] = 0
        for _, d := range dirs{
            area += dfs(r+d[0], c+d[1])
        }
        return area
    }

    for r := 0; r < rows; r++{
        for c := 0; c < cols; c++{
            if grid[r][c] == 1{
                // Count the area
                area := dfs(r, c)
                maxArea = max(maxArea, area)
            }
        }
    }
    return maxArea
}
