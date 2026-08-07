func climbStairs(n int) int {
    memoization := make([]int, n+1)
    var helper func(int) int
    helper = func(n int) int {
        if n == 1 || n == 2 {
            return n
        }
        if memoization[n] != 0 {
            return memoization[n]
        }
        memoization[n] = helper(n-1) + helper(n-2)
        return memoization[n]
    }
    return helper(n)
}