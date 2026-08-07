func findDuplicate(nums []int) int {
    for _, num := range nums{
		idx := abs(num) - 1
		if nums[idx] < 0 {
			// num is existing
			return abs(num)
		}
		nums[idx] *= -1
	}
	return -1
}

func abs(x int) int{
	if x < 0{
		return -x
	}
	return x
}

