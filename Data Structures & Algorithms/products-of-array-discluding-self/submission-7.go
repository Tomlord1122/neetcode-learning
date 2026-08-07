func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	val := 1
	res[0] = 1
	for i := 1; i < len(nums); i++{
		res[i] = nums[i-1] * val
		val = res[i]
	}
	val = 1
	for i := len(nums)-1; i >= 0; i--{
		res[i] = res[i] * val
		val = val * nums[i]
	}
	return res
}
