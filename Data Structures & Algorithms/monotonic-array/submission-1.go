func isMonotonic(nums []int) bool {
    // check the first item and the last item
	// -> This can assume this array is increasing or decreasing
	increasing := true
	if nums[0] > nums[len(nums)-1]{
		increasing = false
	}

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