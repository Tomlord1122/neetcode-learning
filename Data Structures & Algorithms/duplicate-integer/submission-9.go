func hasDuplicate(nums []int) bool {
	numSet := make(map[int]bool)

	for _, num := range nums{
		if _, exist := numSet[num]; exist{
			return true
		}
		numSet[num] = true
	}
	return false
}
