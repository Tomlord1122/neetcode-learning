func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)
	l, r := 0, k - 1
	res := math.MaxInt
	for r < len(nums){	
		res = min(res, nums[r]-nums[l])
		l++
		r++
	}
	return res
}
