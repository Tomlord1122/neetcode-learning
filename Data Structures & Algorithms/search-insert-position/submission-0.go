func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r{
		m := l + (r-l) / 2
		if nums[m] == target{
			return m
		} else if nums[m] < target{
			l = m + 1 
		} else {
			r = m - 1
		}
	}
	return l
}


// return the index 
// sequentially find the index. The time complexity should be O(n)
// Maybe we can use binary search to improve it to O(logn) because 
// the array is sorted