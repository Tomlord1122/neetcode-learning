func productExceptSelf(nums []int) []int {
	val := 1
	n := len(nums)
	res := make([]int, n)
	for i := 0; i < n; i++{
		res[i] = val
		val = val * nums[i]
	}
	val = 1
	for i := n - 1; i >= 0; i--{
		res[i] = res[i] * val
		val = val * nums[i]
	}
	return res
}
