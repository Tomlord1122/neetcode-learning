func productExceptSelf(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	
	val := 1
	for i := 0; i < n; i++{
		res[i] = val
		val *= nums[i]
	}
	val = 1
	for i := n - 1; i >= 0; i--{
		res[i] *= val
		val *= nums[i]
	}
	return res
}
