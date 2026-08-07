func removeElement(nums []int, val int) int {
    l := 0

	for r := 0; r < len(nums); r++{
		if nums[r] == val{
			continue
		} else {
			nums[l] = nums[r]
			l++
		}
	}
	return l
}
