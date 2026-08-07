func productExceptSelf(nums []int) []int {
	output := make([]int, len(nums))
	val := 1
	for i := 0; i < len(nums); i++{
		output[i] = val
		val *= nums[i]
	}
	val = 1
	for i := len(nums)-1; i >= 0; i--{
		output[i] *= val
		val *= nums[i]
	}
	return output
}

// time complexity -> O(n)

