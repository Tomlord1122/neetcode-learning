func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	cur := 1
	for i, val := range nums{
		res[i] = cur
		cur *= val
	}
	cur = 1
	for i := len(nums)-1; i >= 0; i--{
		res[i] *= cur
		cur *= nums[i]
	}
	return res
}
