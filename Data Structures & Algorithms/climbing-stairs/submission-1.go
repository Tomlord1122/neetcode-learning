func climbStairs(n int) int {
    if n == 0 || n == 1{
        return 1
    }
    a, b := 1, 1
    for i := 2; i <= n; i++{
        tmp := b
        b = a + b
        a = tmp
    }
    return b
}
