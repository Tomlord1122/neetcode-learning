func findMaxConsecutiveOnes(nums []int) int {
	count := 0
	res := 0
	for _, val := range nums{
		if val == 1{
			count++
			res = max(res, count)
		} else {
			count = 0
		}
	}
	return res
}
