func canJump(nums []int) bool {
	n := len(nums)
	dp := make([]bool, n)
	dp[n-1] = true
	for i := n-2; i >= 0; i--{
		for j := 1; j <= nums[i] && i+j < n; j++{
			dp[i] = dp[i] || dp[i+j]
			if dp[i]{
				break
			}
		}
	}
	return dp[0]
}
