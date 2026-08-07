func floodFill(image [][]int, sr int, sc int, color int) [][]int {
    originColor := image[sr][sc]
	// edge case: this is for avoiding infinite recursive call
    if originColor == color {
        return image
    }
	dir := [][]int{{0,1}, {0,-1}, {1,0}, {-1, 0}}


	ROW, COL := len(image), len(image[0])
	var dfs func(r, c int)
	dfs = func(r, c int){
		if min(r, c) < 0 || r == ROW || c == COL || image[r][c] != originColor{
			return
		}
		image[r][c] = color
		// recursive call
		for _, d := range dir{
			dfs(r+d[0], c+d[1])
		}
		return
	}

	dfs(sr, sc)
	return image
}