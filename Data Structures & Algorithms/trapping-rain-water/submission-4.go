func trap(height []int) int {
    area := 0
    n := len(height)
    prefixH := make([]int, n)
    postfixH := make([]int, n)
    prefixH[0] = 0
    postfixH[n-1] = 0
    for i := 1; i < n; i++{
        prefixH[i] = max(prefixH[i-1], height[i-1])
    }
    for i := n - 2; i >= 0; i--{
        postfixH[i] = max(postfixH[i+1], height[i+1])
    }

    for i := 0; i < n; i++{
        minH := min(prefixH[i], postfixH[i])
        if minH > height[i]{
            area += minH - height[i]
        }
    }
    return area
}
