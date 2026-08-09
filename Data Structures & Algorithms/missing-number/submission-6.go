func missingNumber(nums []int) int {
	res := 0
	for i := 0; i <= len(nums); i++{
		res ^= i
	}
	for _, num := range nums{
		res ^= num
	}
	return res
}


// 5
// 0, 1, 2, 3, 4
