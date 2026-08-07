func sortedSquares(nums []int) []int {
	l, r := 0, len(nums)-1
	res := []int{}
	for l <= r{
		leftSquare := nums[l] * nums[l]
		rightSquare := nums[r] * nums[r]
		if leftSquare < rightSquare{
			res = append(res, rightSquare)
			r--
		} else {
			res = append(res, leftSquare)
			l++
		}
	}

	// Reverse the res back to get the non-decreasing order
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1{
		res[i], res[j] = res[j], res[i]
	}
	return res
}
