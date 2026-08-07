func longestConsecutive(nums []int) int {
	numSet := make(map[int]bool)
	for _, num := range nums{
		numSet[num] = true
	}
	res := 0
	for _, num := range nums{
		// In this case, we found the start number
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
