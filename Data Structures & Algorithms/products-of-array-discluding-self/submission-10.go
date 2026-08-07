func productExceptSelf(nums []int) []int {
	prefix := make([]int, len(nums))
	postfix := make([]int, len(nums))
	prefix[0] = 1
	postfix[len(nums)-1] = 1
	for i := 1; i < len(nums); i++{
		prefix[i] = prefix[i-1] * nums[i-1]
	}
	for i := len(nums)-2; i >= 0; i--{
		postfix[i] = postfix[i+1] * nums[i+1]
	}
	res := make([]int, len(nums))

	for i := 0; i < len(nums); i++{
		res[i] = prefix[i] * postfix[i]
	}
	return res
}
