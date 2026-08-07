func isMonotonic(nums []int) bool {
    increasing := nums[0] <= nums[len(nums)-1]

	if increasing{
		for i := 1; i < len(nums); i++{
			if nums[i] < nums[i-1]{
				return false
			}
		}
	} else {
		for i := 1; i < len(nums); i++{
			if nums[i] > nums[i-1]{
				return false
			}
		}
	}
	return true
}