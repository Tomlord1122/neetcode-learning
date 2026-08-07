func hasDuplicate(nums []int) bool {
    numMap := make(map[int]bool)
	for _, num := range nums{
		if _, exist := numMap[num]; exist{
			return true
		}
		numMap[num] = true
	}
	return false
}
