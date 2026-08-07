func hasDuplicate(nums []int) bool {
    numSet := make(map[int]bool) // key -> exist
	for _, num := range nums{
		if _, exist := numSet[num]; exist{
			return true
		}
		numSet[num] = true
	}
	return false
}

// N -> time complexity O(N)
// N -> space complexity O(N)

