func longestConsecutive(nums []int) int {
	numSet := make(map[int]bool)
	for _, num := range nums{
		numSet[num] = true
	}
	res := 0
	for _, num := range nums{
		if numSet[num-1]{
			continue
		}
		cur := 1
		for numSet[num+1]{
			cur++
			num++
		}
		res = max(res, cur)
	}
	return res
}
