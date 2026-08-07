func canJump(nums []int) bool {
    maxDistance := 0
	for i := 0; i < len(nums); i++{
		if i > maxDistance{
			return false
		}
		curMax := i + nums[i]
		maxDistance = max(maxDistance, curMax)
	}
	return true
}
