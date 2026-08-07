func productExceptSelf(nums []int) []int {
	n := len(nums)
	val := 1
	res := make([]int, n)
	for i := 0 ; i < n; i++{
		// update nums[i] and val
		res[i] = val
		val = val * nums[i]
	}
	val = 1
	for i := n - 1; i >= 0; i--{
		res[i] *= val
		val = val * nums[i]
	}
	return res
}
