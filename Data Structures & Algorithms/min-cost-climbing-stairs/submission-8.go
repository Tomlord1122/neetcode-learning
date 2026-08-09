func minCostClimbingStairs(cost []int) int {
	n := len(cost)
    dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 0
	if n == 0 || n == 1{
		return cost[n]
	}
	for i := 2; i < n+1; i++{
		dp[i] = min(dp[i-2]+cost[i-2], dp[i-1]+cost[i-1])
	}
	return dp[n]
}


// cost = [1, 2, 3] X
// dp[0] = 0, 
// dp[n] = min(dp[n-2] + cost[n-2], dp[n-1]+cost[n-1])