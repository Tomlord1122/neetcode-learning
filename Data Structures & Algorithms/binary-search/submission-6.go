func search(nums []int, target int) int {
	// search for a target using binary search
	l, r := 0, len(nums)-1
	for l <= r{
		m := l + (r - l) / 2
		if nums[m] == target{
			return m
		} else if nums[m] < target{
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return -1
}
