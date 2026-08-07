func diagonalSum(mat [][]int) int {
	res := 0
	for i := 0; i < len(mat); i++{
		for j := 0; j < len(mat[0]); j++{
			if i - j == 0 && i + j != len(mat)-1{
				res += mat[i][j]
			} else if i + j == len(mat) - 1{
				res += mat[i][j]
			}
		}
	}
	return res
}
