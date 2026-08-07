func canJump(nums []int) bool {
    maxReach := 0
	for i := 0; i < len(nums); i++{
		if i > maxReach{
			return false
		}
		canReach := i + nums[i]
		maxReach = max(maxReach, canReach)
	}
	return true
}
