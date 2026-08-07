func missingNumber(nums []int) int {
	missing := 0
	for i := 0; i < len(nums)+1; i++{
		missing ^= i
	}
	for _, val := range nums{
		missing ^= val
	}
	return missing
}
