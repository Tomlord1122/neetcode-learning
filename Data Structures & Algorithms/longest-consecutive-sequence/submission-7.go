func longestConsecutive(nums []int) int {
	numSet := make(map[int]bool)
	for _, val := range nums{
		numSet[val] = true
	}
	res := 0
	for _, num := range nums{
		if numSet[num-1] == false{
			length := 1
			for numSet[num+1]{
				length++
				num++
			}
			res = max(res, length)
		}
	}
	return res
}
