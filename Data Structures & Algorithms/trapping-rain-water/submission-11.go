func trap(height []int) int {
	n := len(height)
	prefix := make([]int, n)
	postfix := make([]int, n)
	prefix[0] = 0
	postfix[n-1] = 0

	for i := 1; i < n; i++{
		prefix[i] = max(prefix[i-1], height[i-1])
	}
	for i := n - 2; i >= 0; i--{
		postfix[i] = max(postfix[i+1], height[i+1])
	}

	res := 0
	for i := 0; i < len(height); i++{
		minH := min(prefix[i], postfix[i])
		res += max(minH - height[i], 0)
	}
	return res
}
