func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	if image[sr][sc] == color{
		return image
	}
    originColor := image[sr][sc]
	dirs := [][]int{{0,1}, {0,-1}, {1,0}, {-1,0}}
	ROW, COL := len(image), len(image[0])
	var dfs func(r, c int)
	dfs = func(r, c int){
		if min(r, c) < 0 || r == ROW || c == COL || image[r][c] != originColor{
			return
		}
		image[r][c] = color
		for _, d := range dirs{
			dfs(r+d[0], c+d[1])
		}
	}

	dfs(sr, sc)
	return image
}
