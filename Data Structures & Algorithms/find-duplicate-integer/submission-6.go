func findDuplicate(nums []int) int {
    for _, val := range nums{
		idx := abs(val)-1
		if nums[idx] < 0{
			return abs(val)
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