func longestConsecutive(nums []int) int {
	seen := make(map[int]bool)
	for _, num := range nums{
		seen[num] = true
	}

	res := 0

	for _, num := range nums{
		if seen[num-1]{
			continue
		}
		length := 1
		n := num + 1
		for seen[n]{
			n++
			length++
		}
		res = max(res, length)
	}
	return res
}
