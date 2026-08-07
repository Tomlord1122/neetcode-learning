func longestConsecutive(nums []int) int {
	numSet := make(map[int]bool)
	for _, val := range nums{
		numSet[val] = true
	}

	res := 0
	for _, val := range nums{
		if !numSet[val-1]{
			cur := 1
			for numSet[val+1]{
				cur++
				val++
			}
			res = max(res, cur)
		}
	}
	return res
}
